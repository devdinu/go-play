package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFailureTallestGopherWithNoInput(t *testing.T) {
	_, err := tallest(nil)

	assert.EqualError(t, err, "no gophers!")
}

func TestFailureTallestGopherWithEmptySlice(t *testing.T) {
	_, err := tallest([]Gopher{})

	assert.EqualError(t, err, "no gophers!")
}

func TestShouldReturnGopherWhenOnlyOneIsGiven(t *testing.T) {
	g := Gopher{name: "tallest", height: 1.0}
	tg, err := tallest([]Gopher{g})

	assert.NoError(t, err)
	assert.Equal(t, g, tg)
}
