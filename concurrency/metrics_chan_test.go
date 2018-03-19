package concurrency

import (
	"fmt"
	"sync"
	"testing"

	"github.com/bcicen/grmon"
	"github.com/stretchr/testify/assert"
)

func TestConcurrentIncrementsChan(t *testing.T) {
	grmon.Start()
	calls := 10
	done := make(chan bool, 1)
	m := NewCounter(done)

	var wg sync.WaitGroup
	wg.Add(calls)

	for i := 0; i < calls; i++ {
		go func() {
			m.Incr("incr.call")
			wg.Done()
		}()
	}

	wg.Wait()
	done <- true

	fmt.Println("wating for completion")
	m.Close()

	result := m.Get("incr.call")
	fmt.Println("final calls with channels: ", result)
	assert.Equal(t, calls, result, "share by communicating failed")
}

func BenchmarkConcurrentIncrementChan(b *testing.B) {
	done := make(chan bool, 1)
	c := NewCounter(done)
	for i := 0; i < b.N; i++ {
		c.Incr("increment.call")
	}
	done <- true
}

func BenchmarkConcurrentIncrementLock(b *testing.B) {
	done := make(chan bool, 1)
	m := NewMetrics()
	for i := 0; i < b.N; i++ {
		m.Incr("increment.call")
	}
	done <- true
}
