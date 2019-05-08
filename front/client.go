package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/label"
	"github.com/aarzilli/nucular/rect"
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

	// debug
	DebugEnabled bool
	DebugStrings []string

	// Popup
	PSelect []bool

	// Layout
	groupSelectedItem *mookiespb.Item

	// current order
	NameEditor nucular.TextEditor
	order      *mookiespb.Order
	menu       *mookiespb.Menu
	Theme      nstyle.Theme
	client     client
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

	// TlO this need to change dynamically
	l.NameEditor.Flags = nucular.EditField

	l.order = &mookiespb.Order{}

	l.client = client{}
	return l
}

type client struct {
	MenuClient  mookiespb.MenuServiceClient
	OrderClient mookiespb.OrderServiceClient
}

func main() {
	cc, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer cc.Close()
	l := newLayout()
	l.client.MenuClient = mookiespb.NewMenuServiceClient(cc)
	l.client.OrderClient = mookiespb.NewOrderServiceClient(cc)

	l.menu, _ = l.doMenuRequest()
	groupFlags := nucular.WindowFlags(0)
	groupFlags |= nucular.WindowNoScrollbar
	wnd := nucular.NewMasterWindow(groupFlags, "Mookies", l.basicDemo)
	wnd.Main()
}

func (l *layout) doMenuRequest() (*mookiespb.Menu, error) {
	l.debug("Starting to request menu...")
	req := &empty.Empty{}

	res, err := l.client.MenuClient.GetMenu(context.Background(), req)
	if err != nil {
		return nil, err
	}
	l.debug("Response from GetMenu: %v\n", res.GetCategories())
	return res, nil
}

// pass in order as arg
func (l *layout) doSubmitOrderRequest(order *mookiespb.Order) {

	l.debug("Starting order request")
	req := &mookiespb.SubmitOrderRequest{
		Order: order,
	}

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)

	res, err := l.client.OrderClient.SubmitOrder(ctx, req)
	if err != nil {
		log.Fatalf("Error while submitting order RPC: %v\n", err)
	}
	l.debug("Response from SubmitOrder: %v\n", res.GetResult())
}

func (l *layout) basicDemo(w *nucular.Window) {
	l.overviewLayout(w)
}

