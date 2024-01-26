package main

import (
	"github.com/w3gop2p/go_design_patterns/BehavioralPatterns/Command/cmd"
)

func main() {
	tv := &cmd.Tv{}

	onCommand := &cmd.OnCommand{
		Device: tv,
	}

	offCommand := &cmd.OffCommand{
		Device: tv,
	}

	onButton := &cmd.Button{
		Command: onCommand,
	}
	onButton.Press()

	offButton := &cmd.Button{
		Command: offCommand,
	}
	offButton.Press()
}
