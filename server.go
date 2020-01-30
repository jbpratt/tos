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
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/jbpratt78/tos/models"
	tospb "github.com/jbpratt78/tos/protofiles"
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

const (
	topicOrder = "orders"
	// topicComplete = "complete"
)

var (
	reg         = prometheus.NewRegistry()
	grpcMetrics = grpc_prometheus.NewServerMetrics()
	kasp        = keepalive.ServerParameters{
		Time: 60 * time.Second,
	}
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	prnt     = flag.Bool("print", true, "Use printer for complete orders")
	promAddr = flag.String("prom", ":9001", "Port to run metrics HTTP server")
	webAddr  = flag.String("web", ":9090", "Port to run grpc-web HTTP server")
	addr     = flag.String("addr", ":50051", "listen address")
	dbp      = flag.String("database", "/tmp/tos.db", "database to use")
	crt      = flag.String("crt", "cert/server.crt", "TLS cert to use")
	key      = flag.String("key", "cert/server.key", "TLS key to use")
	lpDev    = flag.String("p", "/dev/usb/lp0", "Printer dev file")
	logger   *logrus.Logger
)

type server struct {
	services *models.Services
	orders   []*tospb.Order
	menu     *tospb.Menu
	ps       *pubsub.PubSub
}

func (s *server) GetMenu(ctx context.Context,
	empty *tospb.Empty) (*tospb.Menu, error) {

	if len(s.menu.GetCategories()) == 0 {
		return nil, status.Error(codes.NotFound, "menu is empty")
	}
	return s.menu, nil
}

func (s *server) CreateMenuItem(ctx context.Context,
	req *tospb.Item) (*tospb.CreateMenuItemResponse, error) {

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

	id, err := s.services.Menu.CreateMenuItem(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"services..CreateMenuItem(%v) returned %v", req, err)
	}

	return &tospb.CreateMenuItemResponse{Id: id}, s.loadData()
}

func (s *server) UpdateMenuItem(ctx context.Context,
	req *tospb.Item) (*tospb.Response, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "item id not provided")
	}

	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument,
			"req must have itemid (non 0)")
	}

	if err := s.services.Menu.UpdateMenuItem(req); err != nil {
		return nil, status.Errorf(codes.Internal,
			"services..UpdateMenuItem(%v) returned %v", req, err)
	}

	return &tospb.Response{Response: "success"}, s.loadData()
}

func (s *server) DeleteMenuItem(ctx context.Context,
	req *tospb.DeleteMenuItemRequest) (*tospb.Response, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument,
			"req must not be nil")
	}

	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument,
			"req must have itemid (non 0)")
	}

	if err := s.services.Menu.DeleteMenuItem(req.GetId()); err != nil {
		return nil, status.Errorf(codes.Internal,
			"services..UpdateMenuItem(%v) returned %v", req, err)
	}

	return &tospb.Response{Response: "success"}, s.loadData()
}

func (s *server) CreateMenuItemOption(ctx context.Context,
	req *tospb.Option) (*tospb.Response, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument,
			"item option not provided")
	}

	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *server) SubmitOrder(ctx context.Context,
	req *tospb.Order) (*tospb.Response, error) {

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

	if req.GetName() == "" || req.GetName() == " " {
		return nil, status.Error(codes.InvalidArgument,
			"order name not provided")
	}

	o := req
	o.Status = "active"

	if err := s.services.Order.SubmitOrder(o); err != nil {
		return nil, status.Errorf(codes.Internal, "internal err: %v", err)
	}

	publish(s.ps, o, topicOrder)

	return &tospb.Response{Response: "Order has been placed.."}, s.loadData()
}

func (s *server) SubscribeToOrders(req *tospb.Empty,
	stream tospb.OrderService_SubscribeToOrdersServer) error {

	logger.Infoln("Client has subscribed to orders")

	ch := s.ps.Sub(topicOrder)
	for {
		if o, ok := <-ch; ok {
			logger.Printf("Sending order to client: %v\n", o)
			err := stream.Send(o.(*tospb.Order))
			if err != nil {
				s.ps.Unsub(ch, topicOrder)
				return status.Errorf(codes.Internal,
					"stream.Send(%v) failed with %v", o, err)
			}
		}
	}
}

func publish(ps *pubsub.PubSub, order *tospb.Order, topic string) {
	ps.Pub(order, topic)
}

func (s *server) CompleteOrder(ctx context.Context,
	req *tospb.CompleteOrderRequest) (*tospb.Response, error) {

	if req.GetId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument,
			"order id must be non zero")
	}

	for _, o := range s.orders {
		if req.GetId() == o.GetId() {
			o.Status = "complete"
			if *prnt {
				err := printOrder(o)
				if err != nil {
					return nil, status.Errorf(codes.Internal,
						"printer not established")
				}
			}
			break
		}
	}

	if err := s.services.Order.CompleteOrder(req.GetId()); err != nil {
		return nil, err
	}

	return &tospb.Response{Response: "Order marked as complete"}, s.loadData()
}

func printOrder(o *tospb.Order) error {
	f, err := os.OpenFile(*lpDev, os.O_RDWR, 0)
	if err != nil {
		return err
	}

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

	return f.Close()
}

func (s *server) ActiveOrders(
	ctx context.Context, empty *tospb.Empty) (*tospb.OrdersResponse, error) {

	if s.orders == nil {
		return nil, status.Errorf(codes.Internal,
			"ActiveOrders() failed: server.orders has not been initialized")
	}

	return &tospb.OrdersResponse{Orders: s.orders}, nil
}

func (s *server) loadData() error {
	menu, err := s.services.Menu.GetMenu()
	if err != nil {
		return err
	}

	if len(menu.GetCategories()) == 0 {
		if err = s.services.Menu.SeedMenu(); err != nil {
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
		s.orders = []*tospb.Order{}
	}

	return nil
}

func newServer() (*server, error) {
	services, err := newServices()
	if err != nil {
		return nil, err
	}
	server := &server{services: services, ps: pubsub.New(0)}
	if err = server.loadData(); err != nil {
		return nil, err
	}
	return server, nil
}

func newServices() (*models.Services, error) {
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

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		logger.Fatalf("Failed to listen: %v\n", err)
	}
	logger.Printf("Listening on %q...\n", *addr)

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

	if *dbp == "/tmp/tos.db" {

	}

	server, err := newServer()
	if err != nil {
		logger.Fatal(err)
	}
	defer server.services.Close()

	promHTTPServer := &http.Server{
		Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
		Addr:    *promAddr,
	}

	s := grpc.NewServer(opts...)

	wrappedServer := grpcweb.WrapServer(s)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		setupResponse(&resp, req)
		wrappedServer.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Handler: http.HandlerFunc(handler),
		Addr:    *webAddr,
	}
	logger.Println("gRPC web proxy listening at ", *webAddr)

	tospb.RegisterMenuServiceServer(s, server)
	tospb.RegisterOrderServiceServer(s, server)

	grpcMetrics.InitializeMetrics(s)

	go func() {
		if err := promHTTPServer.ListenAndServe(); err != nil {
			logger.Fatalf("Prom http server failed to start: %v", err)
		}
	}()

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			logger.Fatalf("grpc-web http server failed to start: %v", err)
		}
	}()

	if err := s.Serve(lis); err != nil {
		logger.Fatalf("gRPC server failed to serve: %v", err)
	}
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-grpc-web")
}
