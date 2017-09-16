package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("gotchas...")
	arrayMutate()
	//asyncServe()
	playBool()
	checkReceivers()
}

func mutate(data [3]int) {
	for i, v := range data {

		go func(v int) {
			time.Sleep(2 * time.Second)
			fmt.Println("value:", v)
		}(v) // Should pass the values to goroutines
		data[i] *= 2

	}
	time.Sleep(2 * time.Second)
	data[0] = 1000
	fmt.Println("inside:", data)
}

func arrayMutate() {
	d := [3]int{1, 3, 5}
	mutate(d)
	fmt.Println("after:", d)
}
