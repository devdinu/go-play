package main

import (
	"errors"
	"fmt"
)

// Function slices the src of 0:slicelen and adds elem 500 to new slice
func WierdSlice(src []int, slicelen int) ([]int, error) {
	if slicelen > len(src) {
		return nil, errors.New("invalid slice len")
	}
	sl := src[0:slicelen] // HL
	fmt.Println("new")
	inspect(sl)

	sl = append(sl, 500)
	fmt.Println("original")
	inspect(src)

	fmt.Println("new after append")
	inspect(sl)

	return sl, nil
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	res, _ := WierdSlice(data, 3)
	fmt.Printf("original: %v, sliced: %v", data, res)
}

func inspect(s []int) {
	fmt.Printf("cap %v, len %v, start address: %p\n", cap(s), len(s), s)
}
