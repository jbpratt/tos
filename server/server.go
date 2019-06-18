package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/cskr/pubsub"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/jbpratt78/tos/database"
	mookiespb "github.com/jbpratt78/tos/protofiles"
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
	service *database.Service
	orders  []*mookiespb.Order
	menu    *mookiespb.Menu
	ps      *pubsub.PubSub
}

type DBService interface {
	SeedData() error
	SubmitOrder(o *mookiespb.Order) error
	GetMenu() (*mookiespb.Menu, error)
	GetOrders() ([]*mookiespb.Order, error)
	CompleteOrder(id int32) error
	// CreateItem() error
	// DeleteItem() error
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

	err := s.service.SubmitOrder(o)
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

	err := s.service.CompleteOrder(req.GetId())
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
	menu, err := s.service.GetMenu()
	if err != nil {
		return err
	}
	s.menu = menu

	orders, err := s.service.GetOrders()
	s.orders = orders

	log.Println("Menu and orders have been successfully queried")
	return nil
}

func NewServer(db *sqlx.DB) (*server, error) {
	service := NewService(db)
	server := &server{service: service}
	server.ps = pubsub.New(0)
	err := server.LoadData()
	//err = seedData(s.db)
	if err != nil {
		return nil, err
	}
	return server, nil
}

func NewService(db *sqlx.DB) *database.Service {
	service := &database.Service{db}
	return service
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
