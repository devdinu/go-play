package main

import (
	"errors"
	"fmt"
)

func errs() {
	num := 3
	res, _ := odd(num)
	fmt.Printf("number %d is %t", num, res)
	noErr()
}

func odd(v int) (bool, error) {
	if v%2 == 0 {
		return true, nil
	}
	return false, errors.New("odd number")
}

func noErr() error {
	return nil
}
