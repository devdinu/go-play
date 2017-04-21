package gopherworld

import (
	"fmt"
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

func TestShouldReturnTallestAmongMultipleGophers(t *testing.T) {
	g1 := Gopher{name: "secondTallest", height: 1.0}
	g2 := Gopher{name: "tallest", height: 2.0}

	tg, err := tallest([]Gopher{g1, g2})

	if err != nil {
		fmt.Printf("Should Not Throw any Error, but got: %s", err.Error())
		t.FailNow()
	}
	expectedGopher := Gopher{name: "tallest", height: 2.0}
	if expectedGopher != tg {
		fmt.Println("Not Equal")
		fmt.Printf("Expected: %s", expectedGopher)
		fmt.Printf("Actual: %s", tg)
		t.FailNow()
	}
}
