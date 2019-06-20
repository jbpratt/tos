package models

import (
	"errors"
	"fmt"
	"sync"

	mookiespb "github.com/jbpratt78/tos/protofiles"
	"github.com/jmoiron/sqlx"
)

type MenuDB interface {
	SeedMenu() error
	CreateMenuItem(*mookiespb.Item) error
	// DeleteMenuItem() error
	// CreateMenuItemOption() error
	GetMenu() (*mookiespb.Menu, error)
}

type MenuService interface {
	MenuDB
}

var _ MenuDB = &menuDB{}

type menuService struct {
	MenuDB
}

type menuDB struct {
	sync.RWMutex
	db *sqlx.DB
}

func NewMenuService(db *sqlx.DB) MenuService {
	mdb := &menuDB{db: db}
	return &menuService{mdb}
}

func (m *menuDB) SeedMenu() error {
	m.Lock()
	defer m.Unlock()

	tx, err := m.db.Begin()
	if err != nil {
		return err
	}

	for i, category := range Menu {
		_, err := tx.Exec("INSERT INTO categories (name) VALUES (?)", category.GetName())
		if err != nil {
			tx.Rollback()
			return err
		}
		for _, item := range category.GetItems() {
			result, err := tx.Exec(
				"INSERT INTO items (name, price, category_id) VALUES (?,?,?)",
				item.GetName(), item.GetPrice(), i+1)
			if err != nil {
				tx.Rollback()
				return err
			}
			itemid, _ := result.LastInsertId()
			for _, option := range item.GetOptions() {
				res, err := tx.Exec(
					"INSERT INTO options (name, price, selected) VALUES (?,?,?)",
					option.GetName(), option.GetPrice(), option.GetSelected())
				if err != nil {
					tx.Rollback()
					return err
				}
				optionid, _ := res.LastInsertId()
				_, err = tx.Exec(
					"INSERT INTO item_options (item_id, option_id) VALUES (?,?)",
					itemid, optionid)
				if err != nil {
					tx.Rollback()
					return err
				}
			}
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (m *menuDB) GetMenu() (*mookiespb.Menu, error) {
	m.RLock()
	defer m.RUnlock()
	var categories []*mookiespb.Category
	menu := &mookiespb.Menu{
		Categories: categories,
	}
	err := m.db.Select(&menu.Categories, "SELECT * from categories")
	for _, category := range menu.GetCategories() {
		err = m.db.Select(&category.Items,
			fmt.Sprintf("SELECT * FROM items WHERE category_id = %v", category.GetId()))
		if err != nil {
			return nil, err
		}
		for _, item := range category.GetItems() {
			err = m.db.Select(&item.Options, fmt.Sprintf(
				`
				SELECT name,price,selected,options.id 
				FROM options JOIN item_options as io ON options.id = io.option_id 
				WHERE item_id = %d`, item.GetId()))
			if err != nil {
				return nil, err
			}
		}
	}
	return menu, nil
}

// need to reload
func (m *menuDB) CreateMenuItem(item *mookiespb.Item) error {
	m.Lock()
	defer m.Unlock()

	_, err := m.db.Exec(
		"INSERT INTO items (name, price, category_id) VALUES (?,?,?)",
		item.GetName(), item.GetPrice(), item.GetCategoryID())
	if err != nil {
		return err
	}

	return nil
}

func (m *menuDB) DeleteMenuItem(id int32) error {
	m.Lock()
	defer m.Unlock()

	res, err := m.db.Exec(
		"DELETE FROM items WHERE id = ?", id)

	aff, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if aff == 0 {
		return errors.New("ID not found")
	}

	return nil
}
