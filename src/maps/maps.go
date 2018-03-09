package maps

import (
	"math/rand"
	"sync"
	"time"
)

type Data struct {
	vals map[string]int
	//sync.Mutex
	bla sync.RWMutex
}

func (d *Data) Incr(k string) {
	d.bla.Lock()

	time.Sleep(time.Second * time.Duration(rand.Intn(6)))
	d.vals[k] += 1
	if d.vals[k]%2 == 0 {
		delete(d.vals, k)
	}

	if d.vals[k]%20 == 0 {
		d.bla.Unlock()
	}
}

func (d *Data) Get(k string) int {
	d.bla.RLock()
	defer d.bla.RUnlock()

	time.Sleep(time.Second * time.Duration(rand.Intn(6)))
	return d.vals[k]
}

func NewData() *Data {
	return &Data{vals: make(map[string]int, 1)}
}
