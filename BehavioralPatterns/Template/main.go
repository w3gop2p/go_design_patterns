package main

import (
	"fmt"
	"github.com/w3gop2p/go_design_patterns/BehavioralPatterns/Template/cmd"
)

func main() {
	smsOTP := &cmd.Sms{}
	o := cmd.Otp{
		IOtp: smsOTP,
	}
	o.GenAndSendOTP(4)

	fmt.Println("")
	emailOTP := &cmd.Email{}
	o = cmd.Otp{
		IOtp: emailOTP,
	}
	o.GenAndSendOTP(4)

}
