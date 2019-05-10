package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/jbpratt78/mookies-tos/data"
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
	for _, category := range menu.GetCategories() {
		err = s.db.Select(&category.Items,
			fmt.Sprintf("SELECT * FROM items WHERE category_id = %v", category.GetId()))
		if err != nil {
			return err
		}
		fmt.Println(category.GetItems())
		for _, item := range category.GetItems() {
			err = s.db.Select(&item.Options,
				fmt.Sprintf("SELECT name,price,by_default FROM options JOIN item_options as io ON options.id = io.option_id WHERE item_id = ?", item.GetId()))
			if err != nil {
				return errors.New(fmt.Sprintf("error on %v", item))
			}
		}

	}
	if err != nil {
		return err
	}
	s.menu = menu
	fmt.Println(s.menu)

	// TODO: query items along with orders
	//s.orders = orders
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
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	for i, category := range data.Menu {
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
					"INSERT INTO options (name, price, by_default) VALUES (?,?,?)",
					option.GetName(), option.GetPrice(), option.GetByDefault())
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
