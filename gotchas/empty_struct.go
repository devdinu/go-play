package main

import "fmt"
import "unsafe"

type T struct{}

func (t T) Bla() { fmt.Printf("bla %p\n", &t) }

func main() {
	var x [1000000000]T
	x[0] = T{}
	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(len(x), cap(x))
	fmt.Println(x[0], x[1])
	x[0].Bla()
	x[1].Bla()
	x[2].Bla()
}
