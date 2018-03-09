package main

import "fmt"

func main() {
	fmt.Println("Sending data to a channel")
	talks := make(chan string)
	talks <- "lightning!" // send data
	go func() {
		t := <-talks // receive data
		fmt.Println(t)
	}()
	fmt.Println("Done.")
}
