package main

import (
	"context"
	"flag"
	"fmt"
	"image"
	"log"
	"time"

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

	GroupSelected []bool

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

	od.GroupSelected = make([]bool, 16)

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

	od.GroupWidth = 320
	od.GroupHeight = 200

	return od
}

func main() {
	cc, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	menuClient := mookiespb.NewMenuServiceClient(cc)
	//orderClient := mookiespb.NewOrderServiceClient(cc)

	defer cc.Close()
	doMenuRequest(menuClient)
	//doSubmitOrderRequest(orderClient)

	l := newLayout()
	wnd := nucular.NewMasterWindow(0, "Mookies", l.basicDemo)
	wnd.Main()
}

func doMenuRequest(c mookiespb.MenuServiceClient) {
	fmt.Println("Starting to request menu...")
	req := &empty.Empty{}

	res, err := c.GetMenu(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GetMenu RPC: %v\n", err)
	}
	log.Printf("Response from GetMenu: %v\n", res.GetItems())
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
	if l.ShowMenu {
		l.overviewMenubar(w)
	}
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

func (l *layout) overviewMenubar(w *nucular.Window) {
	w.Row(30).Dynamic(1)
	w.Label(time.Now().Format("15:04:05"), "RT")

	w.MenubarBegin()
	w.Row(25).Dynamic(1)
	if w := w.Menu(label.TA("MENU", "LC"), 120, nil); w != nil {
		w.Row(25).Dynamic(2)
		if w.MenuItem(label.TA("Hide", "LC")) {
			l.ShowMenu = false
		}

		perf := w.Master().GetPerf()
		w.CheckboxText("Show perf", &perf)
		w.Master().SetPerf(perf)
	}
	if w := w.Menu(label.TA("THEME", "CC"), 180, nil); w != nil {
		w.Row(25).Dynamic(1)
		newtheme := l.Theme
		if w.OptionText("Default Theme", newtheme == nstyle.DefaultTheme) {
			newtheme = nstyle.DefaultTheme
		}
		if w.OptionText("White Theme", newtheme == nstyle.WhiteTheme) {
			newtheme = nstyle.WhiteTheme
		}
		if w.OptionText("Red Theme", newtheme == nstyle.RedTheme) {
			newtheme = nstyle.RedTheme
		}
		if w.OptionText("Dark Theme", newtheme == nstyle.DarkTheme) {
			newtheme = nstyle.DarkTheme
		}
		if newtheme != l.Theme {
			l.Theme = newtheme
			w.Master().SetStyle(nstyle.FromTheme(l.Theme, w.Master().Style().Scaling))
			w.Close()
		}
	}
	w.MenubarEnd()
}

func (od *layout) overviewLayout(w *nucular.Window) {
	w.Row(20).Dynamic(1)
	if w.TreePush(nucular.TreeNode, "Widget", false) {
		btn := func() {
			w.Button(label.T("button"), false)
		}

		w.Row(30).Dynamic(1)
		w.Label("Dynamic fixed column layout with generated position and size (LayoutRowDynamic):", "LC")
		w.Row(30).Dynamic(3)
		btn()
		btn()
		btn()

		w.Row(30).Dynamic(1)
		w.Label("Dynamic array-based custom column layout with generated position and custom size (LayoutRowRatio):", "LC")
		w.Row(30).Ratio(0.2, 0.6, 0.2)
		btn()
		btn()
		btn()

		w.Row(30).Dynamic(1)
		w.Label("Static array-based custom column layout with dynamic space in the middle (LayoutRowStatic + LayoutResetStatic):", "LC")
		w.Row(30).Static(100, 100)
		btn()
		btn()
		w.LayoutResetStatic(0, 100, 100)
		w.Spacing(1)
		btn()
		btn()

		w.Row(30).Dynamic(1)
		w.Label("Dynamic immediate mode custom column layout with generated position and custom size (LayoutRowRatio):", "LC")
		w.Row(30).Ratio(0.2, 0.6, 0.2)
		btn()
		btn()
		btn()

		w.TreePop()
	}

	w.Row(20).Dynamic(1)
	if w.TreePush(nucular.TreeNode, "Group", false) {
		groupFlags := nucular.WindowFlags(0)
		if od.GroupBorder {
			groupFlags |= nucular.WindowBorder
		}
		if od.GroupNoScrollbar {
			groupFlags |= nucular.WindowNoScrollbar
		}

		groupFlags |= nucular.WindowNoHScrollbar

		w.Row(30).Dynamic(3)
		w.CheckboxText("Border", &od.GroupBorder)
		w.CheckboxText("No Scrollbar", &od.GroupNoScrollbar)

		w.Row(22).Static(50, 130, 130)
		w.Label("size:", "LC")
		w.PropertyInt("#Width:", 100, &od.GroupWidth, 500, 10, 10)
		w.PropertyInt("#Height:", 100, &od.GroupHeight, 500, 10, 10)

		w.Row(od.GroupHeight).Static(od.GroupWidth, od.GroupWidth)
		if sw := w.GroupBegin("Group", groupFlags); sw != nil {
			sw.Row(18).Static(100)
			for i := range od.GroupSelected {
				sel := "Unselected"
				if od.GroupSelected[i] {
					sel = "Selected"
				}
				sw.SelectableLabel(sel, "CC", &od.GroupSelected[i])
			}
			sw.GroupEnd()
		}
		w.TreePop()
	}
}
