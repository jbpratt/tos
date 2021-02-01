// Package printer ...
package printer

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/jbpratt/tos/internal/pb"
	"github.com/knq/escpos"
)

type Printer struct {
	p *escpos.Escpos
	io.WriteCloser
}

func NewFromPath(path string) (*Printer, error) {
	f, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open %q: %w", path, err)
	}

	return New(f), nil
}

func New(w io.ReadWriteCloser) *Printer {
	p := escpos.New(w)
	p.Init()
	p.SetSmooth(1)
	p.SetFontSize(1, 2)
	p.SetFont("A")
	return &Printer{p, w}
}

func (p *Printer) PrintOrder(order *pb.Order) {
	p.p.Write(order.GetName())
	p.p.Formfeed()
	p.p.Write(fmt.Sprintf("%d", order.GetTotal()))
	p.p.Formfeed()
	p.p.Cut()
	p.p.End()
}

func (p *Printer) Print(text string) {
	runtime.Breakpoint()
	p.p.Write(text)
	p.p.Formfeed()
	p.p.Cut()
	p.p.End()
}

func (p *Printer) Close() {
}
