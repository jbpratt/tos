package models

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	mookiespb "github.com/jbpratt78/tos/protofiles"
	"github.com/jmoiron/sqlx"
)

type OrderDB interface {
	GetOrders() ([]*mookiespb.Order, error)
	CompleteOrder(id int32) error
	SubmitOrder(o *mookiespb.Order) error
}

type OrderService interface {
	OrderDB
}

var _ OrderDB = &orderService{}

type orderService struct {
	OrderDB
}

type orderDB struct {
	sync.RWMutex
	db *sqlx.DB
}

const orderSchema = `
CREATE TABLE IF NOT EXISTS orders (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  status TEXT NOT NULL,
  total DECIMAL NOT NULL,
  time_ordered TEXT NOT NULL,
  time_complete TEXT
);

CREATE TABLE IF NOT EXISTS order_items (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  item_id INTEGER NOT NULL,
  order_id INTEGER NOT NULL,
  FOREIGN KEY (item_id) REFERENCES items(id),
  FOREIGN KEY (order_id) REFERENCES orders(id)
);

CREATE TABLE IF NOT EXISTS order_item_options (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  order_item_id INTEGER NOT NULL,
  option_id INTEGER NOT NULL,
  FOREIGN KEY (order_item_id) REFERENCES order_items(id),
  FOREIGN KEY (option_id) REFERENCES options(id)
);`

func NewOrderService(db *sqlx.DB) (OrderService, error) {
	_, err := db.Exec(orderSchema)
	if err != nil {
		return nil, err
	}
	odb := &orderDB{db: db}
	return &orderService{odb}, nil
}

func (o *orderDB) SubmitOrder(order *mookiespb.Order) error {
	o.Lock()
	defer o.Unlock()
	tx, err := o.db.Begin()
	if err != nil {
		return err
	}
	res, err := tx.Exec(
		"INSERT INTO orders (name, total, status, time_ordered, time_complete) VALUES (?, ?, ?, ?, ?)",
		order.GetName(), order.GetTotal(), order.GetStatus(), time.Now().Format("2006-01-02 15:04:05"), "")
	if err != nil {
		tx.Rollback()
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	order.Id = int32(id)
	if order.GetItems() != nil {
		err = submitOrderItems(tx, order)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func submitOrderItems(tx *sql.Tx, order *mookiespb.Order) error {
	for _, item := range order.GetItems() {
		res, err := tx.Exec(
			"INSERT INTO order_items (item_id, order_id) VALUES (?, ?)",
			item.GetId(), order.GetId())
		if err != nil {
			tx.Rollback()
			return err
		}

		orderItemID, err := res.LastInsertId()
		if err != nil {
			return err
		}
		item.OrderItemID = int32(orderItemID)

		if item.GetOptions() != nil {
			err = submitOrderItemOptions(tx, item)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return nil
}

func submitOrderItemOptions(tx *sql.Tx, item *mookiespb.Item) error {
	for _, option := range item.GetOptions() {
		if option.GetSelected() {
			_, err := tx.Exec(
				"INSERT INTO order_item_options (order_item_id, option_id) VALUES (?, ?)",
				item.GetOrderItemID(), option.GetId(),
			)

			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return nil
}

func (o *orderDB) GetOrders() ([]*mookiespb.Order, error) {
	o.RLock()
	defer o.RUnlock()
	orders, err := o.getOrders()
	if err != nil {
		return nil, err
	}
	for _, order := range orders {
		err = o.getOrderItems(order)
		if err != nil {
			return nil, err
		}
		for _, item := range order.GetItems() {
			err = o.getOrderItemOptions(item, order.GetId())
			if err != nil {
				return nil, err
			}
		}
	}
	return orders, nil
}

func (o *orderDB) getOrders() ([]*mookiespb.Order, error) {
	var orders []*mookiespb.Order
	err := o.db.Select(&orders,
		"SELECT * FROM orders WHERE status = 'active'")
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *orderDB) getOrderItems(order *mookiespb.Order) error {
	err := o.db.Select(&order.Items, fmt.Sprintf(
		`
		SELECT name,price,items.id,order_items.id as order_item_id
		FROM items JOIN order_items ON items.id = order_items.item_id 
		WHERE order_id = %d`, order.GetId()))
	if err != nil {
		return err
	}

	return nil
}

func (o *orderDB) getOrderItemOptions(item *mookiespb.Item, id int32) error {
	err := o.db.Select(&item.Options, fmt.Sprintf(
		`
		SELECT options.name,options.price 
		FROM order_item_options AS oio CROSS JOIN order_items
		CROSS JOIN options WHERE order_item_id = order_items.id
		AND oio.option_id = options.id 
		AND order_id = %d
		AND item_id = %d
		AND order_item_id = %d`, id, item.GetId(), item.GetOrderItemID()))
	if err != nil {
		return err
	}
	for _, option := range item.GetOptions() {
		option.Selected = true
	}

	return nil
}

func (o *orderDB) CompleteOrder(id int32) error {
	o.Lock()
	defer o.Unlock()
	tx, err := o.db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Exec(
		"UPDATE orders SET status = ? WHERE id = ?", "complete", id); err != nil {
		tx.Rollback()
		return err
	}
	if _, err = tx.Exec(
		"UPDATE orders SET time_complete = ? WHERE id = ?",
		time.Now().Format("2006-01-02 15:04:05"), id); err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
