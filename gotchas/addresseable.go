package main

import "fmt"

type gopher struct {
	name string
}

func main() {
	gs := map[string]gopher{}
	gs["tall"] = gopher{"TallGopher"}
	fmt.Println("map: ", gs)
	// This will error out, and could have pointers
	//gs["tall"].name = "change" // HL
	fmt.Println("map: ", gs)

	gl := []gopher{{"tall"}, {"small"}}
	gl[0].name = "asdffds"
	fmt.Println(gl)
}
