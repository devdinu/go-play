package main

import (
	"fmt"
	"sync"
)

type T struct {
	lock sync.Mutex
	val  int
}

func (t *T) Lock() {
	t.lock.Lock()
}

func (t *T) Change() {
	t.val = 100
}

func (t T) Unlock() {
	t.lock.Unlock()
}

func main() {
	t := T{lock: sync.Mutex{}}
	t.Lock()
	fmt.Println("locked...")
	t.Unlock()
	fmt.Println("unlocked...")
	//	t.Lock()
	t.Change()
	fmt.Println("done...", t.val)
}
