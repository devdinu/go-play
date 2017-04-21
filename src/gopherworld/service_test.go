package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindingTallestGopherWithNoInput(t *testing.T) {
	_, err := tallest(nil)

	assert.EqualError(t, err, "no gophers!")
}
