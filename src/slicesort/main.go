package main

import "fmt"

func main() {
	gophers := []Gopher{Gopher{"tall", 0.7, 250},
		Gopher{"fat", 1.5, 150},
		Gopher{"cute", 0.5, 50}}
	fmt.Println("Old way of Sorting : Doing By Height")
	SortGophersOld(gophers)
	fmt.Println("New of Sorting : Doing By Height")
	SortGophersNew(gophers)
}
