package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/label"
	"github.com/aarzilli/nucular/rect"
	"github.com/aarzilli/nucular/style"
	nstyle "github.com/aarzilli/nucular/style"
	mookiespb "github.com/jbpratt78/tos/protofiles"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"golang.org/x/mobile/event/key"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/keepalive"
)

const taxRate = 1.04

type layout struct {
	ShowMenu    bool
	Titlebar    bool
	Border      bool
	Resize      bool
	Movable     bool
	NoScrollbar bool
	Minimizable bool
	Close       bool

	TextVisible bool

	// debug
	DebugEnabled bool
	DebugStrings []string

	//item creation
	catID int64

	Error error

	// Layout
	groupSelectedItem *mookiespb.Item

	// current order
	CurrentItem             *mookiespb.Item
	NameEditor              nucular.TextEditor
	CustomOptionNameEditor  nucular.TextEditor
	CustomOptionPriceEditor nucular.TextEditor
	order                   *mookiespb.Order
	menu                    *mookiespb.Menu
	Theme                   nstyle.Theme
	client                  client
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
	l.Theme = nstyle.DarkTheme

	// TODO: this need to change dynamically
	l.NameEditor.Flags = nucular.EditField
	l.CustomOptionNameEditor.Flags = nucular.EditField
	l.CustomOptionPriceEditor.Flags = nucular.EditField

	l.order = &mookiespb.Order{}

	l.client = client{}
	return l
}

type client struct {
	MenuClient  mookiespb.MenuServiceClient
	OrderClient mookiespb.OrderServiceClient
}

var (
	kacp = keepalive.ClientParameters{
		Time:                60 * time.Second,
		Timeout:             time.Second,
		PermitWithoutStream: true,
	}

	tls  = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	cert = flag.String("cert", "cert/server.crt", "The file containing the CA root cert file")
	addr = flag.String("addr", "server:50051", "server to dial")

	log grpclog.LoggerV2
)

func init() {
	log = grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)
}

func main() {
	flag.Parse()
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

	reg := prometheus.NewRegistry()
	grpcMetrics := grpc_prometheus.NewClientMetrics()
	reg.MustRegister(grpcMetrics)

	opts = append(opts,
		grpc.WithUnaryInterceptor(grpcMetrics.UnaryClientInterceptor()),
		grpc.WithKeepaliveParams(kacp),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor()),
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cc, err := grpc.DialContext(ctx, *addr, opts...)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer cc.Close()

	httpServer := &http.Server{
		Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
		Addr:    fmt.Sprintf("0.0.0.0:%d", 9003)}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal("Unable to start a http server.")
		}
	}()

	l := newLayout()
	l.client.MenuClient = mookiespb.NewMenuServiceClient(cc)
	l.client.OrderClient = mookiespb.NewOrderServiceClient(cc)

	l.menu, _ = l.GetMenu()
	groupFlags := nucular.WindowFlags(0)
	groupFlags |= nucular.WindowNoScrollbar
	wnd := nucular.NewMasterWindow(groupFlags, "Mookies", l.basicDemo)
	wnd.SetStyle(style.FromTheme(l.Theme, 1.0))
	wnd.Main()
}

func (l *layout) GetMenu() (*mookiespb.Menu, error) {
	l.debug("Starting to request menu...")
	req := &mookiespb.Empty{}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	res, err := l.client.MenuClient.GetMenu(ctx, req)
	if err != nil {
		return nil, err
	}
	l.debug("Response from GetMenu: %v\n", res.GetCategories())
	return res, nil
}

func (l *layout) SubmitOrder(order *mookiespb.Order) error {
	l.debug("Starting order request")
	req := order

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	res, err := l.client.OrderClient.SubmitOrder(ctx, req)
	if err != nil {
		return err
	}
	l.debug("Response from SubmitOrder: %v\n", res.GetResponse())
	return nil
}

