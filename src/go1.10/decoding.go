package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type gopher struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func encode() {
	data := `{"id":1,"name":"lilgopher", "x": 1}`

	g := gopher{}
	decoder := json.NewDecoder(strings.NewReader(data))
	// With DisallowUnknownFields option it will panic when we run this or remove 'x' field from data
	// decoder.DisallowUnknownFields()
	if err := decoder.Decode(&g); err != nil {
		panic(err)
	}
	fmt.Printf("encoded: %+v", g)
}
