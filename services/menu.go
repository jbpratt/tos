package services

import (
	"errors"
	"fmt"
	"sync"

	mookiespb "github.com/jbpratt78/tos/protofiles"
	"github.com/jmoiron/sqlx"
)

// MenuDB is everything that interacts with the database
// involving the menu
type MenuDB interface {
	SeedMenu() error
	CreateMenuItem(*mookiespb.Item) (int64, error)
	DeleteMenuItem(int64) error
	UpdateMenuItem(*mookiespb.Item) error
	// CreateMenuItemOption() error
	GetMenu() (*mookiespb.Menu, error)
}

// MenuService the the abstraction for the MenuDB
type MenuService interface {
	MenuDB
}

var _ MenuDB = (*menuDB)(nil)

type menuService struct {
	MenuDB
}

type menuDB struct {
	sync.RWMutex
	db *sqlx.DB
}

const menuSchema = `
CREATE TABLE IF NOT EXISTS categories (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS items (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  price DECIMAL NOT NULL,
  category_id INTEGER NOT NULL, 
  CONSTRAINT fk_categories
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE IF NOT EXISTS options (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  price DECIMAL NOT NULL,
  selected BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS item_options (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  item_id INTEGER NOT NULL,
  option_id INTEGER NOT NULL,
  FOREIGN KEY (item_id) REFERENCES items(id),
  FOREIGN KEY (option_id) REFERENCES options(id)
);`

// NewMenuService creates a menu service for interacting with the database
func NewMenuService(db *sqlx.DB) (MenuService, error) {
	_, err := db.Exec(menuSchema)
	if err != nil {
		return nil, err
	}

	return &menuService{&menuDB{db: db}}, nil
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

// TODO: stop using fmt.Sprintf to format queries
func (m *menuDB) GetMenu() (*mookiespb.Menu, error) {
	m.RLock()
	defer m.RUnlock()
	var categories []*mookiespb.Category
	menu := &mookiespb.Menu{
		Categories: categories,
	}
	err := m.db.Select(&menu.Categories, "SELECT * from categories")
	for _, category := range menu.GetCategories() {
		if err = m.db.Select(&category.Items,
			fmt.Sprintf("SELECT * FROM items WHERE category_id = %v", category.GetId())); err != nil {
			return nil, err
		}
		for _, item := range category.GetItems() {
			if err = m.db.Select(&item.Options, fmt.Sprintf(
				`
				SELECT name,price,selected,options.id 
				FROM options JOIN item_options as io ON options.id = io.option_id 
				WHERE item_id = %d`, item.GetId())); err != nil {
				return nil, err
			}
		}
	}
	return menu, nil
}

// need to reload
func (m *menuDB) CreateMenuItem(item *mookiespb.Item) (int64, error) {
	m.Lock()
	defer m.Unlock()

	res, err := m.db.Exec(
		"INSERT INTO items (name, price, category_id) VALUES (?,?,?)",
		item.GetName(), item.GetPrice(), item.GetCategoryID())
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (m *menuDB) UpdateMenuItem(item *mookiespb.Item) error {
	m.Lock()
	defer m.Unlock()

	_, err := m.db.Exec(
		"UPDATE items SET name = ?, price = ?, category_id = ? WHERE id = ?",
		item.GetName(), item.GetPrice(), item.GetCategoryID(), item.GetId())

	if err != nil {
		return err
	}
	return nil
}

func (m *menuDB) DeleteMenuItem(id int64) error {
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
