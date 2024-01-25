package cmd

type Computer interface {
	Print()
	SetPrinter(Printer)
}
