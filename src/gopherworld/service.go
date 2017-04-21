package gopherworld

import (
	"errors"
	"sort"
)

func tallest(gophers []Gopher) (Gopher, error) {
	if gophers == nil || len(gophers) == 0 {
		return Gopher{}, errors.New("no gophers!")
	}
	byHeight := func(i, j int) bool {
		return gophers[i].height > gophers[j].height
	}
	sort.Slice(gophers, byHeight)
	return gophers[0], nil
}
