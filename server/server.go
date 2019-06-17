package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/cskr/pubsub"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/jbpratt78/mookies-tos/data"
	mookiespb "github.com/jbpratt78/mookies-tos/protofiles"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const topicOrder = "orders"
const topicComplete = "complete"

var (
	reg         = prometheus.NewRegistry()
	grpcMetrics = grpc_prometheus.NewServerMetrics()
	promAddr    = flag.String("prom", ":9001", "Port to run metrics HTTP server")
	listen      = flag.String("listen", ":50051", "listen address")
	dbp         = flag.String("database", "./mookies.db", "database to use")
	lpDev       = flag.String("p", "/dev/usb/lp0", "Printer dev file")
	crt         = flag.String("crt", "server.crt", "TLS cert to use")
	key         = flag.String("key", "server.key", "TLS key to use")
	kasp        = keepalive.ServerParameters{Time: 5 * time.Second}
)

type server struct {
	db     *sqlx.DB
	orders []*mookiespb.Order
	menu   *mookiespb.Menu
	ps     *pubsub.PubSub
}

func (s *server) GetMenu(ctx context.Context, empty *mookiespb.Empty) (*mookiespb.Menu, error) {
	log.Println("Client has requested the menu")
	res := s.menu
	return res, nil
}

func (s *server) SubmitOrder(ctx context.Context,
	req *mookiespb.SubmitOrderRequest) (*mookiespb.SubmitOrderResponse, error) {

	log.Println("An order was received")
	o := req.GetOrder()
	// expecting it to be right id
	o.Status = "active"

	err := submitOrder(s.db, o)
	if err != nil {
		return nil, err
	}

	res := &mookiespb.SubmitOrderResponse{
		Result: "Order has been placed..",
	}

	go publish(s.ps, o, topicOrder)

	return res, s.LoadData()
}

func (s *server) SubscribeToOrders(req *mookiespb.Empty,
	stream mookiespb.OrderService_SubscribeToOrdersServer) error {

	log.Printf("Client has subscribed to orders: %v\n", req)
	ch := s.ps.Sub(topicOrder)
	for {
		if o, ok := <-ch; ok {
			log.Printf("Sending order to client: %v\n", o)
			err := stream.Send(o.(*mookiespb.Order))
			if err != nil {
				return err
			}
		}
	}
}

func (s *server) SubscribeToCompleteOrders(req *mookiespb.Empty,
	stream mookiespb.OrderService_SubscribeToCompleteOrdersServer) error {

	log.Printf("Client has subscribed to orders: %v\n", req)
	ch := s.ps.Sub(topicComplete)
	for {
		if o, ok := <-ch; ok {
			log.Printf("Sending order to client: %v\n", o)
			err := stream.Send(o.(*mookiespb.Order))
			if err != nil {
				return err
			}
		}
	}
}

func publish(ps *pubsub.PubSub, order *mookiespb.Order, topic string) {
	ps.Pub(order, topic)
}

func (s *server) CompleteOrder(ctx context.Context,
	req *mookiespb.CompleteOrderRequest) (*mookiespb.CompleteOrderResponse, error) {

	log.Printf("Client is completing order: %v\n", req)
	// update order to be complete
	for _, o := range s.orders {
		if req.GetId() == o.GetId() {
			o.Status = "complete"
			go publish(s.ps, o, topicComplete)
		}
	}

	err := completeOrder(s.db, req.GetId())
	if err != nil {
		return nil, err
	}
	res := &mookiespb.CompleteOrderResponse{
		Result: "Order marked as complete",
	}

	return res, s.LoadData()
}

func (s *server) ActiveOrders(
	ctx context.Context, empty *mookiespb.Empty) (*mookiespb.OrdersResponse, error) {

	log.Println("Client has requested active orders")
	res := &mookiespb.OrdersResponse{
		Orders: s.orders,
	}
	return res, nil
}

func (s *server) LoadData() error {
	menu, err := getMenu(s.db)
	if err != nil {
		return err
	}
	s.menu = menu

	orders, err := getOrders(s.db)
	s.orders = orders

	log.Println("Menu and orders have been successfully queried")
	return nil
}

func NewServer(db *sqlx.DB) (*server, error) {
	server := &server{db: db}
	server.ps = pubsub.New(0)
	err := server.LoadData()
	//err = seedData(s.db)
	if err != nil {
		return nil, err
	}
	return server, nil
}

func init() {
	reg.MustRegister(grpcMetrics)
}

func main() {
	flag.Parse()

	/*creds, err := credentials.NewServerTLSFromFile(*crt, *key)
	if err != nil {
		log.Fatalf("Could not load server/key paid: %s", err)
	}*/

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

	httpServer := &http.Server{
		Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
		Addr:    *promAddr,
	}

	s := grpc.NewServer(
		grpc.KeepaliveParams(kasp),
		grpc.StreamInterceptor(grpcMetrics.StreamServerInterceptor()),
		grpc.UnaryInterceptor(grpcMetrics.UnaryServerInterceptor()),
		/*grpc.Creds(creds),*/
	)
	mookiespb.RegisterMenuServiceServer(s, server)
	mookiespb.RegisterOrderServiceServer(s, server)

	grpcMetrics.InitializeMetrics(s)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal("Unable to start a http server.")
		}
	}()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func seedData(db *sqlx.DB) error {
	tx, err := db.Begin()
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

func submitOrder(db *sqlx.DB, o *mookiespb.Order) error {
	tx, err := db.Begin()
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

func getMenu(db *sqlx.DB) (*mookiespb.Menu, error) {
	var categories []*mookiespb.Category
	menu := &mookiespb.Menu{
		Categories: categories,
	}
	err := db.Select(&menu.Categories, "SELECT * from categories")
	for _, category := range menu.GetCategories() {
		err = db.Select(&category.Items,
			fmt.Sprintf("SELECT * FROM items WHERE category_id = %v", category.GetId()))
		if err != nil {
			return nil, err
		}
		for _, item := range category.GetItems() {
			err = db.Select(&item.Options, fmt.Sprintf(
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

func getOrders(db *sqlx.DB) ([]*mookiespb.Order, error) {
	var orders []*mookiespb.Order
	err := db.Select(&orders,
		"SELECT * FROM orders WHERE status = 'active'")
	if err != nil {
		return nil, err
	}
	for _, order := range orders {
		err = db.Select(&order.Items, fmt.Sprintf(
			`
			SELECT name,price,items.id,order_items.id as order_item_id
			FROM items JOIN order_items ON items.id = order_items.item_id 
			WHERE order_id = %d`, order.GetId()))
		if err != nil {
			return nil, err
		}
		for _, item := range order.GetItems() {
			err = db.Select(&item.Options, fmt.Sprintf(
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

func completeOrder(db *sqlx.DB, id int32) error {
	tx, err := db.Begin()
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
