package main

import (
	"github.com/w3gop2p/go_design_patterns/BehavioralPatterns/Observer/cmd"
)

func main() {

	shirtItem := cmd.NewItem("Nike Shirt")

	observerFirst := &cmd.Customer{Id: "abc@gmail.com"}
	observerSecond := &cmd.Customer{Id: "xyz@gmail.com"}

	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)

	shirtItem.UpdateAvailability()
}