func (l *layout) CreateMenuItem(item *mookiespb.Item) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	req := item

	// check res here
	_, err := l.client.MenuClient.CreateMenuItem(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (l *layout) DeleteMenuItem(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	req := &mookiespb.DeleteMenuItemRequest{Id: id}

	_, err := l.client.MenuClient.DeleteMenuItem(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (l *layout) basicDemo(w *nucular.Window) {
	l.overviewLayout(w)
}

// need to pass error into here
func (l *layout) errorPopup(w *nucular.Window) {
	w.Row(25).Dynamic(1)
	w.Label(l.Error.Error(), "CC")
	w.Row(25).Dynamic(1)
	if w.Button(label.T("OK"), false) {
		w.Close()
	}
}

func (l *layout) itemOptionPopup(w *nucular.Window) {
	options := l.CurrentItem.GetOptions()

	sort.Slice(options, func(i, j int) bool {
		return options[i].Selected && !options[j].Selected
	})

	for _, option := range options {
		w.Row(30).Static(w.Bounds.W-100, 83)
		w.CheckboxText(option.GetName(), &option.Selected)
		w.Label(fmt.Sprintf("$ %.2f", option.GetPrice()/100.0), "RC")
	}

	w.Row(30).Dynamic(1)
	if w.Button(label.T("Custom Option"), false) {
		l.TextVisible = !l.TextVisible
	}
	if l.TextVisible {
		w.Row(30).Static(w.Bounds.W-90, 10, 83)
		l.CustomOptionNameEditor.Edit(w)
		w.Label("$", "CC")
		l.CustomOptionPriceEditor.Edit(w)
		w.Row(30).Dynamic(1)
		if w.Button(label.T("Add"), false) {
			price := string(l.CustomOptionPriceEditor.Buffer)
			if s, err := strconv.ParseFloat(price, 32); err == nil {
				newOption := new(mookiespb.Option)
				newOption.Price = float32(s * 100)
				newOption.Name = string(l.CustomOptionNameEditor.Buffer)
				newOption.Selected = true
				l.CustomOptionPriceEditor.Buffer = nil
				l.CustomOptionNameEditor.Buffer = nil
				l.CurrentItem.Options = append(options, newOption)
				l.TextVisible = false
			}
		}
	}

	w.Row(30).Dynamic(2)
	if w.Button(label.T("OK"), false) {
		l.order.Items = append(l.order.Items, l.CurrentItem)
		if l.CurrentItem.GetCategoryID() == 2 || l.CurrentItem.GetCategoryID() == 3 {
			count := 0
			for _, option := range l.CurrentItem.GetOptions() {
				if option.GetSelected() {
					count++
				}
			}
			if count > 2 || count < 2 {
				l.Error = errors.New("Item has too many or too few options")
				w.Master().PopupOpen("Error", nucular.WindowMovable|nucular.WindowTitle|nucular.WindowDynamic|nucular.WindowNoScrollbar, rect.Rect{400, 100, 300, 150}, true, l.errorPopup)
			}
		}
		w.Close()
	}
	if w.Button(label.T("Cancel"), false) {
		l.CurrentItem = nil
		w.Close()
	}
}

func (l *layout) newMenuItemPopup(w *nucular.Window) {
	item := &mookiespb.Item{}

	var catList []string

	for _, cat := range l.menu.GetCategories() {
		catList = append(catList, cat.GetName())
	}

	w.Row(30).Dynamic(1)
	l.catID = int64(w.ComboSimple(catList, int(l.catID), 30))

	w.Row(30).Static(w.Bounds.W-90, 10, 83)
	l.CustomOptionNameEditor.Edit(w)
	w.Label("$", "CC")
	l.CustomOptionPriceEditor.Edit(w)
	w.Row(30).Dynamic(1)
	if w.Button(label.T("ADD"), false) {
		price := string(l.CustomOptionPriceEditor.Buffer)
		if s, err := strconv.ParseFloat(price, 32); err == nil && len(l.CustomOptionNameEditor.Buffer) >= 1 {
			item.Name = string(l.CustomOptionNameEditor.Buffer)
			l.CustomOptionNameEditor.Buffer = nil
			item.Price = float32(s * 100)
			item.CategoryID = l.catID + 1
			l.CreateMenuItem(item)
			log.Infof("added item %v with price $ %v in category %v.", item.Name, item.Price/100, item.CategoryID)
			l.menu, _ = l.GetMenu()
			// TODO: does not work atm
			w.Close()
		}
	}
	if w.Button(label.T("Cancel"), false) {
		l.CurrentItem = nil
		w.Close()
	}
}

func (l *layout) settings(w *nucular.Window) {
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

func (l *layout) overviewLayout(w *nucular.Window) {
	l.keybindings(w)
	w.Row(30).Ratio(0.1, 0.8, 0.1)
	w.Spacing(1)
	w.Label(time.Now().Format("3:04PM"), "CC")
	if w.Button(label.T("settings"), false) {
		w.Master().PopupOpen("Select options:",
			nucular.WindowMovable, rect.Rect{200, 100, 230, 200}, true, l.settings)
	}
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
						l.CurrentItem = cloneItem(item)
						if len(item.GetOptions()) < 1 {
							l.order.Items = append(l.order.Items, l.CurrentItem)
						} else {
							w.Master().PopupOpen("Select options:",
								nucular.WindowMovable|nucular.WindowTitle|nucular.WindowDynamic|nucular.WindowNoScrollbar, rect.Rect{(w.Bounds.W / 2) - 115, 100, 230, (len(item.GetOptions()) * 40) + 140}, true, l.itemOptionPopup)
						}
					}
					newRow++
				}
				if newRow == 4 {
					newRow = 0
					sw.Row(100).Dynamic(4)
				}
				if sw.Button(label.T(wrapText("Create new item", int(float64(sw.Bounds.W)/4.0/8.0))), false) {
					l.catID = category.Id - 1
					w.Master().PopupOpen("Create new Item:", nucular.WindowMovable|nucular.WindowTitle|nucular.WindowDynamic|nucular.WindowNoScrollbar, rect.Rect{(w.Bounds.W / 2) - 115, 100, 230, (5 * 40) + 140}, true, l.newMenuItemPopup)
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
	if sw := w.GroupBegin("orders", groupFlags); sw != nil {
		newHeight := sw.Bounds.H - 117
		sw.Row(newHeight).Dynamic(1)
		groupFlags = nucular.WindowFlags(0)
		groupFlags |= nucular.WindowNoHScrollbar
		groupFlags |= nucular.WindowBorder
		if orderWindow := sw.GroupBegin("orders", groupFlags); orderWindow != nil {
			if len(l.order.Items) > 0 {
				for itemNumber, item := range l.order.Items {

					lines := strings.Split(wrapText(item.Name, int(float64(orderWindow.Bounds.W-115)/8.0)), "\n")
					numberOfOptionsActive := 0
					for _, option := range item.GetOptions() {
						if option.GetSelected() {
							numberOfOptionsActive++
						}
					}
					orderWindow.Row((len(lines)*12)+15+numberOfOptionsActive*15).Static(orderWindow.Bounds.W-50, 30)
					groupFlags = nucular.WindowFlags(0)
					groupFlags |= nucular.WindowNoScrollbar
					if itemWindow := orderWindow.GroupBegin(item.GetName(), groupFlags); itemWindow != nil {
						for i, line := range lines {
							// more spacing between items
							if i == 0 {
								itemWindow.Row(12).Static(itemWindow.Bounds.W-65, 10)
								itemWindow.Label("• "+line, "LT")
								itemWindow.Label(
									fmt.Sprintf("$ %5v", fmt.Sprintf("%.2f", item.GetPrice()/100)), "RT")
							} else {
								itemWindow.Row(12).Dynamic(1)
								itemWindow.Label("  "+line, "LT")
							}
						}
						for _, option := range item.GetOptions() {
							if option.GetSelected() {
								itemWindow.Row(12).Static(itemWindow.Bounds.W-65, 10)
								itemWindow.Label("    • "+option.GetName(), "LT")
								itemWindow.Label(fmt.Sprintf("$ %5v", fmt.Sprintf("%.2f", option.GetPrice()/100)), "RT")
							}
						}

						itemWindow.GroupEnd()
					}
					if orderWindow.Button(label.T("X"), false) {
						l.order.Items = append(l.order.Items[:itemNumber], l.order.Items[itemNumber+1:]...)
					}
				}
			}
			orderWindow.GroupEnd()
		}

		sum := calculateSum(l.order)
		price := calculatePrice(l.order)
		sw.Row(20).Dynamic(2)
		sw.Label("Sum:", "LC")
		sw.Label(fmt.Sprintf("$ %.2f", sum), "RC")
		sw.Row(20).Dynamic(2)
		sw.Label("After Tax:", "LC")
		sw.Label(fmt.Sprintf("$ %.2f", price), "RC")

		sw.Row(25).Dynamic(1)
		l.NameEditor.Edit(sw)

		sw.Row(40).Dynamic(1)
		if sw.Button(label.T("ORDER"), false) {
			l.sendOrder(w)
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
	log.Info(msg)
	// append to debug slice
	l.DebugStrings = append(l.DebugStrings, msg)
}

func cloneItem(originalItem *mookiespb.Item) *mookiespb.Item {
	newItem := new(mookiespb.Item)
	newItem.Id = originalItem.GetId()
	newItem.Name = originalItem.GetName()
	newItem.Price = originalItem.GetPrice()
	newItem.Options = make([]*mookiespb.Option, len(originalItem.GetOptions()))
	newItem.CategoryID = originalItem.GetCategoryID()
	for i, option := range originalItem.GetOptions() {
		newItem.Options[i] = new(mookiespb.Option)
		newItem.Options[i].Id = option.GetId()
		newItem.Options[i].Name = option.GetName()
		newItem.Options[i].Price = option.GetPrice()
		newItem.Options[i].Selected = option.GetSelected()
	}
	return newItem
}

func (l *layout) keybindings(w *nucular.Window) {
	mw := w.Master()
	if in := w.Input(); in != nil {
		k := in.Keyboard
		for _, e := range k.Keys {
			scaling := mw.Style().Scaling
			switch {
			case (e.Code == key.CodeReturnEnter):
				l.sendOrder(w)
			case (e.Code == key.CodeF12):
				l.DebugEnabled = !l.DebugEnabled
			case (e.Modifiers == key.ModControl || e.Modifiers == key.ModControl|key.ModShift) && (e.Code == key.CodeZ):
				// TODO: theme pop up to pick from theme list
				log.Infoln("pop up theme list")
			case (e.Modifiers == key.ModControl || e.Modifiers == key.ModControl|key.ModShift) && (e.Code == key.CodeEqualSign):
				mw.Style().Scale(scaling + 0.1)
			case (e.Modifiers == key.ModControl || e.Modifiers == key.ModControl|key.ModShift) && (e.Code == key.CodeHyphenMinus):
				mw.Style().Scale(scaling - 0.1)
			}
		}
	}
}

func (l *layout) sendOrder(w *nucular.Window) {
	sum := calculatePrice(l.order)
	if len(l.NameEditor.Buffer) > 0 && len(l.order.Items) > 0 {
		l.order.Name = string(l.NameEditor.Buffer)
		l.order.Total = float32(math.Round(float64(sum*100)) / 100)
		l.SubmitOrder(l.order)
		l.order.Reset()
		l.NameEditor.Buffer = nil
	} else {
		l.Error = errors.New("Please give the order a name")
		w.Master().PopupOpen("Error",
			nucular.WindowMovable|nucular.WindowTitle|nucular.WindowDynamic|nucular.WindowNoScrollbar,
			rect.Rect{20, 100, 230, 150}, true, l.errorPopup)
	}
}

func calculateSum(order *mookiespb.Order) float32 {
	var sum float32
	for _, item := range order.GetItems() {
		sum += item.GetPrice() / 100
		for _, option := range item.GetOptions() {
			if option.GetSelected() {
				sum += option.GetPrice() / 100
			}
		}
	}
	sum *= 1.04
	return sum
}

func calculatePrice(order *mookiespb.Order) float32 {
	return calculateSum(order) * taxRate
}
