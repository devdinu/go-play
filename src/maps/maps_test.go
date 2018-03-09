package maps

import (
	"fmt"
	"testing"
)

func BenchmarkParallelUpdate(b *testing.B) {
	d := NewData()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error occured")
		}
	}()

	for i := 0; i < b.N; i++ {
		go d.Incr("a")
		go d.Incr("b")
		go d.Incr("a")
		go d.Get("a")
		go d.Get("b")
		go d.Get("a")
	}

	fmt.Println(d.vals)

}
