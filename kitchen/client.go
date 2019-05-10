package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/label"
	nstyle "github.com/aarzilli/nucular/style"

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
	Orders             []*mookiespb.Order
	LastCompletedOrder *mookiespb.Order

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

var wnd nucular.MasterWindow

func main() {
	cc, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer cc.Close()

	l := newLayout()
	l.client = mookiespb.NewOrderServiceClient(cc)

	err = l.requestOrders()
	if err != nil {
		log.Fatal(err)
	}
	go l.subscribeToOrders()

	groupFlags := nucular.WindowFlags(0)
	groupFlags |= nucular.WindowNoScrollbar
	wnd = nucular.NewMasterWindow(groupFlags, "Mookies", l.basicDemo)
	wnd.Main()
}

func (l *layout) completeOrder(id int32) error {
	fmt.Println("Starting complete order request...")
	// take this order req in as param
	req := &mookiespb.CompleteOrderRequest{
		Id: id,
	}
	res, err := l.client.CompleteOrder(context.Background(), req)
	if err != nil {
		return err
	}
	log.Printf("Response from CompleteOrder: %v\n", res.GetResult())
	return nil
}

func (l *layout) requestOrders() error {
	req := &mookiespb.Empty{}

	res, err := l.client.ActiveOrders(context.Background(), req)
	if err != nil {
		return err
	}
	l.Orders = res.GetOrders()
	log.Printf("Response from Orders: %v\n", l.Orders)
	return nil
}

func (l *layout) subscribeToOrders() error {

	fmt.Println("Subscribing to orders...")
	req := &mookiespb.SubscribeToOrderRequest{
		Request: "send me orders",
	}

	stream, err := l.client.SubscribeToOrders(context.Background(), req)
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
		wnd.Changed()
	}
	return nil
}

func (l *layout) overviewLayout(w *nucular.Window) {
	w.Row(30).Ratio(0.1, 0.8, 0.1)
	if w.Button(label.T("debug"), false) {
		l.DebugEnabled = !l.DebugEnabled
	}
	w.Label(time.Now().Format("3:04PM"), "CC")
	w.Spacing(1)
	w.Row(20).Dynamic(1)
	w.Label("Orders:", "LC")

	groupFlags := nucular.WindowFlags(0)
	// looks like there is no option to only have a horizontal scrollbar
	//groupFlags |= nucular.WindowNoScrollbar

	// create group that contains all orders
	w.Row(0).Dynamic(1)
	if ordersWindow := w.GroupBegin("orders", groupFlags); ordersWindow != nil {
		// create a row with a column for every order
		widths := make([]int, len(l.Orders))
		for index := 0; index < len(widths); index++ {
			widths[index] = 200
		}
		ordersWindow.Row(ordersWindow.Bounds.H - 15).Static(widths...)
		for i, order := range l.Orders {
			groupFlags = nucular.WindowFlags(0)
			groupFlags |= nucular.WindowNoHScrollbar
			groupFlags |= nucular.WindowBorder
			// create group for each order
			if singleOrderWindow := ordersWindow.GroupBegin(fmt.Sprint(order.Id), groupFlags); singleOrderWindow != nil {
				singleOrderWindow.Row(20).Dynamic(1)
				singleOrderWindow.Label(order.GetName(), "CC")
				singleOrderWindow.Row(20).Dynamic(1)
				if singleOrderWindow.Button(label.T("DONE"), false) {
					l.completeOrder(order.GetId())
					l.Orders = append(l.Orders[:i], l.Orders[i+1:]...)
					// so we have the option for an undo button if the order was dismissed too early
					l.LastCompletedOrder = order
				}

				for _, item := range order.Items {
					lines := strings.Split(wrapText(item.Name, 24), "\n")
					for i, line := range lines {
						// more spacing between items
						if i == len(lines)-1 {
							singleOrderWindow.Row(20).Dynamic(1)
						} else {
							singleOrderWindow.Row(12).Dynamic(1)
						}
						// bullet point for new orders
						if i == 0 {
							singleOrderWindow.Label("• "+line, "LT")
						} else {
							singleOrderWindow.Label("  "+line, "LT")
						}
					}
				}
				singleOrderWindow.GroupEnd()
			}
		}
		ordersWindow.GroupEnd()
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

func (l *layout) basicDemo(w *nucular.Window) {
	l.overviewLayout(w)
}

func (l *layout) debug(message string, v ...interface{}) {
	msg := fmt.Sprintf(message, v...)
	log.Printf(msg)
	// append to debug slice
	l.DebugStrings = append(l.DebugStrings, msg)
}