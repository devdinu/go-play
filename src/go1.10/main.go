package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	r := strings.NewReader("123456789")
	data := make([]byte, 5)

	r.Read(data)
	w.Write(data)
	fmt.Printf("Size writer: %d, reader: %d\n", w.Size(), r.Size()) // HL
	encode()                                                        //OMIT
}
