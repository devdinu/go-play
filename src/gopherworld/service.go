package main

import "errors"

func tallest(gophers []Gopher) (Gopher, error) {
	return Gopher{}, errors.New("no gophers!")
}
