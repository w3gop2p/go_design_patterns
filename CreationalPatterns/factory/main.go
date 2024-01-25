package factory

import (
	"fmt"
	cmd "github.com/w3gop2p/go_design_patterns/CreationalPatterns/factory/cmd"
)

func main() {
	ak47, _ := cmd.GetGun("ak47")
	musket, _ := cmd.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g cmd.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
