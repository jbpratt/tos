package main

import (
	"context"
	"flag"
	"fmt"
	"log"

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

	HeaderAlign nstyle.HeaderAlign

	// Menu status
	// Selectable
	Selected  []bool
	Selected2 []bool

	// Popup
	PSelect []bool

	// Layout
	GroupBorder             bool
	GroupNoScrollbar        bool
	GroupWidth, GroupHeight int
	groupCurrent            int
	groupSelectedItem       *mookiespb.Item
	GroupSelected           []bool

	// current order
	order *mookiespb.Order
	menu  *mookiespb.Menu
	Theme nstyle.Theme
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
	l.HeaderAlign = nstyle.HeaderRight
	//Layout
	l.GroupBorder = true
	l.GroupNoScrollbar = false

	// TlO this need to change dynamically
	l.GroupSelected = make([]bool, 10000)
	l.groupCurrent = -1

	l.GroupWidth = 0
	l.GroupHeight = 0

	l.order = &mookiespb.Order{}

	return l
}

type Client struct {
	MenuClient  mookiespb.MenuServiceClient
	OrderClient mookiespb.OrderServiceClient
}

var client Client

func main() {
	cc, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer cc.Close()
	client.MenuClient = mookiespb.NewMenuServiceClient(cc)
	client.OrderClient = mookiespb.NewOrderServiceClient(cc)

	l := newLayout()
	l.menu, _ = doMenuRequest(client.MenuClient)
	wnd := nucular.NewMasterWindow(0, "Mookies", l.basicDemo)
	wnd.Main()
}

func doMenuRequest(c mookiespb.MenuServiceClient) (*mookiespb.Menu, error) {
	fmt.Println("Starting to request menu...")
	req := &empty.Empty{}

	res, err := c.GetMenu(context.Background(), req)
	if err != nil {
		return nil, err
	}
	log.Printf("Response from GetMenu: %v\n", res.GetCategories())
	return res, nil
}

// pass in order as arg
func doSubmitOrderRequest(
	c mookiespb.OrderServiceClient, order *mookiespb.Order) {

	fmt.Println("Starting order request")
	req := &mookiespb.SubmitOrderRequest{
		Order: order,
	}

	res, err := c.SubmitOrder(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while submitting order RPC: %v\n", err)
	}
	log.Printf("Response from SubmitOrder: %v\n", res.GetResult())
}

func (l *layout) basicDemo(w *nucular.Window) {
	l.overviewLayout(w)
}

// need to pass error into here
func errorPopup(w *nucular.Window) {
	w.Row(25).Dynamic(1)
	w.Label("Error", "CC")
	w.Row(25).Dynamic(2)
	if w.Button(label.T("OK"), false) {
		w.Close()
	}
	if w.Button(label.T("Cancel"), false) {
		w.Close()
	}
}

func (od *layout) overviewLayout(w *nucular.Window) {
	// creates a row of height 20 with 1 column
	w.Row(30).Dynamic(1)
	// puts this text in the column with alignment x:center - y:center
	w.Label("AYAYA", "CC")
	// creates a row of height 20 with 2 columns with dybnamic width
	w.Row(30).Ratio(0.4, 0.2, 0.4)
	w.Spacing(1)
	if w.Button(label.T("Send Order"), false) {
		doSubmitOrderRequest(client.OrderClient, od.order)
		od.order.Reset()
	}
	w.Spacing(1)

	// creates a row of height 20 with 1 column
	w.Row(20).Dynamic(1)
	// puts this text in the column with alignment x:left - y:center
	w.Label("Menu:", "LC")

	// creates a row of height 0 (Dynamic sizing) with 2 columns
	w.Row(0).Ratio(0.7, 0.3)

	// create flags for the group we're about to create, turn off horizontal scrollbar and turn on borders
	groupFlags := nucular.WindowFlags(0)
	groupFlags |= nucular.WindowBorder
	groupFlags |= nucular.WindowNoHScrollbar

	// creates a group and puts it in the first column
	if sw := w.GroupBegin("Group", groupFlags); sw != nil {
		sw.Row(18).Static(od.GroupWidth)
		categories := menu.GetCategories()
		for _, category := range categories {
			if sw.Button(label.T(category.GetName()), false) {
				fmt.Println(category)
				//od.order.Items = append(od.order.Items, item)
				//od.groupSelectedItem = item
			}
		}
		sw.GroupEnd()
	}

	// create new flags, turn off horizontal scrollbar
	groupFlags = nucular.WindowFlags(0)
	groupFlags |= nucular.WindowNoHScrollbar

	// creates a second group and puts it in the second column
	if sw := w.GroupBegin("asdasd", groupFlags); sw != nil {
		if od.groupSelectedItem != nil {
			sw.Row(20).Dynamic(2)
			sw.Label("name: ", "RC")
			sw.Label(od.groupSelectedItem.GetName(), "LC")
			sw.Row(20).Dynamic(2)
			sw.Label("price: ", "RC")
			sw.Label(fmt.Sprintf("$ %.2f", od.groupSelectedItem.GetPrice()/100), "LC")
			sw.Row(20).Dynamic(2)
			sw.Label("ID: ", "RC")
			sw.Label(fmt.Sprintf("%d", od.groupSelectedItem.GetId()), "LC")
		}

		sw.Row(10).Dynamic(1)
		sw.Spacing(1)

		if len(od.order.Items) > 0 {
			var sum float32
			for i, item := range od.order.Items {
				sum += item.GetPrice() / 100
				sw.Row(20).Ratio(0.7, 0.2, 0.1)
				sw.Label(fmt.Sprintf("%v", item.GetName()), "LC")
				sw.Label(fmt.Sprintf("$ %.2f", item.GetPrice()/100), "RC")
				if sw.Button(label.T("X"), false) {
					od.order.Items = append(od.order.Items[:i], od.order.Items[i+1:]...)
				}
			}

			sw.Row(10).Dynamic(1)
			sw.Spacing(1)

			sw.Row(20).Dynamic(2)
			sw.Label("Sum:", "LC")
			sw.Label(fmt.Sprintf("$ %.2f", sum), "RC")
		}
		sw.GroupEnd()
	}
}
