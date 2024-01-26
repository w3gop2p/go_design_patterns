package main

import (
	"github.com/w3gop2p/go_design_patterns/BehavioralPatterns/ChainOfResponsability/cmd"
)

func main() {

	cashier := &cmd.Cashier{}

	//Set next for medical department
	medical := &cmd.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &cmd.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &cmd.Reception{}
	reception.SetNext(doctor)

	patient := &cmd.Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)
}
