package models

import (
	"fmt"
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
	db *sqlx.DB
}

func NewOrderService(db *sqlx.DB) OrderService {
	odb := &orderDB{db}
	return &orderService{odb}
}

func (o *orderDB) SubmitOrder(order *mookiespb.Order) error {
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
	for _, item := range order.GetItems() {
		res, err := tx.Exec(
			"INSERT INTO order_items (item_id, order_id) VALUES (?, ?)",
			item.GetId(), order.GetId())
		if err != nil {
			tx.Rollback()
			return err
		}
		orderItemID, _ := res.LastInsertId()
		item.OrderItemID = int32(orderItemID)

		for _, option := range item.GetOptions() {
			if option.GetSelected() {
				res, err = tx.Exec(
					"INSERT INTO order_item_options (order_item_id, option_id) VALUES (?, ?)",
					orderItemID, option.GetId(),
				)

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

func (o *orderDB) GetOrders() ([]*mookiespb.Order, error) {
	var orders []*mookiespb.Order
	err := o.db.Select(&orders,
		"SELECT * FROM orders WHERE status = 'active'")
	if err != nil {
		return nil, err
	}
	for _, order := range orders {
		err = o.db.Select(&order.Items, fmt.Sprintf(
			`
			SELECT name,price,items.id,order_items.id as order_item_id
			FROM items JOIN order_items ON items.id = order_items.item_id 
			WHERE order_id = %d`, order.GetId()))
		if err != nil {
			return nil, err
		}
		for _, item := range order.GetItems() {
			err = o.db.Select(&item.Options, fmt.Sprintf(
				`
				SELECT options.name,options.price 
				FROM order_item_options AS oio CROSS JOIN order_items
				CROSS JOIN options WHERE order_item_id = order_items.id
				AND oio.option_id = options.id 
				AND order_id = %d
				AND item_id = %d
				AND order_item_id = %d`, order.GetId(), item.GetId(), item.GetOrderItemID()))
			if err != nil {
				return nil, err
			}
			for _, option := range item.GetOptions() {
				option.Selected = true
			}
		}
	}
	return orders, nil
}

func (o *orderDB) CompleteOrder(id int32) error {
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
