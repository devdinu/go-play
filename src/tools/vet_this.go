package main

import "fmt"

func vet() {
	val := 1
	fmt.Printf("Testing Go Vet Errors %s", val)
	suspiciousIf(val)
	some := "stringval"
	some = "other"
	fmt.Println(some)
	if true {
		return
	}
}

func suspiciousIf(x int) {
	if x != 1 || x != 2 {
		fmt.Println("This will be always True!!!")
	}
}