// need to pass error into here
func (l *layout) errorPopup(w *nucular.Window) {
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

func (l *layout) overviewLayout(w *nucular.Window) {
	w.Row(30).Ratio(0.1, 0.8, 0.1)
	if w.Button(label.T("debug"), false) {
		l.DebugEnabled = !l.DebugEnabled
	}
	w.Label(time.Now().Format("3:04PM"), "CC")
	w.Spacing(1)
	// creates a row of height 20 with 1 column
	w.Row(20).Dynamic(1)
	// puts this text in the column with alignment x:left - y:center
	w.Label("Menu:", "LC")

	// creates a row of height 0 (Dynamic sizing) with 2 columns
	w.Row(w.Bounds.H-70).Ratio(0.7, 0.3)

	// create flags for the group we're about to create, turn off horizontal scrollbar and turn on borders
	groupFlags := nucular.WindowFlags(0)
	groupFlags |= nucular.WindowBorder
	groupFlags |= nucular.WindowNoHScrollbar

	// creates a group and puts it in the first column
	if sw := w.GroupBegin("Group", groupFlags); sw != nil {

		if l.DebugEnabled {
			sw.Row(int(sw.Bounds.H / 2)).Dynamic(1)
			if debugWindow := sw.GroupBegin("debug", groupFlags); debugWindow != nil {
				for _, line := range l.DebugStrings {
					wrapped := wrapText(line, 100)
					lines := strings.Split(wrapped, "\n")
					for _, subLine := range lines {
						debugWindow.Row(20).Dynamic(1)
						debugWindow.Label(subLine, "LC")
					}
				}
				debugWindow.GroupEnd()
			}
		}

		categories := l.menu.GetCategories()
		for _, category := range categories {
			if sw.TreePush(nucular.TreeTab, category.GetName(), false) {
				newRow := 4
				for _, item := range category.GetItems() {
					if newRow == 4 {
						newRow = 0
						sw.Row(100).Dynamic(4)
					}
					text := wrapText(item.GetName(), int(float64(sw.Bounds.W)/4.0/8.0))
					if sw.Button(label.T(text), false) {
						l.order.Items = append(l.order.Items, item)
					}
					newRow++
				}
				sw.TreePop()
			}
		}
		sw.GroupEnd()
	}

	// create new flags, turn off horizontal scrollbar
	groupFlags = nucular.WindowFlags(0)
	groupFlags |= nucular.WindowNoScrollbar

	// creates a second group and puts it in the second column
	if sw := w.GroupBegin("asdasd", groupFlags); sw != nil {

		var sum float32
		newHeight := sw.Bounds.H - 117
		sw.Row(newHeight).Dynamic(1)
		groupFlags = nucular.WindowFlags(0)
		groupFlags |= nucular.WindowBorder
		if orderWindow := sw.GroupBegin("asdasd", groupFlags); sw != nil {
			if len(l.order.Items) > 0 {
				for itemNumber, item := range l.order.Items {
					sum += item.GetPrice() / 100
					lines := strings.Split(wrapText(item.Name, int(float64(orderWindow.Bounds.W-95)/8.0)), "\n")
					for i, line := range lines {
						// more spacing between items
						if i == 0 {
							if len(lines) > 1 {
								orderWindow.Row(12).Static(orderWindow.Bounds.W-95, 55, 20)
							} else {
								orderWindow.Row(20).Static(orderWindow.Bounds.W-95, 55, 20)
							}
							orderWindow.Label("• "+line, "LT")
							orderWindow.Label(fmt.Sprintf("$ %5v", fmt.Sprintf("%.2f", item.GetPrice()/100)), "RT")
							// button height need to be fixed
							if orderWindow.Button(label.T("X"), false) {
								l.order.Items = append(l.order.Items[:itemNumber], l.order.Items[itemNumber+1:]...)
							}
						} else {
							if i == len(lines)-1 {
								orderWindow.Row(20).Dynamic(1)
							} else {
								orderWindow.Row(12).Dynamic(1)
							}
							orderWindow.Label("  "+line, "LT")
						}
					}

				}
			}
			orderWindow.GroupEnd()
		}

		sw.Row(20).Dynamic(2)
		sw.Label("Sum:", "LC")
		sw.Label(fmt.Sprintf("$ %.2f", sum), "RC")
		sw.Row(20).Dynamic(2)
		sw.Label("After Tax:", "LC")
		sw.Label(fmt.Sprintf("$ %.2f", sum*1.04), "RC")

		sw.Row(25).Dynamic(1)
		l.NameEditor.Edit(sw)

		sw.Row(40).Dynamic(1)
		if sw.Button(label.T("ORDER"), false) {
			if len(l.NameEditor.Buffer) > 0 && len(l.order.Items) > 0 {
				l.order.Name = string(l.NameEditor.Buffer)
				l.order.Total = float32(math.Round(float64(sum*100)) / 100)
				l.doSubmitOrderRequest(l.order)
				l.order.Reset()
				l.NameEditor.Buffer = nil
			} else {
				w.Master().PopupOpen("Please give the order a name :)", nucular.WindowMovable|nucular.WindowTitle|nucular.WindowDynamic|nucular.WindowNoScrollbar, rect.Rect{20, 100, 230, 150}, true, l.errorPopup)
			}
		}
		sw.GroupEnd()
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

func (l *layout) debug(message string, v ...interface{}) {
	msg := fmt.Sprintf(message, v...)
	log.Printf(msg)
	// append to debug slice
	l.DebugStrings = append(l.DebugStrings, msg)
}
