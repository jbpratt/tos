package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/label"
	nstyle "github.com/aarzilli/nucular/style"

	"github.com/golang/protobuf/ptypes/empty"
	mookiespb "github.com/jbpratt78/mookies-tos/protofiles"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", ":50051", "address to dial")
)

type layout struct {
	ShowMenu    bool
	Titlebar    bool
	Border      bool
	Resize      bool
	Movable     bool
	NoScrollbar bool
	Minimizable bool
	Close       bool

	// orders
	Orders []*mookiespb.Order

	// debug
	DebugEnabled bool
	DebugStrings []string

	// Popup
	PSelect []bool

	// current order
	Theme  nstyle.Theme
	client mookiespb.OrderServiceClient
}

func newLayout() (l *layout) {
	l = &layout{}
	l.ShowMenu = true
	l.Titlebar = true
	l.Border = true
	l.Resize = true
	l.Movable = true
	l.NoScrollbar = false
	l.Close = true

	return l
}

func main() {
	cc, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer cc.Close()

	l := newLayout()
	l.client = mookiespb.NewOrderServiceClient(cc)

	err = l.requestOrders(l.client)
	if err != nil {
		log.Fatal(err)
	}
	go l.subscribeToOrders(l.client)

	wnd := nucular.NewMasterWindow(0, "Mookies", l.basicDemo)
	wnd.Main()
}

func completeOrder(c mookiespb.OrderServiceClient) error {
	fmt.Println("Starting complete order request...")
	// take this order req in as param
	req := &mookiespb.CompleteOrderRequest{
		Id: 1,
	}
	res, err := c.CompleteOrder(context.Background(), req)
	if err != nil {
		return err
	}
	log.Printf("Response from CompleteOrder: %v\n", res.GetResult())
	return nil
}

func (l *layout) requestOrders(c mookiespb.OrderServiceClient) error {
	req := &empty.Empty{}

	res, err := c.ActiveOrders(context.Background(), req)
	if err != nil {
		return err
	}
	l.Orders = res.GetOrders()
	//log.Printf("Response from Orders: %v\n", l.Orders)
	return nil
}

func (l *layout) subscribeToOrders(c mookiespb.OrderServiceClient) error {

	fmt.Println("Subscribing to orders...")
	req := &mookiespb.SubscribeToOrderRequest{
		Request: "send me orders",
	}

	stream, err := c.SubscribeToOrders(context.Background(), req)
	if err != nil {
		return err
	}
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			// end of stream, never hope to hit?
			// or call subscribeToOrders often
			break
		}
		if err != nil {
			return nil
		}
		l.Orders = append(l.Orders, order)
		log.Printf("Received order: %v\n", order)
		log.Printf("Order status: %v\n", order.GetStatus())
	}
	return nil
}

func (od *layout) overviewLayout(w *nucular.Window) {
	w.Row(30).Ratio(0.1, 0.8, 0.1)
	if w.Button(label.T("debug"), false) {
		od.DebugEnabled = !od.DebugEnabled
	}
	w.Label(time.Now().Format("3:04PM"), "CC")
	w.Spacing(1)
	w.Row(20).Dynamic(1)
	w.Label("Orders:", "LC")

	groupFlags := nucular.WindowFlags(0)
	groupFlags |= nucular.WindowBorder

	// create group that contains all orders
	w.Row(0).Dynamic(1)
	if ordersWindow := w.GroupBegin("orders", groupFlags); ordersWindow != nil {
		// create a row with a column for every order
		widths := make([]int, len(od.Orders))
		for index := 0; index < len(widths); index++ {
			widths[index] = 150
		}
		ordersWindow.Row(0).Static(widths...)
		for _, order := range od.Orders {
			// create group for each order
			if singleOrderWindow := ordersWindow.GroupBegin("fuck you", groupFlags); singleOrderWindow != nil {
				// create a row for text
				singleOrderWindow.Row(20).Dynamic(1)
				singleOrderWindow.Label(order.GetName(), "LC")
				singleOrderWindow.GroupEnd()
			}
		}
		ordersWindow.GroupEnd()
	}
}

func (l *layout) basicDemo(w *nucular.Window) {
	l.overviewLayout(w)
}
