package main

import "errors"

func tallest(gophers []Gopher) (Gopher, error) {
	if gophers == nil || len(gophers) == 0 {
		return Gopher{}, errors.New("no gophers!")
	}
	return gophers[0], nil
}
