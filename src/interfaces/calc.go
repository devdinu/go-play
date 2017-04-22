package main

import "strconv"

type Calculator interface {
	Add(int, int) string
	Sub(int, int) string
}
type calc struct {
}

func (c calc) Add(a, b int) string {
	return strconv.Itoa(a + b)
}

func (c calc) Sub(a, b int) string {
	return strconv.Itoa(a - b)
}

func NewCalculator() Calculator {
	return calc{}
	//return appendCalc{}
}
