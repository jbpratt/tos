package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/cskr/pubsub"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/jbpratt78/tos/models"
	mookiespb "github.com/jbpratt78/tos/protofiles"
	"github.com/knq/escpos"
	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
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
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	promAddr = flag.String("prom", ":9001", "Port to run metrics HTTP server")
	listen   = flag.String("listen", ":50051", "listen address")
	dbp      = flag.String("database", "./mookies.db", "database to use")
	crt      = flag.String("crt", "cert/server.crt", "TLS cert to use")
	key      = flag.String("key", "cert/server.key", "TLS key to use")
	lpDev    = flag.String("p", "/dev/usb/lp0", "Printer dev file")
	logger   *logrus.Logger
)

type server struct {
	services *models.Services
	orders   []*mookiespb.Order
	menu     *mookiespb.Menu
	ps       *pubsub.PubSub
}

func (s *server) GetMenu(ctx context.Context,
	empty *mookiespb.Empty) (*mookiespb.Menu, error) {

	if len(s.menu.GetCategories()) == 0 {
		return nil, status.Error(codes.NotFound, "menu is empty")
	}
	return s.menu, nil
}

func (s *server) CreateMenuItem(ctx context.Context,
	req *mookiespb.Item) (*mookiespb.Response, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "no item provided")
	}

	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "item must have a name")
	}

	if req.GetCategoryID() == 0 {
		return nil, status.Error(codes.InvalidArgument,
			"item must have a categoryID (non 0)")
	}

	// ignoring price in case of $0 item?

	// check if item aleady exists
	for _, c := range s.menu.GetCategories() {
		if c.GetId() == req.GetCategoryID() {
			for _, i := range c.GetItems() {
				if i.GetName() == req.GetName() {
					return nil, status.Error(codes.FailedPrecondition,
						"item already exists")
				}
			}
		}
	}

	err := s.services.Menu.CreateMenuItem(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"services..CreateMenuItem(%v) returned %v", req, err)
	}

	res := &mookiespb.Response{Response: "success"}
	return res, s.loadData()
}

func (s *server) UpdateMenuItem(ctx context.Context,
	req *mookiespb.Item) (*mookiespb.Response, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "item id not provided")
	}

	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument,
			"req must have itemid (non 0)")
	}

	//err := s.services.Menu.UpdateMenuItem(req)
	//if err != nil {
	//	return nil, status.Errorf(codes.Internal,
	//		"services..UpdateMenuItem(%v) returned %v", req, err)
	//}

	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *server) DeleteMenuItem(ctx context.Context,
	req *mookiespb.DeleteMenuItemRequest) (*mookiespb.Response, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument,
			"req must not be nil")
	}

	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument,
			"req must have itemid (non 0)")
	}

	err := s.services.Menu.DeleteMenuItem(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"services..UpdateMenuItem(%v) returned %v", req, err)
	}

	res := &mookiespb.Response{Response: "success"}
	return res, s.loadData()
}

func (s *server) CreateMenuItemOption(ctx context.Context,
	req *mookiespb.Option) (*mookiespb.Response, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument,
			"item option not provided")
	}

	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *server) SubmitOrder(ctx context.Context,
	req *mookiespb.Order) (*mookiespb.Response, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument,
			"order not provided")
	}

	if req.GetItems() == nil {
		return nil, status.Error(codes.InvalidArgument,
			"order items not provided")
	}

	if req.GetTotal() == 0 {
		return nil, status.Error(codes.InvalidArgument,
			"order price not provided")
	}

	o := req
	o.Status = "active"

	err := s.services.Order.SubmitOrder(o)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal err: %v", err)
	}

	res := &mookiespb.Response{Response: "Order has been placed.."}

	go publish(s.ps, o, topicOrder)

	return res, s.loadData()
}

func (s *server) SubscribeToOrders(req *mookiespb.Empty,
	stream mookiespb.OrderService_SubscribeToOrdersServer) error {

	logger.Infoln("Client has subscribed to orders")

	ch := s.ps.Sub(topicOrder)
	for {
		if o, ok := <-ch; ok {
			logger.Printf("Sending order to client: %v\n", o)
			err := stream.Send(o.(*mookiespb.Order))
			if err != nil {
				return status.Errorf(codes.Internal,
					"stream.Send(%v) failed with %v", o, err)
			}
		}
	}
}

