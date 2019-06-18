package database

import (
	"fmt"
	"time"

	mookiespb "github.com/jbpratt78/tos/protofiles"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	DB *sqlx.DB
}

func (s *Service) SeedData() error {
	tx, err := s.DB.Begin()
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

func (s *Service) SubmitOrder(o *mookiespb.Order) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	res, err := tx.Exec(
		"INSERT INTO orders (name, total, status, time_ordered, time_complete) VALUES (?, ?, ?, ?, ?)",
		o.GetName(), o.GetTotal(), o.GetStatus(), time.Now().Format("2006-01-02 15:04:05"), "")
	if err != nil {
		tx.Rollback()
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	o.Id = int32(id)
	for _, item := range o.GetItems() {
		res, err := tx.Exec(
			"INSERT INTO order_items (item_id, order_id) VALUES (?, ?)",
			item.GetId(), o.GetId())
		if err != nil {
			tx.Rollback()
			return err
		}
		orderItemID, _ := res.LastInsertId()
		item.OrderItemID = int32(orderItemID)

		for _, option := range item.GetOptions() {
			if option.GetSelected() {
				res, err = tx.Exec(
					"INSERT INTO order_item_option (order_item_id, option_id) VALUES (?, ?)",
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

func (s *Service) GetMenu() (*mookiespb.Menu, error) {
	var categories []*mookiespb.Category
	menu := &mookiespb.Menu{
		Categories: categories,
	}
	err := s.DB.Select(&menu.Categories, "SELECT * from categories")
	for _, category := range menu.GetCategories() {
		err = s.DB.Select(&category.Items,
			fmt.Sprintf("SELECT * FROM items WHERE category_id = %v", category.GetId()))
		if err != nil {
			return nil, err
		}
		for _, item := range category.GetItems() {
			err = s.DB.Select(&item.Options, fmt.Sprintf(
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

func (s *Service) GetOrders() ([]*mookiespb.Order, error) {
	var orders []*mookiespb.Order
	err := s.DB.Select(&orders,
		"SELECT * FROM orders WHERE status = 'active'")
	if err != nil {
		return nil, err
	}
	for _, order := range orders {
		err = s.DB.Select(&order.Items, fmt.Sprintf(
			`
			SELECT name,price,items.id,order_items.id as order_item_id
			FROM items JOIN order_items ON items.id = order_items.item_id 
			WHERE order_id = %d`, order.GetId()))
		if err != nil {
			return nil, err
		}
		for _, item := range order.GetItems() {
			err = s.DB.Select(&item.Options, fmt.Sprintf(
				`
				SELECT options.name,options.price 
				FROM order_item_option AS oio CROSS JOIN order_items
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

func (s *Service) CompleteOrder(id int32) error {
	tx, err := s.DB.Begin()
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

var Menu = []*mookiespb.Category{
	{
		Name: "Sandwiches",
		Items: []*mookiespb.Item{
			{Name: "LG Smoked Pulled Pork", Price: 495, Options: []*mookiespb.Option{
				{Name: "pickles", Price: 0, Selected: true},
				{Name: "bbq sauce", Price: 0, Selected: true},
				{Name: "white sauce", Price: 0, Selected: false},
				{Name: "ketchup", Price: 0, Selected: false},
				{Name: "mayo", Price: 0, Selected: false},
				{Name: "coleslaw", Price: 25, Selected: false},
				{Name: "tomato", Price: 25, Selected: false},
				{Name: "lettuce", Price: 25, Selected: false},
				{Name: "cheese", Price: 25, Selected: false},
				{Name: "onion", Price: 25, Selected: false},
			}},
			{Name: "RG Smoked Pulled Pork", Price: 395, Options: []*mookiespb.Option{
				{Name: "pickles", Price: 0, Selected: true},
				{Name: "bbq sauce", Price: 0, Selected: true},
				{Name: "white sauce", Price: 0, Selected: false},
				{Name: "ketchup", Price: 0, Selected: false},
				{Name: "mayo", Price: 0, Selected: false},
				{Name: "coleslaw", Price: 25, Selected: false},
				{Name: "tomato", Price: 25, Selected: false},
				{Name: "lettuce", Price: 25, Selected: false},
				{Name: "cheese", Price: 25, Selected: false},
				{Name: "onion", Price: 25, Selected: false},
			}},
			{Name: "LG Smoked Chicken Breast", Price: 495, Options: []*mookiespb.Option{
				{Name: "pickles", Price: 0, Selected: false},
				{Name: "bbq sauce", Price: 0, Selected: false},
				{Name: "white sauce", Price: 0, Selected: false},
				{Name: "ketchup", Price: 0, Selected: false},
				{Name: "mayo", Price: 0, Selected: true},
				{Name: "coleslaw", Price: 25, Selected: false},
				{Name: "tomato", Price: 25, Selected: true},
				{Name: "lettuce", Price: 25, Selected: true},
				{Name: "cheese", Price: 25, Selected: false},
				{Name: "onion", Price: 25, Selected: false},
			}},
			{Name: "RG Smoked Chicken Breast", Price: 395, Options: []*mookiespb.Option{
				{Name: "pickles", Price: 0, Selected: false},
				{Name: "bbq sauce", Price: 0, Selected: false},
				{Name: "white sauce", Price: 0, Selected: false},
				{Name: "ketchup", Price: 0, Selected: true},
				{Name: "mayo", Price: 0, Selected: true},
				{Name: "coleslaw", Price: 25, Selected: false},
				{Name: "tomato", Price: 25, Selected: true},
				{Name: "lettuce", Price: 25, Selected: true},
				{Name: "cheese", Price: 25, Selected: false},
				{Name: "onion", Price: 25, Selected: false},
			}},
			{Name: "'The Molly'", Price: 395, Options: []*mookiespb.Option{
				{Name: "pickles", Price: 0, Selected: true},
				{Name: "bbq sauce", Price: 0, Selected: true},
				{Name: "white sauce", Price: 0, Selected: true},
				{Name: "ketchup", Price: 0, Selected: false},
				{Name: "mayo", Price: 0, Selected: false},
				{Name: "coleslaw", Price: 25, Selected: true},
				{Name: "tomato", Price: 25, Selected: false},
				{Name: "lettuce", Price: 25, Selected: false},
				{Name: "cheese", Price: 25, Selected: false},
				{Name: "onion", Price: 25, Selected: false},
			}},
			{Name: "LG Hamburger", Price: 495, Options: []*mookiespb.Option{
				{Name: "pickles", Price: 0, Selected: true},
				{Name: "bbq sauce", Price: 0, Selected: true},
				{Name: "white sauce", Price: 0, Selected: false},
				{Name: "ketchup", Price: 0, Selected: false},
				{Name: "mayo", Price: 0, Selected: false},
				{Name: "coleslaw", Price: 25, Selected: false},
				{Name: "tomato", Price: 25, Selected: false},
				{Name: "lettuce", Price: 25, Selected: false},
			}},
			{Name: "Hamburger", Price: 395, Options: []*mookiespb.Option{
				{Name: "pickles", Price: 0, Selected: true},
				{Name: "bbq sauce", Price: 0, Selected: true},
				{Name: "white sauce", Price: 0, Selected: false},
				{Name: "ketchup", Price: 0, Selected: false},
				{Name: "mayo", Price: 0, Selected: false},
				{Name: "coleslaw", Price: 25, Selected: false},
				{Name: "tomato", Price: 25, Selected: false},
				{Name: "lettuce", Price: 25, Selected: false},
			}},
			{Name: "LG Cheeseburger", Price: 550, Options: []*mookiespb.Option{
				{Name: "pickles", Price: 0, Selected: true},
				{Name: "bbq sauce", Price: 0, Selected: true},
				{Name: "white sauce", Price: 0, Selected: false},
				{Name: "ketchup", Price: 0, Selected: false},
				{Name: "mayo", Price: 0, Selected: false},
				{Name: "coleslaw", Price: 25, Selected: false},
				{Name: "tomato", Price: 25, Selected: false},
				{Name: "lettuce", Price: 25, Selected: false},
			}},
			{Name: "Cheeseburger", Price: 425, Options: []*mookiespb.Option{
				{Name: "pickles", Price: 0, Selected: true},
				{Name: "bbq sauce", Price: 0, Selected: true},
				{Name: "white sauce", Price: 0, Selected: false},
				{Name: "ketchup", Price: 0, Selected: false},
				{Name: "mayo", Price: 0, Selected: false},
				{Name: "coleslaw", Price: 25, Selected: false},
				{Name: "tomato", Price: 25, Selected: false},
				{Name: "lettuce", Price: 25, Selected: false},
			}},
			{Name: "Grilled Cheese", Price: 300, Options: []*mookiespb.Option{
				{Name: "pickles", Price: 0, Selected: true},
				{Name: "bbq sauce", Price: 0, Selected: true},
				{Name: "white sauce", Price: 0, Selected: false},
				{Name: "ketchup", Price: 0, Selected: false},
				{Name: "mayo", Price: 0, Selected: false},
				{Name: "coleslaw", Price: 25, Selected: false},
				{Name: "tomato", Price: 25, Selected: false},
				{Name: "lettuce", Price: 25, Selected: false},
			}},
			{Name: "Pulled Pork Melt", Price: 395, Options: []*mookiespb.Option{
				{Name: "pickles", Price: 0, Selected: true},
				{Name: "bbq sauce", Price: 0, Selected: true},
				{Name: "white sauce", Price: 0, Selected: false},
				{Name: "ketchup", Price: 0, Selected: false},
				{Name: "mayo", Price: 0, Selected: false},
				{Name: "coleslaw", Price: 25, Selected: false},
				{Name: "tomato", Price: 25, Selected: false},
				{Name: "lettuce", Price: 25, Selected: false},
			}},
		},
	},
	{
		Name: "Plates",
		Items: []*mookiespb.Item{
			{Name: "Smoked Pulled Pork", Price: 990, Options: []*mookiespb.Option{
				{Name: "potato salad", Price: 0, Selected: false},
				{Name: "baked beans", Price: 0, Selected: false},
				{Name: "coleslaw", Price: 0, Selected: false},
				{Name: "chips", Price: 0, Selected: false},
				{Name: "fries", Price: 0, Selected: false},
				{Name: "baked potato", Price: 100, Selected: false},
			}},
			{Name: "RG Rib", Price: 995, Options: []*mookiespb.Option{
				{Name: "potato salad", Price: 0, Selected: false},
				{Name: "baked beans", Price: 0, Selected: false},
				{Name: "coleslaw", Price: 0, Selected: false},
				{Name: "chips", Price: 0, Selected: false},
				{Name: "fries", Price: 0, Selected: false},
				{Name: "baked potato", Price: 100, Selected: false},
			}},
			{Name: "½ Smoked Chicken", Price: 995, Options: []*mookiespb.Option{
				{Name: "potato salad", Price: 0, Selected: false},
				{Name: "baked beans", Price: 0, Selected: false},
				{Name: "coleslaw", Price: 0, Selected: false},
				{Name: "chips", Price: 0, Selected: false},
				{Name: "fries", Price: 0, Selected: false},
				{Name: "baked potato", Price: 100, Selected: false},
			}},
			{Name: "Smoked Chicken Breast", Price: 725, Options: []*mookiespb.Option{
				{Name: "potato salad", Price: 0, Selected: false},
				{Name: "baked beans", Price: 0, Selected: false},
				{Name: "coleslaw", Price: 0, Selected: false},
				{Name: "chips", Price: 0, Selected: false},
				{Name: "fries", Price: 0, Selected: false},
				{Name: "baked potato", Price: 100, Selected: false},
			}},
			{Name: "Smoked Wing (8 wings)", Price: 999, Options: []*mookiespb.Option{
				{Name: "potato salad", Price: 0, Selected: false},
				{Name: "baked beans", Price: 0, Selected: false},
				{Name: "coleslaw", Price: 0, Selected: false},
				{Name: "chips", Price: 0, Selected: false},
				{Name: "fries", Price: 0, Selected: false},
				{Name: "baked potato", Price: 100, Selected: false},
			}},
			{Name: "BBQ Loaded Nachos", Price: 875, Options: []*mookiespb.Option{
				{Name: "cheddar cheese", Price: 0, Selected: true},
				{Name: "green onions", Price: 0, Selected: true},
				{Name: "jalapeños", Price: 0, Selected: true},
				{Name: "bbq sauce", Price: 0, Selected: false},
				{Name: "white sauce", Price: 0, Selected: false},
			}},
			{Name: "Chicken Loaded Nachos", Price: 875, Options: []*mookiespb.Option{
				{Name: "cheddar cheese", Price: 0, Selected: true},
				{Name: "green onions", Price: 0, Selected: true},
				{Name: "jalapeños", Price: 0, Selected: true},
				{Name: "bbq sauce", Price: 0, Selected: false},
				{Name: "white sauce", Price: 0, Selected: false},
			}},
		},
	},
	{
		Name: "Baskets",
		Items: []*mookiespb.Item{
			{Name: "Smoked Wing", Price: 500, Options: []*mookiespb.Option{
				{Name: "potato salad", Price: 0, Selected: false},
				{Name: "baked beans", Price: 0, Selected: false},
				{Name: "coleslaw", Price: 0, Selected: false},
				{Name: "chips", Price: 0, Selected: false},
				{Name: "fries", Price: 0, Selected: false},
				{Name: "baked potato", Price: 100, Selected: false},
			}},
			{Name: "Rib", Price: 500, Options: []*mookiespb.Option{
				{Name: "potato salad", Price: 0, Selected: false},
				{Name: "baked beans", Price: 0, Selected: false},
				{Name: "coleslaw", Price: 0, Selected: false},
				{Name: "chips", Price: 0, Selected: false},
				{Name: "fries", Price: 0, Selected: false},
				{Name: "baked potato", Price: 100, Selected: false},
			}},
		},
	},
	{
		Name: "Potatoes",
		Items: []*mookiespb.Item{
			{Name: "Loaded Pork", Price: 899, Options: []*mookiespb.Option{
				{Name: "cheese", Price: 0, Selected: true},
				{Name: "bacon bits", Price: 0, Selected: true},
				{Name: "green onions", Price: 0, Selected: true},
				{Name: "butter", Price: 0, Selected: true},
				{Name: "sour cream", Price: 0, Selected: true},
			}},
			{Name: "Loaded Chicken", Price: 899, Options: []*mookiespb.Option{
				{Name: "cheese", Price: 0, Selected: true},
				{Name: "bacon bits", Price: 0, Selected: true},
				{Name: "green onions", Price: 0, Selected: true},
				{Name: "butter", Price: 0, Selected: true},
				{Name: "sour cream", Price: 0, Selected: true},
			}},
			{Name: "Loaded Potato (no meat)", Price: 699, Options: []*mookiespb.Option{
				{Name: "cheese", Price: 0, Selected: true},
				{Name: "bacon bits", Price: 0, Selected: true},
				{Name: "green onions", Price: 0, Selected: true},
				{Name: "butter", Price: 0, Selected: true},
				{Name: "sour cream", Price: 0, Selected: true},
			}},
			{Name: "(Pork) Smothered and Covered Fries", Price: 899, Options: []*mookiespb.Option{
				{Name: "cheese", Price: 0, Selected: true},
				{Name: "bacon bits", Price: 0, Selected: true},
				{Name: "green onions", Price: 0, Selected: true},
				{Name: "butter", Price: 0, Selected: true},
				{Name: "sour cream", Price: 0, Selected: true},
			}},
			{Name: "(Chicken) Smothered and Covered Fries", Price: 899, Options: []*mookiespb.Option{
				{Name: "cheese", Price: 0, Selected: true},
				{Name: "bacon bits", Price: 0, Selected: true},
				{Name: "green onions", Price: 0, Selected: true},
				{Name: "butter", Price: 0, Selected: true},
				{Name: "sour cream", Price: 0, Selected: true},
			}},
		},
	},
	{
		Name: "Sides",
		Items: []*mookiespb.Item{
			{Name: "SM Cole Slaw", Price: 300},
			{Name: "LG Cole Slaw", Price: 600},
			{Name: "SM Baked Beans", Price: 300},
			{Name: "LG Baked Beans", Price: 600},
			{Name: "SM Potato Salad", Price: 300},
			{Name: "LG Potato Salad", Price: 600},
			{Name: "Plain Chips", Price: 100},
			{Name: "Fries", Price: 300},
		},
	},
	{
		Name: "Drinks",
		Items: []*mookiespb.Item{
			{Name: "Canned Drink", Price: 100},
			{Name: "Bottled Water", Price: 150},
		},
	},
	{
		Name: "Desserts",
		Items: []*mookiespb.Item{
			{Name: "Oreo Dream", Price: 350},
			{Name: "Lemon Delight", Price: 350},
			{Name: "Strawberry Pizza", Price: 350},
			{Name: "Whole Dessert", Price: 3000},
			{Name: "SM Banana Pudding", Price: 350},
			{Name: "LG Banana Pudding", Price: 700},
		},
	},
	{
		Name: "Sauces",
		Items: []*mookiespb.Item{
			{Name: "Extra Ranch", Price: 50},
			{Name: "Extra White", Price: 50},
			{Name: "Extra Buffalo Wing", Price: 50},
			{Name: "Extra BBQ", Price: 50},
		},
	},
}
