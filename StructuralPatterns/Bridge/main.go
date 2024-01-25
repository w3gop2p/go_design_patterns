package main

import (
	"fmt"
	cmd "github.com/w3gop2p/go_design_patterns/StructuralPatterns/Bridge/cmd"
)

func main() {

	hpPrinter := &cmd.Hp{}
	epsonPrinter := &cmd.Epson{}

	macComputer := &cmd.Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &cmd.WindowsOs{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
