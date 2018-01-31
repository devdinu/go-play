package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstEvenNumberResult(t *testing.T) {
	data := []int{3, 5, 10, 7, 9}

	result := getFirstEven(data)

	assert.Equal(t, 1, len(result))
	assert.Equal(t, 10, *result[0], "Why This Test Fails!")
}
