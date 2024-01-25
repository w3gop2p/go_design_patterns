package main

import (
	"fmt"
	"github.com/w3gop2p/go_design_patterns/StructuralPatterns/Facade/cmd"
	"log"
)

func main() {
	fmt.Println()
	walletFacade := cmd.NewWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
