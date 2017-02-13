package main

import (
	"fmt"

	"github.com/golang/go/src/sort"
)

func SortGophersNew(gophers []Gopher) {
	sortByHeight := func(i, j int) bool { return gophers[i].height < gophers[i].height } // HL
	sort.Slice(gophers, sortByHeight)                                                    // HL
	fmt.Println(gophers)
}
