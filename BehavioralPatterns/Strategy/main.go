package main

import (
	"github.com/w3gop2p/go_design_patterns/BehavioralPatterns/Strategy/cmd"
)

func main() {
	lfu := &cmd.Lfu{}
	cache := cmd.InitCache(lfu)

	cache.Add("a", "1")
	cache.Add("b", "2")

	cache.Add("c", "3")

	lru := &cmd.Lru{}
	cache.SetEvictionAlgo(lru)

	cache.Add("d", "4")

	fifo := &cmd.Fifo{}
	cache.SetEvictionAlgo(fifo)

	cache.Add("e", "5")

}
