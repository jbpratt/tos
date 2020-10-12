package printer

import (
	"fmt"
	"io"

	"github.com/jbpratt/tos/pkg/pb"
)

type Printer struct {
	io.Writer
}

func (p *Printer) PrintOrder(order *pb.Order) error {
	return fmt.Errorf("unimplemented")
}
