package main

import (
	"fmt"
	"sort"
)

type ByWeight []Gopher

func (gs ByWeight) Len() int           { return len(gs) }
func (gs ByWeight) Swap(i, j int)      { gs[i], gs[j] = gs[j], gs[i] }
func (gs ByWeight) Less(i, j int) bool { return gs[i].weight < gs[j].weight }

func SortGophersOld(gophers []Gopher) {
	sort.Sort(ByWeight(gophers))
	fmt.Println(gophers)
}
