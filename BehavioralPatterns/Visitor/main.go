package main

import (
	"fmt"
	"github.com/w3gop2p/go_design_patterns/BehavioralPatterns/Visitor/cmd"
)

func main() {
	square := &cmd.Square{Side: 2}
	circle := &cmd.Circle{Radius: 3}
	rectangle := &cmd.Rectangle{L: 2, B: 3}

	areaCalculator := &cmd.AreaCalculator{}

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &cmd.MiddleCoordinates{}
	square.Accept(middleCoordinates)
	circle.Accept(middleCoordinates)
	rectangle.Accept(middleCoordinates)
}
