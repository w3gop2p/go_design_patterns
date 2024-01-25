package cmd

import "fmt"

type WindowsOs struct {
	printer Printer
}

func (w *WindowsOs) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *WindowsOs) SetPrinter(p Printer) {
	w.printer = p
}
