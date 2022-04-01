package main

import (
	"github.com/suborbital/reactr/api/tinygo/runnable" 
)

type Yo struct{}

func (h Yo) Run(input []byte) ([]byte, error) {
	return []byte("😀 yo (from TinyGo), " + string(input)), nil
}

// initialize runnable, do not edit //
func main() {
	runnable.Use(Yo{})
}
