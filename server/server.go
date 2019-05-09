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
				{Name: "Large Smoked Pulled Pork", Id: 1, Price: 495, CategoryID: 1, Options: []*mookiespb.Option{}},
				{Name: "Regular Smoked Pulled Pork", Id: 2, Price: 395, CategoryID: 1, Options: []*mookiespb.Option{}},
				{Name: "Large Smoked Chicken Breast", Id: 3, Price: 495, CategoryID: 1, Options: []*mookiespb.Option{}},
				{Name: "Regular Smoked Chicken Breast", Id: 4, Price: 395, CategoryID: 1, Options: []*mookiespb.Option{}},
				{Name: "'The Molly'", Id: 5, Price: 395, CategoryID: 1, Options: []*mookiespb.Option{}},
				{Name: "Large Hamburger", Id: 6, Price: 495, CategoryID: 1, Options: []*mookiespb.Option{}},
				{Name: "Hamburger", Id: 7, Price: 395, CategoryID: 1, Options: []*mookiespb.Option{}},
				{Name: "Large Cheeseburger", Id: 8, Price: 550, CategoryID: 1, Options: []*mookiespb.Option{}},
				{Name: "Cheeseburger", Id: 9, Price: 425, CategoryID: 1, Options: []*mookiespb.Option{}},
				{Name: "Grilled Cheese", Id: 10, Price: 300, CategoryID: 1, Options: []*mookiespb.Option{}},
				{Name: "Pulled Pork Melt", Id: 11, Price: 395, CategoryID: 1, Options: []*mookiespb.Option{}},
			},
		},
		{
			Name: "Plates",
			Items: []*mookiespb.Item{
				{Name: "Smoked Pulled Pork", Id: 12, Price: 990, CategoryID: 2},
				{Name: "Regular Rib", Id: 13, Price: 995, CategoryID: 2},
				{Name: "1/2 Smoked Chicken", Id: 14, Price: 995, CategoryID: 2},
				{Name: "Smoked Chicken Breast", Id: 15, Price: 725, CategoryID: 2},
				{Name: "Smoked Wing (8 wings)", Id: 16, Price: 999, CategoryID: 2},
				{Name: "Loaded Nachos (BBQ or Chicken)", Id: 17, Price: 875, CategoryID: 2},
			},
		},
		{
			Name: "Baskets",
			Items: []*mookiespb.Item{
				{Name: "Smoked Wing", Id: 18, Price: 500, CategoryID: 3},
				{Name: "Rib", Id: 19, Price: 500, CategoryID: 3},
			},
		},
		{
			Name: "Potatoes",
			Items: []*mookiespb.Item{
				{Name: "Loaded Pork", Id: 20, Price: 899, CategoryID: 4},
				{Name: "Loaded Chicken", Id: 21, Price: 899, CategoryID: 4},
				{Name: "Loaded Potato (no meat)", Id: 22, Price: 699, CategoryID: 4},
				{Name: "Smothered and Covered Fries", Id: 23, Price: 899, CategoryID: 4},
			},
		},
		{
			Name: "Sides",
			Items: []*mookiespb.Item{
				{Name: "Small Cole Slaw", Id: 24, Price: 300, CategoryID: 5},
				{Name: "Large Cole Slaw", Id: 25, Price: 600, CategoryID: 5},
				{Name: "Small Baked Beans", Id: 26, Price: 300, CategoryID: 5},
				{Name: "Large Baked Beans", Id: 27, Price: 600, CategoryID: 5},
				{Name: "Small Potato Salad", Id: 28, Price: 300, CategoryID: 5},
				{Name: "Large Potato Salad", Id: 29, Price: 600, CategoryID: 5},
				{Name: "Plain Chips", Id: 30, Price: 100, CategoryID: 5},
				{Name: "Fries", Id: 31, Price: 300, CategoryID: 5},
			},
		},
		{
			Name: "Drinks",
			Items: []*mookiespb.Item{
				{Name: "Canned Drink", Id: 32, Price: 100, CategoryID: 6},
				{Name: "Bottled Water", Id: 33, Price: 150, CategoryID: 6},
			},
		},
		{
			Name: "Desserts",
			Items: []*mookiespb.Item{
				{Name: "Oreo Dream", Id: 34, Price: 350, CategoryID: 7},
				{Name: "Lemon Delight", Id: 35, Price: 350, CategoryID: 7},
				{Name: "Strawberry Pizza", Id: 36, Price: 350, CategoryID: 7},
				{Name: "Whole Dessert", Id: 37, Price: 3000, CategoryID: 7},
				{Name: "Small Banana Pudding", Id: 38, Price: 350, CategoryID: 7},
				{Name: "Large Banana Pudding", Id: 39, Price: 700, CategoryID: 7},
			},
		},
		{
			Name: "Sauces",
			Items: []*mookiespb.Item{
				{Name: "Extra sauce", Id: 40, Price: 50, CategoryID: 8},
			},
		},
	}

	for _, c := range data {
		t := "INSERT INTO categories (name) VALUES ('%s');"
		cmd := fmt.Sprintf(t, c.GetName())
		fmt.Println(cmd)
		//res := s.db.MustExec(cmd)
	}

	for _, c := range data {
		for _, i := range c.GetItems() {
			t := "INSERT INTO items (name, price, category_id) VALUES ('%v', '%v', '%v');"
			cmd := fmt.Sprintf(t, i.GetName(), i.GetPrice(), i.GetCategoryID())
			fmt.Println(cmd)
		}
	}
}
