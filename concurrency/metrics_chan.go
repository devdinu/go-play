package concurrency

import (
	"fmt"
)

type Counter struct {
	occurence chan string
	data      map[string]int
	stop      <-chan bool
}

func (c *Counter) Incr(key string) {
	c.occurence <- key
}

func (c *Counter) process() {
	for {
		select {
		case key := <-c.occurence:
			c.data[key]++
		case <-c.stop:
			return
		}
	}
}

func (c *Counter) Report() {
	for k, v := range c.data {
		fmt.Printf("%s: %d", k, v)
	}
}

func (c *Counter) Get(key string) int {
	return c.data[key]
}

func NewCounter(close <-chan bool) *Counter {
	c := &Counter{
		data:      make(map[string]int, 1000),
		occurence: make(chan string),
		stop:      close,
	}
	go c.process()
	return c
}

func (c *Counter) Close() {
	close(c.occurence)
}
