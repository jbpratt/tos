package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	mookiespb "github.com/jbpratt78/mookies-tos/protofiles"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

var (
	listen  = flag.String("listen", ":50051", "listen address")
	dbp     = flag.String("database", "./mookies.db", "database to use")
	reqChan = make(chan *mookiespb.Order, 1024)
)

type server struct {
	db     *sqlx.DB
	orders []*mookiespb.Order
	menu   *mookiespb.Menu
}

func (s *server) GetMenu(ctx context.Context, empty *mookiespb.Empty) (*mookiespb.Menu, error) {
	log.Println("Menu function was invoked")
	res := s.menu
	return res, nil
}

func (s *server) SubmitOrder(ctx context.Context,
	req *mookiespb.SubmitOrderRequest) (*mookiespb.SubmitOrderResponse, error) {

	log.Println("An order was received")
	o := req.GetOrder()
	// expecting it to be right id
	o.Id = int32(len(s.orders))
	o.Status = "active"

	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	_, err = tx.Exec(
		"INSERT INTO orders (name, total, status, time_ordered, time_complete) VALUES (?, ?, ?, ?, ?)",
		o.GetName(), o.GetTotal(), o.GetStatus(), time.Now().UTC().String(), "")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, item := range o.GetItems() {
		_, err := tx.Exec(
			"INSERT INTO order_item (itemid, orderid) VALUES (?, ?)",
			item.GetId(), o.GetId())
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	res := &mookiespb.SubmitOrderResponse{
		Result: "Order has been placed..",
	}

	reqChan <- o

	return res, s.LoadData()
}

func (*server) SubscribeToOrders(req *mookiespb.SubscribeToOrderRequest,
	stream mookiespb.OrderService_SubscribeToOrdersServer) error {

	log.Printf("SubscribeToOrders function was invoked with %v\n", req)
	for {
		res := <-reqChan
		log.Printf("Sending order to client: %v\n", res)
		err := stream.Send(res)
		if err != nil {
			return err
		}
		//time.Sleep(time.Millisecond * 1000)
	}
}

func (s *server) CompleteOrder(ctx context.Context,
	req *mookiespb.CompleteOrderRequest) (*mookiespb.CompleteOrderResponse, error) {

	log.Printf("CompleteOrder function was invoked with %v\n", req)
	// update order to be complete
	for _, o := range s.orders {
		if req.GetId() == o.GetId() {
			o.Status = "complete"
		}
	}

	// update query at req.GetId()
	_, err := s.db.Exec("UPDATE orders SET status = 'complete' WHERE id = ?", req.GetId())
	if err != nil {
		return nil, err
	}
	res := &mookiespb.CompleteOrderResponse{
		Result: "Order marked as complete",
	}
	return res, s.LoadData()
}

func (s *server) ActiveOrders(ctx context.Context, empty *mookiespb.Empty) (*mookiespb.OrdersResponse, error) {
	log.Println("Active orders function was invoked")
	res := &mookiespb.OrdersResponse{
		Orders: s.orders,
	}
	return res, nil
}

func (s *server) LoadData() error {
	var categories []*mookiespb.Category
	menu := &mookiespb.Menu{
		Categories: categories,
	}
	// get menu
	err := s.db.Select(&menu.Categories, "SELECT * from categories")
	for _, c := range menu.GetCategories() {
		err = s.db.Select(&c.Items, fmt.Sprintf("SELECT * FROM items WHERE category_id = %v", c.GetId()))
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	s.menu = menu
	// get all orders
	var orders []*mookiespb.Order
	err = s.db.Select(&orders, "SELECT * FROM orders WHERE status = 'active'")
	if err != nil {
		return err
	}

	// TODO: query items along with orders

	s.orders = orders
	return nil
}

func NewServer(db *sqlx.DB) (*server, error) {
	server := &server{db: db}
	//err := server.seedData()
	err := server.LoadData()
	if err != nil {
		return nil, err
	}
	return server, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	log.Printf("Listening on %q...\n", *listen)

	db, err := sqlx.Open("sqlite3", *dbp)
	if err != nil {
		log.Fatalf("Failed to open DB: %v\n", err)
	}
	defer db.Close()
	server, err := NewServer(db)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	mookiespb.RegisterMenuServiceServer(s, server)
	mookiespb.RegisterOrderServiceServer(s, server)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) seedData() error {
	data := []*mookiespb.Category{
		{
			Name: "Sandwich",
			Items: []*mookiespb.Item{
				{Name: "LG Smoked Pulled Pork", Price: 495, Options: []*mookiespb.Option{
					{Name: "pickles", Price: 0, ByDefault: true},
					{Name: "bbq sauce", Price: 0, ByDefault: true},
					{Name: "white sauce", Price: 0, ByDefault: false},
					{Name: "ketchup", Price: 0, ByDefault: false},
					{Name: "mayo", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 25, ByDefault: false},
					{Name: "tomato", Price: 25, ByDefault: false},
					{Name: "lettuce", Price: 25, ByDefault: false},
					{Name: "cheese", Price: 25, ByDefault: false},
					{Name: "onion", Price: 25, ByDefault: false},
				}},
				{Name: "RG Smoked Pulled Pork", Price: 395, Options: []*mookiespb.Option{
					{Name: "pickles", Price: 0, ByDefault: true},
					{Name: "bbq sauce", Price: 0, ByDefault: true},
					{Name: "white sauce", Price: 0, ByDefault: false},
					{Name: "ketchup", Price: 0, ByDefault: false},
					{Name: "mayo", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 25, ByDefault: false},
					{Name: "tomato", Price: 25, ByDefault: false},
					{Name: "lettuce", Price: 25, ByDefault: false},
					{Name: "cheese", Price: 25, ByDefault: false},
					{Name: "onion", Price: 25, ByDefault: false},
				}},
				{Name: "LG Smoked Chicken Breast", Price: 495, Options: []*mookiespb.Option{
					{Name: "pickles", Price: 0, ByDefault: false},
					{Name: "bbq sauce", Price: 0, ByDefault: false},
					{Name: "white sauce", Price: 0, ByDefault: false},
					{Name: "ketchup", Price: 0, ByDefault: false},
					{Name: "mayo", Price: 0, ByDefault: true},
					{Name: "coleslaw", Price: 25, ByDefault: false},
					{Name: "tomato", Price: 25, ByDefault: true},
					{Name: "lettuce", Price: 25, ByDefault: true},
					{Name: "cheese", Price: 25, ByDefault: false},
					{Name: "onion", Price: 25, ByDefault: false},
				}},
				{Name: "RG Smoked Chicken Breast", Price: 395, Options: []*mookiespb.Option{
					{Name: "pickles", Price: 0, ByDefault: false},
					{Name: "bbq sauce", Price: 0, ByDefault: false},
					{Name: "white sauce", Price: 0, ByDefault: false},
					{Name: "ketchup", Price: 0, ByDefault: true},
					{Name: "mayo", Price: 0, ByDefault: true},
					{Name: "coleslaw", Price: 25, ByDefault: false},
					{Name: "tomato", Price: 25, ByDefault: true},
					{Name: "lettuce", Price: 25, ByDefault: true},
					{Name: "cheese", Price: 25, ByDefault: false},
					{Name: "onion", Price: 25, ByDefault: false},
				}},
				{Name: "'The Molly'", Price: 395, Options: []*mookiespb.Option{
					{Name: "pickles", Price: 0, ByDefault: true},
					{Name: "bbq sauce", Price: 0, ByDefault: true},
					{Name: "white sauce", Price: 0, ByDefault: true},
					{Name: "ketchup", Price: 0, ByDefault: false},
					{Name: "mayo", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 25, ByDefault: true},
					{Name: "tomato", Price: 25, ByDefault: false},
					{Name: "lettuce", Price: 25, ByDefault: false},
					{Name: "cheese", Price: 25, ByDefault: false},
					{Name: "onion", Price: 25, ByDefault: false},
				}},
				{Name: "LG Hamburger", Price: 495, Options: []*mookiespb.Option{
					{Name: "pickles", Price: 0, ByDefault: true},
					{Name: "bbq sauce", Price: 0, ByDefault: true},
					{Name: "white sauce", Price: 0, ByDefault: false},
					{Name: "ketchup", Price: 0, ByDefault: false},
					{Name: "mayo", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 25, ByDefault: false},
					{Name: "tomato", Price: 25, ByDefault: false},
					{Name: "lettuce", Price: 25, ByDefault: false},
				}},
				{Name: "Hamburger", Price: 395, Options: []*mookiespb.Option{
					{Name: "pickles", Price: 0, ByDefault: true},
					{Name: "bbq sauce", Price: 0, ByDefault: true},
					{Name: "white sauce", Price: 0, ByDefault: false},
					{Name: "ketchup", Price: 0, ByDefault: false},
					{Name: "mayo", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 25, ByDefault: false},
					{Name: "tomato", Price: 25, ByDefault: false},
					{Name: "lettuce", Price: 25, ByDefault: false},
				}},
				{Name: "LG Cheeseburger", Price: 550, Options: []*mookiespb.Option{
					{Name: "pickles", Price: 0, ByDefault: true},
					{Name: "bbq sauce", Price: 0, ByDefault: true},
					{Name: "white sauce", Price: 0, ByDefault: false},
					{Name: "ketchup", Price: 0, ByDefault: false},
					{Name: "mayo", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 25, ByDefault: false},
					{Name: "tomato", Price: 25, ByDefault: false},
					{Name: "lettuce", Price: 25, ByDefault: false},
				}},
				{Name: "Cheeseburger", Price: 425, Options: []*mookiespb.Option{
					{Name: "pickles", Price: 0, ByDefault: true},
					{Name: "bbq sauce", Price: 0, ByDefault: true},
					{Name: "white sauce", Price: 0, ByDefault: false},
					{Name: "ketchup", Price: 0, ByDefault: false},
					{Name: "mayo", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 25, ByDefault: false},
					{Name: "tomato", Price: 25, ByDefault: false},
					{Name: "lettuce", Price: 25, ByDefault: false},
				}},
				{Name: "Grilled Cheese", Price: 300, Options: []*mookiespb.Option{
					{Name: "pickles", Price: 0, ByDefault: true},
					{Name: "bbq sauce", Price: 0, ByDefault: true},
					{Name: "white sauce", Price: 0, ByDefault: false},
					{Name: "ketchup", Price: 0, ByDefault: false},
					{Name: "mayo", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 25, ByDefault: false},
					{Name: "tomato", Price: 25, ByDefault: false},
					{Name: "lettuce", Price: 25, ByDefault: false},
				}},
				{Name: "Pulled Pork Melt", Price: 395, Options: []*mookiespb.Option{
					{Name: "pickles", Price: 0, ByDefault: true},
					{Name: "bbq sauce", Price: 0, ByDefault: true},
					{Name: "white sauce", Price: 0, ByDefault: false},
					{Name: "ketchup", Price: 0, ByDefault: false},
					{Name: "mayo", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 25, ByDefault: false},
					{Name: "tomato", Price: 25, ByDefault: false},
					{Name: "lettuce", Price: 25, ByDefault: false},
				}},
			},
		},
		{
			Name: "Plates",
			Items: []*mookiespb.Item{
				{Name: "Smoked Pulled Pork", Price: 990, Options: []*mookiespb.Option{
					{Name: "potato salad", Price: 0, ByDefault: false},
					{Name: "baked beans", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 0, ByDefault: false},
					{Name: "chips", Price: 0, ByDefault: false},
					{Name: "fries", Price: 0, ByDefault: false},
					{Name: "baked potato", Price: 100, ByDefault: false},
				}},
				{Name: "RG Rib", Price: 995, Options: []*mookiespb.Option{
					{Name: "potato salad", Price: 0, ByDefault: false},
					{Name: "baked beans", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 0, ByDefault: false},
					{Name: "chips", Price: 0, ByDefault: false},
					{Name: "fries", Price: 0, ByDefault: false},
					{Name: "baked potato", Price: 100, ByDefault: false},
				}},
				{Name: "½ Smoked Chicken", Price: 995, Options: []*mookiespb.Option{
					{Name: "potato salad", Price: 0, ByDefault: false},
					{Name: "baked beans", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 0, ByDefault: false},
					{Name: "chips", Price: 0, ByDefault: false},
					{Name: "fries", Price: 0, ByDefault: false},
					{Name: "baked potato", Price: 100, ByDefault: false},
				}},
				{Name: "Smoked Chicken Breast", Price: 725, Options: []*mookiespb.Option{
					{Name: "potato salad", Price: 0, ByDefault: false},
					{Name: "baked beans", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 0, ByDefault: false},
					{Name: "chips", Price: 0, ByDefault: false},
					{Name: "fries", Price: 0, ByDefault: false},
					{Name: "baked potato", Price: 100, ByDefault: false},
				}},
				{Name: "Smoked Wing (8 wings)", Price: 999, Options: []*mookiespb.Option{
					{Name: "potato salad", Price: 0, ByDefault: false},
					{Name: "baked beans", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 0, ByDefault: false},
					{Name: "chips", Price: 0, ByDefault: false},
					{Name: "fries", Price: 0, ByDefault: false},
					{Name: "baked potato", Price: 100, ByDefault: false},
				}},
				{Name: "BBQ Loaded Nachos", Price: 875, Options: []*mookiespb.Option{
					{Name: "cheddar cheese", Price: 0, ByDefault: true},
					{Name: "green onions", Price: 0, ByDefault: true},
					{Name: "jalapeños", Price: 0, ByDefault: true},
					{Name: "bbq sauce", Price: 0, ByDefault: false},
					{Name: "white sauce", Price: 0, ByDefault: false},
				}},
				{Name: "Chicken Loaded Nachos", Price: 875, Options: []*mookiespb.Option{
					{Name: "cheddar cheese", Price: 0, ByDefault: true},
					{Name: "green onions", Price: 0, ByDefault: true},
					{Name: "jalapeños", Price: 0, ByDefault: true},
					{Name: "bbq sauce", Price: 0, ByDefault: false},
					{Name: "white sauce", Price: 0, ByDefault: false},
				}},
			},
		},
		{
			Name: "Baskets",
			Items: []*mookiespb.Item{
				{Name: "Smoked Wing", Price: 500, Options: []*mookiespb.Option{
					{Name: "potato salad", Price: 0, ByDefault: false},
					{Name: "baked beans", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 0, ByDefault: false},
					{Name: "chips", Price: 0, ByDefault: false},
					{Name: "fries", Price: 0, ByDefault: false},
					{Name: "baked potato", Price: 100, ByDefault: false},
				}},
				{Name: "Rib", Price: 500, Options: []*mookiespb.Option{
					{Name: "potato salad", Price: 0, ByDefault: false},
					{Name: "baked beans", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 0, ByDefault: false},
					{Name: "chips", Price: 0, ByDefault: false},
					{Name: "fries", Price: 0, ByDefault: false},
					{Name: "baked potato", Price: 100, ByDefault: false},
				}},
			},
		},
		{
			Name: "Potatoes",
			Items: []*mookiespb.Item{
				{Name: "Loaded Pork", Price: 899, Options: []*mookiespb.Option{
					{Name: "cheese", Price: 0, ByDefault: true},
					{Name: "bacon bits", Price: 0, ByDefault: true},
					{Name: "green onions", Price: 0, ByDefault: true},
					{Name: "butter", Price: 0, ByDefault: true},
					{Name: "sour cream", Price: 0, ByDefault: true},
				}},
				{Name: "Loaded Chicken", Price: 899, Options: []*mookiespb.Option{
					{Name: "cheese", Price: 0, ByDefault: true},
					{Name: "bacon bits", Price: 0, ByDefault: true},
					{Name: "green onions", Price: 0, ByDefault: true},
					{Name: "butter", Price: 0, ByDefault: true},
					{Name: "sour cream", Price: 0, ByDefault: true},
				}},
				{Name: "Loaded Potato (no meat)", Price: 699, Options: []*mookiespb.Option{
					{Name: "cheese", Price: 0, ByDefault: true},
					{Name: "bacon bits", Price: 0, ByDefault: true},
					{Name: "green onions", Price: 0, ByDefault: true},
					{Name: "butter", Price: 0, ByDefault: true},
					{Name: "sour cream", Price: 0, ByDefault: true},
				}},
				{Name: "(Pork) Smothered and Covered Fries", Price: 899, Options: []*mookiespb.Option{
					{Name: "cheese", Price: 0, ByDefault: true},
					{Name: "bacon bits", Price: 0, ByDefault: true},
					{Name: "green onions", Price: 0, ByDefault: true},
					{Name: "butter", Price: 0, ByDefault: true},
					{Name: "sour cream", Price: 0, ByDefault: true},
				}},
				{Name: "(Chicken) Smothered and Covered Fries", Price: 899, Options: []*mookiespb.Option{
					{Name: "cheese", Price: 0, ByDefault: true},
					{Name: "bacon bits", Price: 0, ByDefault: true},
					{Name: "green onions", Price: 0, ByDefault: true},
					{Name: "butter", Price: 0, ByDefault: true},
					{Name: "sour cream", Price: 0, ByDefault: true},
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

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	for i, category := range data {
		_, err := tx.Exec("INSERT INTO categories (name) VALUES (?)", category.GetName())
		if err != nil {
			tx.Rollback()
			return err
		}
		for x, item := range category.GetItems() {
			_, err = tx.Exec(
				"INSERT INTO items (name, price, category_id) VALUES (?,?,?)",
				item.GetName(), item.GetPrice(), i+1)
			if err != nil {
				tx.Rollback()
				return err
			}
			for o, option := range item.GetOptions() {
				_, err = tx.Exec(
					"INSERT INTO options (name, price, by_default) VALUES (?,?,?)",
					option.GetName(), option.GetPrice(), option.GetByDefault())
				if err != nil {
					tx.Rollback()
					return err
				}
				_, err = tx.Exec(
					"INSERT INTO item_options (item_id, option_id) VALUES (?,?)",
					x+1, o+1)
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
	/*for _, c := range data {
		t := "INSERT INTO categories (name) VALUES ('%s');"
		cmd := fmt.Sprintf(t, c.GetName())
		fmt.Println(cmd)
		//res := tx.MustExec(cmd)
	}

	for x, c := range data {
		for _, i := range c.GetItems() {
			t := "INSERT INTO items (name, price, category_id) VALUES ('%v', '%v', '%v');"
			cmd := fmt.Sprintf(t, i.GetName(), i.GetPrice(), x)
			fmt.Println(cmd)
		}
	}*/

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
