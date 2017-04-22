package main

import "strconv"

type appendCalc struct {
}

func (ac appendCalc) Add(a, b int) string {
	return strconv.Itoa(a) + strconv.Itoa(b)
}

func (ac appendCalc) Sub(a, b int) string {
	return strconv.Itoa(a - b)
}
