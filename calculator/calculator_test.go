package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldAdd(t *testing.T) {
	c := new(Calculator)

	c.Add(3)

	assert.Equal(t, 3, c.Result())
}

func TestShouldSub(t *testing.T) {
	c := new(Calculator)
	c.Press(10)

	c.Sub(3)

	assert.Equal(t, 7, c.Result())
}

func TestPress(t *testing.T) {
	c := new(Calculator)

	c.Press(10)

	assert.Equal(t, 10, c.Result())
}
