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

func seedData() {
	data := []*mookiespb.Category{
		{
			Name: "Sandwich",
			Items: []*mookiespb.Item{
				{Name: "Large Smoked Pulled Pork", Price: 495, Options: []*mookiespb.Option{
					{Name: "pickles", Price: 0, ByDefault: true},
					{Name: "bbq sauce", Price: 0, ByDefault: true},
					{Name: "white sauce", Price: 0, ByDefault: false},
					{Name: "ketchup", Price: 0, ByDefault: false},
					{Name: "mayo", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 25, ByDefault: false},
					{Name: "tomato", Price: 25, ByDefault: false},
					{Name: "lettuce", Price: 25, ByDefault: false},
					{Name: "cheese", Price: 25, ByDefault: false},
				}},
				{Name: "Regular Smoked Pulled Pork", Price: 395, Options: []*mookiespb.Option{
					{Name: "pickles", Price: 0, ByDefault: true},
					{Name: "bbq sauce", Price: 0, ByDefault: true},
					{Name: "white sauce", Price: 0, ByDefault: false},
					{Name: "ketchup", Price: 0, ByDefault: false},
					{Name: "mayo", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 25, ByDefault: false},
					{Name: "tomato", Price: 25, ByDefault: false},
					{Name: "lettuce", Price: 25, ByDefault: false},
				}},
				{Name: "Large Smoked Chicken Breast", Price: 495, Options: []*mookiespb.Option{
					{Name: "pickles", Price: 0, ByDefault: true},
					{Name: "bbq sauce", Price: 0, ByDefault: true},
					{Name: "white sauce", Price: 0, ByDefault: false},
					{Name: "ketchup", Price: 0, ByDefault: false},
					{Name: "mayo", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 25, ByDefault: false},
					{Name: "tomato", Price: 25, ByDefault: false},
					{Name: "lettuce", Price: 25, ByDefault: false},
				}},
				{Name: "Regular Smoked Chicken Breast", Price: 395, Options: []*mookiespb.Option{
					{Name: "pickles", Price: 0, ByDefault: true},
					{Name: "bbq sauce", Price: 0, ByDefault: true},
					{Name: "white sauce", Price: 0, ByDefault: false},
					{Name: "ketchup", Price: 0, ByDefault: false},
					{Name: "mayo", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 25, ByDefault: false},
					{Name: "tomato", Price: 25, ByDefault: false},
					{Name: "lettuce", Price: 25, ByDefault: false},
				}},
				{Name: "'The Molly'", Price: 395, Options: []*mookiespb.Option{
					{Name: "pickles", Price: 0, ByDefault: true},
					{Name: "bbq sauce", Price: 0, ByDefault: true},
					{Name: "white sauce", Price: 0, ByDefault: false},
					{Name: "ketchup", Price: 0, ByDefault: false},
					{Name: "mayo", Price: 0, ByDefault: false},
					{Name: "coleslaw", Price: 25, ByDefault: false},
					{Name: "tomato", Price: 25, ByDefault: false},
					{Name: "lettuce", Price: 25, ByDefault: false},
				}},
				{Name: "Large Hamburger", Price: 495, Options: []*mookiespb.Option{
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
				{Name: "Large Cheeseburger", Price: 550, Options: []*mookiespb.Option{
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
					{Name: "coleslaw"}, {Name: "pickles"}, {Name: "lettuce"}, {Name: "tomato"}, {Name: "mayo"},
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
				{Name: "Smoked Pulled Pork", Price: 990, Options: []*mookiespb.Option{}},
				{Name: "Regular Rib", Price: 995},
				{Name: "1/2 Smoked Chicken", Price: 995},
				{Name: "Smoked Chicken Breast", Price: 725},
				{Name: "Smoked Wing (8 wings)", Price: 999},
				{Name: "Loaded Nachos (BBQ or Chicken)", Price: 875},
			},
		},
		{
			Name: "Baskets",
			Items: []*mookiespb.Item{
				{Name: "Smoked Wing", Price: 500},
				{Name: "Rib", Price: 500},
			},
		},
		{
			Name: "Potatoes",
			Items: []*mookiespb.Item{
				{Name: "Loaded Pork", Price: 899},
				{Name: "Loaded Chicken", Price: 899},
				{Name: "Loaded Potato (no meat)", Price: 699},
				{Name: "Smothered and Covered Fries", Price: 899},
			},
		},
		{
			Name: "Sides",
			Items: []*mookiespb.Item{
				{Name: "Small Cole Slaw", Price: 300},
				{Name: "Large Cole Slaw", Price: 600},
				{Name: "Small Baked Beans", Price: 300},
				{Name: "Large Baked Beans", Price: 600},
				{Name: "Small Potato Salad", Price: 300},
				{Name: "Large Potato Salad", Price: 600},
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
				{Name: "Small Banana Pudding", Price: 350},
				{Name: "Large Banana Pudding", Price: 700},
			},
		},
		{
			Name: "Sauces",
			Items: []*mookiespb.Item{
				{Name: "Extra sauce", Price: 50},
			},
		},
	}

	for _, c := range data {
		t := "INSERT INTO categories (name) VALUES ('%s');"
		cmd := fmt.Sprintf(t, c.GetName())
		fmt.Println(cmd)
		//res := s.db.MustExec(cmd)
	}

	for x, c := range data {
		for _, i := range c.GetItems() {
			t := "INSERT INTO items (name, price, category_id) VALUES ('%v', '%v', '%v');"
			cmd := fmt.Sprintf(t, i.GetName(), i.GetPrice(), x)
			fmt.Println(cmd)
		}
	}
}
