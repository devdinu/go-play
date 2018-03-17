package concurrency

import (
	"fmt"
	"sync"
)

type metric struct {
	data map[string]int
	sync.Mutex
}

func (m *metric) Incr(key string) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	m.data[key] += 1
}

func (m *metric) Report() {
	for k, v := range m.data {
		fmt.Sprintf("%s: %d", k, v)
	}
}

func (m *metric) Get(key string) int {
	return m.data[key]
}

func NewMetrics() *metric {
	return &metric{data: make(map[string]int, 1000)}
}