func publish(ps *pubsub.PubSub, order *mookiespb.Order, topic string) {
	ps.Pub(order, topic)
}

func (s *server) CompleteOrder(ctx context.Context,
	req *mookiespb.CompleteOrderRequest) (*mookiespb.Response, error) {

	if req.GetId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument,
			"order id must be non zero")
	}

	// update order to be complete
	// TODO: handle if not found
	for _, o := range s.orders {
		if req.GetId() == o.GetId() {
			o.Status = "complete"
			// err := printOrder(o)
			// if err != nil {
			// 	return nil, status.Errorf(codes.Internal,
			//	"printer not established")
			// }
		}
	}

	err := s.services.Order.CompleteOrder(req.GetId())
	if err != nil {
		return nil, err
	}

	res := &mookiespb.Response{Response: "Order marked as complete"}

	return res, s.loadData()
}

func printOrder(o *mookiespb.Order) error {
	f, err := os.OpenFile(*lpDev, os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer f.Close()

	bw := bufio.NewWriter(f)

	w := bufio.NewReadWriter(nil, bw)
	p := escpos.New(w)
	p.Init()
	p.SetSmooth(1)
	p.SetFontSize(1, 2)
	p.SetFont("A")
	p.Write("Mookies Smokehouse")
	p.Formfeed()

	p.Write(o.GetName())
	p.Formfeed()
	p.Write(fmt.Sprintf("%f", o.GetTotal()))
	p.Formfeed()

	p.Cut()
	p.End()

	w.Flush()
	bw.Flush()
	return nil
}

func (s *server) ActiveOrders(
	ctx context.Context, empty *mookiespb.Empty) (*mookiespb.OrdersResponse, error) {

	if s.orders == nil {
		return nil, status.Errorf(codes.Internal,
			"ActiveOrders() failed: server.orders has not been initialized")
	}

	res := &mookiespb.OrdersResponse{Orders: s.orders}

	return res, nil
}

func (s *server) loadData() error {
	menu, err := s.services.Menu.GetMenu()
	if err != nil {
		return err
	}

	if len(menu.GetCategories()) == 0 {
		err = s.services.Menu.SeedMenu()
		if err != nil {
			return err
		}

		menu, err = s.services.Menu.GetMenu()
		if err != nil {
			return err
		}
	}
	s.menu = menu

	orders, err := s.services.Order.GetOrders()
	if err != nil {
		return err
	}
	s.orders = orders
	if s.orders == nil {
		s.orders = []*mookiespb.Order{}
	}

	return nil
}

func NewServer() (*server, error) {
	services, err := NewServices()
	if err != nil {
		return nil, err
	}
	server := &server{services: services, ps: pubsub.New(0)}
	err = server.loadData()
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
	logger = logrus.StandardLogger()
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
		DisableSorting:  true,
	})
	// Should only be done from init functions
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(logger.Out, logger.Out, logger.Out))
	reg.MustRegister(grpcMetrics)
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *listen)
	if err != nil {
		logger.Fatalf("Failed to listen: %v\n", err)
	}
	logger.Printf("Listening on %q...\n", *listen)

	var opts []grpc.ServerOption
	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*crt, *key)
		if err != nil {
			logger.Fatalf("Could not load server/key pair: %s", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	logrusEntry := logrus.NewEntry(logger)

	opts = append(opts,
		grpc.KeepaliveParams(kasp),
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             10 * time.Second,
				PermitWithoutStream: true,
			},
		),
		grpc_middleware.WithStreamServerChain(
			grpcMetrics.StreamServerInterceptor(),
			grpc_logrus.StreamServerInterceptor(logrusEntry),
		),
		grpc_middleware.WithUnaryServerChain(
			grpcMetrics.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
		),
	)

	server, err := NewServer()
	if err != nil {
		logger.Fatal(err)
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
			logger.Fatal("Unable to start a http server.")
		}
	}()

	if err := s.Serve(lis); err != nil {
		logger.Fatalf("Failed to serve: %v", err)
	}
}
