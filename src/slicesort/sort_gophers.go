package main

import (
	"fmt"
	"sort"
)

func SortGophersByHeightNew(gophers []Gopher) {
	sortByHeight := func(i, j int) bool { return gophers[i].height < gophers[j].height } // HL
	sort.Slice(gophers, sortByHeight)                                                    // HL
	fmt.Println(gophers)
}
