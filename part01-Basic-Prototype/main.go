package main

import (
	"fmt"
	"publicChain/part01-Basic-Prototype/BLC"
)

func main() {
	block := BLC.NewBlock("TiN Block", 1, make([]byte, 32))
	fmt.Printf("block: %#v\n", block)
}
