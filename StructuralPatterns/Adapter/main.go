package main

import (
	cmd "github.com/w3gop2p/go_design_patterns/StructuralPatterns/Adapter/cmd"
)

func main() {

	client := &cmd.Client{}
	mac := &cmd.Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &cmd.Windows{}
	windowsMachineAdapter := &cmd.WindowsAdapter{
		WindowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
