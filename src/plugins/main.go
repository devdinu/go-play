package main

import (
	"fmt"
	"log"
	"plugin"
)

func main() {
	p, err := plugin.Open("./calculator.so")
	if err != nil {
		log.Fatalf("failed to open plugin calculator")
	}

	add, err := p.Lookup("Add")
	if err != nil {
		log.Fatalf("failed to open plugin calculator")
	}

	fmt.Println(add.(func(int, int) int)(1, 2))

	sub, err := p.Lookup("Sub")
	if err != nil {
		log.Fatalf("failed to open plugin calculator")
	}

	fmt.Println(sub.(func(int, int) int)(1, 2))
}
