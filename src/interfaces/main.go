package main

import "fmt"

func main() {
	c := NewCalculator()
	added := c.Add(100, 50)
	sub := c.Sub(50, 10)
	fmt.Println("Add:", added, " Sub:", sub)
}
