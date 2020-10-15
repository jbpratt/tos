package main

import (
	"os"
	"strings"

	"github.com/jbpratt/tos/pkg/printer"
)

func main() {
	printer, err := printer.NewFromPath("/dev/usb/lp0")
	if err != nil {
		panic(err)
	}
	printer.Print(strings.Join(os.Args[1:], " "))
}
