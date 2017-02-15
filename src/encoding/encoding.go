package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

// START OMIT
func main() {
	data1 := "WvLTlMrX9NpYDQlEIFlnDA=="
	data2 := "WvLTlMrX9NpYDQlEIFlnDB=="
	showDecodedOutput(data1, data2)
	strictTest(data1, data2)
}

func showDecodedOutput(data1, data2 string) {
	val1, err1 := base64.StdEncoding.DecodeString(data1)
	val2, err2 := base64.StdEncoding.DecodeString(data2)
	fmt.Println("Data1:", string(val1))
	fmt.Println("Data2:", string(val2))
	fmt.Println("Same?:", bytes.Equal(val1, val2), "No errors?:", (err1 == nil) && (err2 == nil))
}

// END OMIT
func strictTest(data1, data2 string) {
	fmt.Println("Strict Decoding...")
	x, err1 := base64.StdEncoding.Strict().DecodeString(data1)
	y, err2 := base64.StdEncoding.Strict().DecodeString(data2)
	fmt.Println("Error1:", err1, "Error2:", err2)
	fmt.Println(x, y)
}
