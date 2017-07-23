package main

import "fmt"

type Gopher struct {
	name string
}

func (g Gopher) Code() {
	fmt.Println("can code.")
}

type Dancer struct {
	*Gopher
}

func (d *Dancer) Dance() {
	fmt.Println("can dance")
}

type DancingGopher interface{}

func testGopherThing() {
	var g *Gopher
	g = &Gopher{"gopher"}
	dancer := Dancer{g}

	fmt.Println("base gopher...")
	g.Code()

	fmt.Println("\ngopher dancer...")
	dancer.Dance()
	dancer.Code()

	var dg DancingGopher
	dg = dancer

	extracted, ok := dg.(Gopher)
	if !ok {
		fmt.Println("oops can't cast to Gopher")
	}
	fmt.Println("cool extracted...")
	extracted.Code()
}
