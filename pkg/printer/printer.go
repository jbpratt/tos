package printer

import (
	"fmt"
	"io"
	"os"

	"github.com/jbpratt/tos/pkg/pb"
	"github.com/knq/escpos"
)

type Printer struct {
	*escpos.Escpos
}

func NewFromPath(path string) (*Printer, error) {
	f, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open %q: %w", path, err)
	}
	return New(f), nil
}

func New(w io.ReadWriter) *Printer {
	p := escpos.New(w)
	p.Init()
	p.SetSmooth(1)
	p.SetFontSize(1, 2)
	p.SetFont("A")
	return &Printer{p}
}

func (p *Printer) PrintOrder(order *pb.Order) {
	p.Write(order.GetName())
	p.Formfeed()
	p.Write(fmt.Sprintf("%f", order.GetTotal()))
	p.Formfeed()
	p.Cut()
	p.End()
}
