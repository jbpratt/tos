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
	"github.com/jbpratt78/tos/models"
	mookiespb "github.com/jbpratt78/tos/protofiles"
	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
)

const topicOrder = "orders"
const topicComplete = "complete"

var (
	reg         = prometheus.NewRegistry()
	grpcMetrics = grpc_prometheus.NewServerMetrics()
	kasp        = keepalive.ServerParameters{
		Time: 60 * time.Second,
	}
	promAddr = flag.String("prom", ":9001", "Port to run metrics HTTP server")
	listen   = flag.String("listen", ":50051", "listen address")
	dbp      = flag.String("database", "./mookies.db", "database to use")
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	crt      = flag.String("crt", "server.crt", "TLS cert to use")
	key      = flag.String("key", "server.key", "TLS key to use")
)

type server struct {
	services *models.Services
	orders   []*mookiespb.Order
	menu     *mookiespb.Menu
	ps       *pubsub.PubSub
}

func (s *server) GetMenu(ctx context.Context, empty *mookiespb.Empty) (*mookiespb.Menu, error) {
	log.Println("Client has requested the menu")
	res := s.menu
	return res, nil
}

func (s *server) CreateMenuItem(ctx context.Context,
	req *mookiespb.CreateMenuItemRequest) (*mookiespb.CreateMenuItemResponse, error) {

	i := req.GetItem()
	err := s.services.Menu.CreateMenuItem(i)
	if err != nil {
		return nil, err
	}

	res := &mookiespb.CreateMenuItemResponse{
		Result: "success",
	}

	return res, s.LoadData()
}

func (s *server) UpdateMenuItem(ctx context.Context,
	req *mookiespb.UpdateMenuItemRequest) (*mookiespb.UpdateMenuItemResponse, error) {

	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *server) DeleteMenuItem(ctx context.Context,
	req *mookiespb.DeleteMenuItemRequest) (*mookiespb.DeleteMenuItemResponse, error) {

	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *server) CreateMenuItemOption(ctx context.Context,
	req *mookiespb.CreateMenuItemOptionRequest) (*mookiespb.CreateMenuItemOptionResponse, error) {

	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *server) SubmitOrder(ctx context.Context,
	req *mookiespb.SubmitOrderRequest) (*mookiespb.SubmitOrderResponse, error) {

	log.Println("An order was received")
	o := req.GetOrder()
	o.Status = "active"

	err := s.services.Order.SubmitOrder(o)
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

	err := s.services.Order.CompleteOrder(req.GetId())
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
	menu, err := s.services.Menu.GetMenu()
	if err != nil {
		return err
	}
	s.menu = menu

	orders, err := s.services.Order.GetOrders()
	if err != nil {
		return err
	}
	s.orders = orders

	return nil
}

func NewServer() (*server, error) {
	services, err := NewServices()
	if err != nil {
		return nil, err
	}
	server := &server{services: services, ps: pubsub.New(0)}
	//err = server.services.Menu.SeedMenu()
	err = server.LoadData()
	if err != nil {
		return nil, err
	}
	return server, nil
}

func NewServices() (*models.Services, error) {
	services, err := models.NewServices(
		models.WithSqlite(*dbp),
		models.WithMenu(),
		models.WithOrder(),
	)
	if err != nil {
		return nil, err
	}
	return services, nil
}

func init() {
	reg.MustRegister(grpcMetrics)
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	log.Printf("Listening on %q...\n", *listen)

	var opts []grpc.ServerOption
	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*crt, *key)
		if err != nil {
			log.Fatalf("Could not load server/key pair: %s", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	opts = append(opts,
		grpc.KeepaliveParams(kasp),
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             10 * time.Second,
				PermitWithoutStream: true,
			},
		),
		grpc.StreamInterceptor(grpcMetrics.StreamServerInterceptor()),
		grpc.UnaryInterceptor(grpcMetrics.UnaryServerInterceptor()))

	server, err := NewServer()
	if err != nil {
		log.Fatal(err)
	}
	defer server.services.Close()

	httpServer := &http.Server{
		Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
		Addr:    *promAddr,
	}

	s := grpc.NewServer(opts...)

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
