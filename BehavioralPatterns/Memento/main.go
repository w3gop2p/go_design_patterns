package main

import (
	"fmt"
	"github.com/w3gop2p/go_design_patterns/BehavioralPatterns/Memento/cmd"
)

func main() {

	caretaker := &cmd.Caretaker{
		MementoArray: make([]*cmd.Memento, 0),
	}

	originator := &cmd.Originator{
		State: "A",
	}

	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState("B")
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState("C")
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.RestoreMemento(caretaker.GetMemento(1))
	fmt.Printf("Restored to State: %s\n", originator.GetState())

	originator.RestoreMemento(caretaker.GetMemento(0))
	fmt.Printf("Restored to State: %s\n", originator.GetState())

}
