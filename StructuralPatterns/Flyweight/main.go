package main

import (
	"fmt"
	"github.com/w3gop2p/go_design_patterns/StructuralPatterns/Flyweight/cmd"
)

func main() {
	game := cmd.NewGame()

	//Add Terrorist
	game.AddTerrorist(cmd.TerroristDressType)
	game.AddTerrorist(cmd.TerroristDressType)
	game.AddTerrorist(cmd.TerroristDressType)
	game.AddTerrorist(cmd.TerroristDressType)

	//Add CounterTerrorist
	game.AddCounterTerrorist(cmd.CounterTerrroristDressType)
	game.AddCounterTerrorist(cmd.CounterTerrroristDressType)
	game.AddCounterTerrorist(cmd.CounterTerrroristDressType)

	dressFactoryInstance := cmd.GetDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.DressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
	}
}
