package main

import (
	"context"
	"flag"
	"fmt"
	"image"
	"log"

	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/label"
	nstyle "github.com/aarzilli/nucular/style"
	"github.com/golang/protobuf/ptypes"
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
	Mprog   int
	Mslider int
	Mcheck  bool
	Prog    int
	Slider  int
	Check   bool

	// Basic widgets status
	IntSlider                          int
	FloatSlider                        float64
	ProgValue                          int
	PropertyFloat                      float64
	PropertyInt                        int
	PropertyNeg                        int
	RangeMin, RangeValue, RangeMax     float64
	RangeIntMin, RangeInt, RangeIntMax int
	Checkbox                           bool

	// Selectable
	Selected  []bool
	Selected2 []bool

	// Popup
	PSelect []bool
	PProg   int
	PSlider int

	// Layout
	GroupBorder             bool
	GroupNoScrollbar        bool
	GroupWidth, GroupHeight int
	groupCurrent int
	groupSelectedItem *mookiespb.Item
	GroupSelected []bool

	// order list
	orderContents []*mookiespb.Item

	// Vertical Split
	A, B, C    int
	HA, HB, HC int

	Img *image.RGBA

	Resizing1, Resizing2, Resizing3, Resizing4 bool

	edEntry1, edEntry2, edEntry3 nucular.TextEditor

	Theme nstyle.Theme
}

func newLayout() (od *layout) {
	od = &layout{}
	od.ShowMenu = true
	od.Titlebar = true
	od.Border = true
	od.Resize = true
	od.Movable = true
	od.NoScrollbar = false
	od.Close = true
	od.HeaderAlign = nstyle.HeaderRight
	od.Mprog = 60
	od.Mslider = 8
	od.Mcheck = true
	od.Prog = 40
	od.Slider = 10
	od.Check = true
	od.IntSlider = 5
	od.FloatSlider = 2.5
	od.ProgValue = 40
	od.Selected = []bool{false, false, true, false}
	od.Selected2 = []bool{true, false, false, false, false, true, false, false, false, false, true, false, false, false, false, true}
	od.PSelect = make([]bool, 4)
	od.PProg = 0
	od.PSlider = 10

	//Layout
	od.GroupBorder = true
	od.GroupNoScrollbar = false

	// TODO this need to change dynamically
	od.GroupSelected = make([]bool, 10000)
	od.groupCurrent = -1

	od.A = 100
	od.B = 100
	od.C = 100

	od.HA = 100
	od.HB = 100
	od.HC = 100

	od.PropertyFloat = 2
	od.PropertyInt = 10
	od.PropertyNeg = 10

	od.RangeMin = 0
	od.RangeValue = 50
	od.RangeMax = 100

	od.RangeIntMin = 0
	od.RangeInt = 2048
	od.RangeIntMax = 4096

	od.GroupWidth = 0
	od.GroupHeight = 0

	return od
}

var menu *mookiespb.Menu
var client mookiespb.MenuServiceClient

func main() {
	cc, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer cc.Close()
	//orderClient := mookiespb.NewOrderServiceClient(cc)
	client = mookiespb.NewMenuServiceClient(cc)
	updateMenu()

	//items := menu.GetItems()
	//doSubmitOrderRequest(orderClient)

	l := newLayout()
	wnd := nucular.NewMasterWindow(0, "Mookies", l.basicDemo)
	wnd.Main()
}

func updateMenu() {
	menu, _ = doMenuRequest(client)
}

func doMenuRequest(c mookiespb.MenuServiceClient) (*mookiespb.Menu, error) {
	fmt.Println("Starting to request menu...")
	req := &empty.Empty{}

	res, err := c.GetMenu(context.Background(), req)
	if err != nil {
		return nil, err
	}
	log.Printf("Response from GetMenu: %v\n", res.GetItems())
	return res, nil
}

func doSubmitOrderRequest(c mookiespb.OrderServiceClient) {
	fmt.Println("Starting order request")
	req := &mookiespb.SubmitOrderRequest{
		Order: &mookiespb.Order{
			Id:   1,
			Name: "Majora",
			Items: []*mookiespb.Item{
				{Name: "Large Smoked Pulled Pork", Id: 1, Price: 495, Category: "Sandwich"},
			},
			Total:       495,
			TimeOrdered: ptypes.TimestampNow(),
		},
	}

	res, err := c.SubmitOrder(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while submitting order RPC: %v\n", err)
	}
	log.Printf("Response from SubmitOrder: %v\n", res.GetResult())
}

func (l *layout) basicDemo(w *nucular.Window) {
	l.overviewLayout(w)

	//if l.Debug {
	//	l.showDebug(w)
	//}
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
	w.Row(30).Ratio(0.4,0.2,0.4)
	w.Spacing(1)
	if w.Button(label.T("update menu"), false) {
		updateMenu()
	}
	w.Spacing(1)
	
	// creates a row of height 20 with 1 column
	w.Row(20).Dynamic(1)
	// puts this text in the column with alignment x:left - y:center
	w.Label("Menu:", "LC")
	
	// creates a row of height 0 (Dynamic sizing) with 2 columns
	w.Row(0).Dynamic(2)
	
	// create flags for the group we're about to create, turn off horizontal scrollbar and turn on borders
	groupFlags := nucular.WindowFlags(0)
	groupFlags |= nucular.WindowBorder
	groupFlags |= nucular.WindowNoHScrollbar

	// creates a group and puts it in the first column
	if sw := w.GroupBegin("Group", groupFlags); sw != nil {
		sw.Row(18).Static(od.GroupWidth)
		items := menu.GetItems()
		for _, item := range items {
			if sw.Button(label.T(item.GetName()), false) {
				od.orderContents = append(od.orderContents, item)
				od.groupSelectedItem = item
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
			sw.Label("category: ", "RC")
			sw.Label(od.groupSelectedItem.GetCategory(), "LC")
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

		if len(od.orderContents) > 0 {
			var sum float32
			for i, item := range od.orderContents {
				sum += item.GetPrice()/100
				sw.Row(20).Ratio(0.7,0.2,0.1)
				sw.Label(fmt.Sprintf("%v", item.GetName()), "LC")
				sw.Label(fmt.Sprintf("$ %.2f", item.GetPrice()/100), "RC")
				if sw.Button(label.T("X"), false) {
					od.orderContents = append(od.orderContents[:i], od.orderContents[i+1:]...)
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
