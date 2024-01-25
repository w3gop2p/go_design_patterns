package main

import (
	"fmt"
	"github.com/w3gop2p/go_design_patterns/StructuralPatterns/Decorator/cmd"
)

func main() {

	pizza := &cmd.VeggieMania{}

	//Add cheese topping
	pizzaWithCheese := &cmd.CheeseTopping{
		Pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &cmd.TomatoTopping{
		Pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrice())
}
