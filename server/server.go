package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	mookiespb "github.com/jbpratt78/mookies-tos/protofiles"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

var (
	listen  = flag.String("listen", ":50051", "listen address")
	dbp     = flag.String("database", "./mookies.db", "database to use")
	reqChan = make(chan *mookiespb.Order)
)

var Orders []*mookiespb.Order

type server struct {
	db *sqlx.DB
}

func (s *server) GetMenu(ctx context.Context, empty *empty.Empty) (*mookiespb.Menu, error) {
	log.Println("Menu function was invoked")
	var categories []*mookiespb.Category
	err := s.db.Select(&categories, "SELECT * from categories")
	for _, c := range categories {
		err = s.db.Select(&c.Items, fmt.Sprintf("SELECT * FROM items WHERE category_id = %v", c.GetId()))
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	res := &mookiespb.Menu{
		Categories: categories,
	}
	return res, nil
}

func (*server) SubmitOrder(ctx context.Context,
	req *mookiespb.SubmitOrderRequest) (*mookiespb.SubmitOrderResponse, error) {

	log.Println("An order was received")
	o := req.GetOrder()
	o.Id = int32(len(Orders) + 1)
	o.Status = mookiespb.Order_ACTIVE
	Orders = append(Orders, o)
	res := &mookiespb.SubmitOrderResponse{
		Result: "Order has been placed..",
	}
	go func() { reqChan <- o }()

	return res, nil
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
		time.Sleep(time.Millisecond * 1000)
	}
}

func (*server) CompleteOrder(ctx context.Context,
	req *mookiespb.CompleteOrderRequest) (*mookiespb.CompleteOrderResponse, error) {

	log.Printf("CompleteOrder function was invoked with %v\n", req)
	for _, o := range Orders {
		if req.GetId() == o.GetId() {
			o.Status = mookiespb.Order_COMPLETE
		}
	}
	res := &mookiespb.CompleteOrderResponse{
		Result: "Order marked as complete",
	}
	return res, nil
}

func (*server) Orders(ctx context.Context,
	empty *empty.Empty) (*mookiespb.OrdersResponse, error) {

	log.Println("Orders function was invoked")
	active := []*mookiespb.Order{}

	for _, o := range Orders {
		if o.GetStatus() == mookiespb.Order_ACTIVE {
			active = append(active, o)
		}
	}

	res := &mookiespb.OrdersResponse{
		Orders: active,
	}

	return res, nil
}

func NewServer(db *sqlx.DB) *server {
	return &server{
		db: db,
	}
}

func main() {
	seedData()
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
	server := NewServer(db)
	err = server.db.Ping()
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

const itemDbSchema = `
CREATE TABLE IF NOT EXISTS items (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	price DECIMAL NOT NULL
	category_id INTEGER,
	CONSTRAINT fk_categories
		FOREIGN KEY (category_id)
		REFERENCES categories(id)
);`

const categoryDbSchema = `
CREATE TABLE IF NOT EXISTS categories (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL
);`

func seedData() {
	data := []*mookiespb.Category{
		{
			Name: "Sandwich",
			Items: []*mookiespb.Item{
				{Name: "Large Smoked Pulled Pork", Id: 1, Price: 495, CategoryID: 1},
				{Name: "Regular Smoked Pulled Pork", Id: 2, Price: 395, CategoryID: 1},
				{Name: "Large Smoked Chicken Breast", Id: 3, Price: 495, CategoryID: 1},
				{Name: "Regular Smoked Chicken Breast", Id: 4, Price: 395, CategoryID: 1},
				{Name: "'The Molly'", Id: 5, Price: 395, CategoryID: 1},
				{Name: "Large Hamburger", Id: 6, Price: 495, CategoryID: 1},
				{Name: "Hamburger", Id: 7, Price: 395, CategoryID: 1},
				{Name: "Large Cheeseburger", Id: 8, Price: 550, CategoryID: 1},
				{Name: "Cheeseburger", Id: 9, Price: 425, CategoryID: 1},
				{Name: "Grilled Cheese", Id: 10, Price: 300, CategoryID: 1},
				{Name: "Pulled Pork Melt", Id: 11, Price: 395, CategoryID: 1},
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
