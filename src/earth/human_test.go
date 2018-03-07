package earth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var pets []Gopher
var tallestGopher = Gopher{"tallest", 2.0, 250}
var petWorld *petWorldMock

func init() {
	petWorld = new(petWorldMock)
	pets = []Gopher{
		tallestGopher,
		Gopher{"secondTallest", 1.5, 150},
	}
}

func TestTallestPetOfHuman(t *testing.T) {
	h := human{petWorld, "homosapien", pets}
	petWorld.On("Tallest", pets).Return(Gopher{}, nil)

	_, err := h.tallestPet()

	assert.NoError(t, err)
	petWorld.AssertExpectations(t)
	// assert.Equal(t, tallestGopher, pet)
}

func TestTallestWithoutMocksLongerTimeTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Longer Test....")
	}
	time.Sleep(2 * time.Second)
	assert.Equal(t, "y", "y")
}

func TestRace(t *testing.T) {
	h := human{petWorld, "homosapien", pets}
	h.race()
}
