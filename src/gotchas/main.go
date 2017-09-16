package main

import "fmt"

func main() {
	//var x := nil
	fmt.Println("gotchas...")

	//arrayMutate()
	asyncServe()
}

func mutate(data [3]int) {
	data[0] = 1000
	fmt.Println("inside:", data)
}

func arrayMutate() {
	d := [3]int{1, 3, 5}
	mutate(d)
	fmt.Println("after:", d)
}
