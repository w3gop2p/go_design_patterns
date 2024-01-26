package main

import (
	"github.com/w3gop2p/go_design_patterns/BehavioralPatterns/Mediator/cmd"
)

func main() {
	stationManager := cmd.NewStationManger()

	passengerTrain := &cmd.PassengerTrain{
		Mediator: stationManager,
	}
	freightTrain := &cmd.FreightTrain{
		Mediator: stationManager,
	}

	passengerTrain.Arrive()
	freightTrain.Arrive()
	passengerTrain.Depart()
}
