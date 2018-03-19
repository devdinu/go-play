package concurrency

import (
	"fmt"
	"sync"
)

type Metric struct {
	data map[string]int
	sync.Mutex
}

func (m *Metric) Incr(key string) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	m.data[key]++
}

func (m *Metric) Report() {
	for k, v := range m.data {
		fmt.Printf("%s: %d", k, v)
	}
}

func (m *Metric) Get(key string) int {
	return m.data[key]
}

func NewMetrics() *Metric {
	return &Metric{data: make(map[string]int, 1000)}
}
