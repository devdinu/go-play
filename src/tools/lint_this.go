package main

import "fmt"

type Gopher struct {
	name   string
	weight int
}

func lint() {
	someUrl := "http://google.com" // HL
	fmt.Println("Linters check...")
	fmt.Println(someUrl)
}
