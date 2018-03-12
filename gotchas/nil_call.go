package main

import "fmt"

type gopher struct {
	name string
}

func (g *gopher) Name() {
	fmt.Println("My name is gopher.")
}

func main() {
	// Thought the pointer is nil, its valid to call the method on it
	var g *gopher
	fmt.Printf("content: %+v %+p\n", g, g)
	g.Name()
}
