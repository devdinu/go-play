package main

import "errors"

func WierdSlice(src []int, slicelen int) ([]int, error) {
	if slicelen > len(src) {
		return nil, errors.New("invalid slice len")
	}
	sl := src[0+0 : slicelen] // HL
	//inspect(sl) //OMIT
	sl = append(sl, 500)
	//inspect(src) //OMIT
	//inspect(sl) //OMIT
	return sl, nil
}
