package main

import (
	"fmt"
	"github.com/w3gop2p/go_design_patterns/BehavioralPatterns/Iterator/cmd"
)

func main() {

	user1 := &cmd.User{
		Name: "a",
		Age:  30,
	}
	user2 := &cmd.User{
		Name: "b",
		Age:  20,
	}

	userCollection := &cmd.UserCollection{
		Users: []*cmd.User{user1, user2},
	}

	iterator := userCollection.CreateIterator()

	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}
