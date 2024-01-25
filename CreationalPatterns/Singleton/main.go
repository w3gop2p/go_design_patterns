package Singleton

import (
	"fmt"
	"sync"
)

var (
	singleInstance *single
	once           sync.Once
)

type single struct {
}

func getInstance() *single {
	once.Do(func() {
		fmt.Println("Creating single instance now.")
		singleInstance = &single{}
	})

	fmt.Println("Single instance already created.")
	return singleInstance
}

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
