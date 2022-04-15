package main

import (
	"github.com/suborbital/reactr/api/tinygo/runnable" 
)

type HelloWorld struct{}

func (h HelloWorld) Run(input []byte) ([]byte, error) {
	return []byte("👋 Hello World 🌍"), nil
}

// initialize runnable, do not edit //
func main() {
	runnable.Use(HelloWorld{})
}
