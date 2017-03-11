package main

import "fmt"

type Gopher struct {
	name   string
	weight int
}

func lint() {
	someUrl := "http://google.com"
	fmt.Println("Linters check...")
	fmt.Println(someUrl)
}
