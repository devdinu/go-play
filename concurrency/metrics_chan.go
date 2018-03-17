package concurrency

import (
	"fmt"
)

type counter struct {
	occurence chan string
	data      map[string]int
}

func (c *counter) Incr(key string) {
	c.occurence <- key
}

func (c *counter) process() {
	//fmt.Println("processing...")
	for {
		key := <-c.occurence
		//		fmt.Println("received data")
		if key == "END" {
			return
		}
		c.data[key] += 1
	}
}

func (c *counter) Report() {
	for k, v := range c.data {
		fmt.Sprintf("%s: %d", k, v)
	}
}

func (c *counter) Get(key string) int {
	return c.data[key]
}

func NewCounter() *counter {
	c := &counter{
		data:      make(map[string]int, 1000),
		occurence: make(chan string),
	}
	go c.process()
	return c
}

func (c *counter) Close() {
	close(c.occurence)
}
