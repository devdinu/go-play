package main

import (
	"fmt"
)

func slices(total int) {
	fmt.Println("Slices: Cap and len")
	s := make([]int, 0, 3)
	for i := 0; i < total; i++ {
		s = append(s, i)
		fmt.Printf("cap %v, len %v, %p\n", cap(s), len(s), s)
	}
}

func getFirstEven(data []int) []*int {
	var result []*int
	for _, e := range data {
		fmt.Printf("%v\n data: %d\n", &e, e)
		if e%2 == 0 {
			fmt.Println("This is the place", e)
			result = append(result, &e)
		}
	}
	fmt.Println(result)
	return result
}

func slicing(sl []int) {
	fmt.Println("\nSlicing the slice:")
	fmt.Printf("%+v, %p\n", sl, &sl)
	slx := sl[2:]
	fmt.Printf("%+v, %p\n", slx, &slx)
}
