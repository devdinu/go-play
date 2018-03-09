package main

import (
	"errors"
	"fmt"
)

func WierdSlice(src []int, slicelen int) ([]int, error) {
	if slicelen > len(src) {
		return nil, errors.New("invalid slice len")
	}
	sl := src[0:slicelen] // HL
	fmt.Println("new")
	inspect(sl) //OMIT

	sl = append(sl, 500)
	fmt.Println("original")

	inspect(src) //OMIT
	fmt.Println("new after append")
	inspect(sl) //OMIT

	return sl, nil
}
