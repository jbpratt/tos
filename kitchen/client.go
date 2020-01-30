package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	tospb "github.com/jbpratt78/tos/protofiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/keepalive"

	"github.com/jroimartin/gocui"
)

type app struct {
	CompleteOrderIndex int

	// orders
	Orders             []*tospb.Order
	LastCompletedOrder *tospb.Order

	// debug
	DebugEnabled bool
	DebugStrings []string

	client tospb.OrderServiceClient
}

func newapp() *app {
	return &app{CompleteOrderIndex: 0}
}

var (
	kacp = keepalive.ClientParameters{
		Time:                60 * time.Second,
		Timeout:             time.Second,
		PermitWithoutStream: true,
	}

	tls  = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	cert = flag.String("cert", "cert/server.crt", "The file containing the CA root cert file")
	addr = flag.String("addr", "server:50051", "server address to dial")

	log grpclog.LoggerV2
	reg = prometheus.NewRegistry()

	orders []*tospb.Order
)

func init() {
	log = grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
	grpclog.SetLoggerV2(log)
}

func connecttospberver() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	if *tls {
		creds, err := credentials.NewClientTLSFromFile(*cert, "")
		if err != nil {
			log.Fatal(err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))

	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	grpcMetrics := grpc_prometheus.NewClientMetrics()
	reg.MustRegister(grpcMetrics)

	opts = append(opts,
		grpc.WithStreamInterceptor(grpcMetrics.StreamClientInterceptor()),
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor()),
		grpc.WithUnaryInterceptor(grpcMetrics.UnaryClientInterceptor()),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor()),
		grpc.WithKeepaliveParams(kacp),
	)

	return grpc.Dial(*addr, opts...)
}

func main() {
	flag.Parse()

	cc, err := connecttospberver()
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer cc.Close()

	httpServer := &http.Server{
		Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
		Addr:    fmt.Sprintf("0.0.0.0:%d", 9004),
	}

	// Start your http server for prometheus.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal("Unable to start a http server.")
		}
	}()

	l := newapp()
	l.client = tospb.NewOrderServiceClient(cc)

	if err = l.requestOrders(); err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			err := l.subscribeToOrders()
			log.Warningf("subscription failed: %v", err)
			time.Sleep(time.Second * 5)
		}
	}()

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()
	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Fatal(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatal(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		for _, order := range orders {
			fmt.Fprintln(v, order.Id)
		}
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (l *app) completeOrder(order *tospb.Order, i int) error {
	id := order.GetId()
	log.Infoln("Starting complete order request...")
	// take this order req in as param
	req := &tospb.CompleteOrderRequest{
		Id: id,
	}
	res, err := l.client.CompleteOrder(context.Background(), req)
	if err != nil {
		return err
	}
	log.Infof("Response from CompleteOrder: %v\n", res.GetResponse())

	l.Orders = append(l.Orders[:i], l.Orders[i+1:]...)
	// so we have the option for an undo button if the order was dismissed too early
	l.LastCompletedOrder = order
	return nil
}

func (l *app) requestOrders() error {
	req := &tospb.Empty{}

	res, err := l.client.ActiveOrders(context.Background(), req)
	if err != nil {
		return err
	}
	l.Orders = res.GetOrders()
	orders = l.Orders
	log.Infof("Response from Orders: %v\n", l.Orders)
	return nil
}

func (l *app) subscribeToOrders() error {
	log.Infoln("Subscribing to orders...")
	req := &tospb.Empty{}

	stream, err := l.client.SubscribeToOrders(context.Background(), req)
	if err != nil {
		return err
	}
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			// end of stream, never hope to hit?
			// or call subscribeToOrders often
			return err
		}
		if err != nil {
			fmt.Printf("got error trying to get order %v", err)
			return err
		}
		l.Orders = append(l.Orders, order)
		log.Infof("Received order: %v\n", order)
	}
}

func wrapText(text string, width int) string {
	if len(text) <= width {
		return text
	}
	var output string
	subStrings := strings.Split(text, " ")
	for i, subString := range subStrings {
		if len(output)+len(subString)+1 <= width || i == 0 {
			output += subString + " "
		} else {
			output += "\n" + wrapText(strings.Join(subStrings[i:], " "), width)
			break
		}
	}
	return output
}

func (l *app) debug(message string, v ...interface{}) {
	msg := fmt.Sprintf(message, v...)
	log.Infof(msg)
	// append to debug slice
	l.DebugStrings = append(l.DebugStrings, msg)
}

func (l *app) addToCompleteOrderIndex(i int) error {
	if i > 9 {
		return errors.New("number too large")
	}
	l.CompleteOrderIndex *= 10
	l.CompleteOrderIndex += i
	return nil
}

func (l *app) resetCompleteOrderIndex() {
	l.CompleteOrderIndex = 0
}
