package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/bcicen/grmon"
	"github.com/stretchr/testify/assert"
)

func TestConcurrentIncrementsChan(t *testing.T) {
	grmon.Start()
	calls := 10
	m := NewCounter()
	done := make(chan bool, 1)

	var wg sync.WaitGroup
	wg.Add(calls)

	for i := 0; i < calls; i++ {
		go func() {
			time.Sleep(2 * time.Second)
			m.Incr("incr.call")
			wg.Done()
		}()
	}

	wg.Wait()
	m.Incr("END")
	done <- true

	fmt.Println("wating for completion")
	<-done
	m.Close()

	result := m.Get("incr.call")
	fmt.Println("final calls with channels: ", result)
	assert.Equal(t, calls, result, "share by communicating failed")
}

func BenchmarkConcurrentIncrementChan(b *testing.B) {
	//CHAN
	c := NewCounter()
	for i := 0; i < b.N; i++ {
		c.Incr("increment.call")
	}
}

func BenchmarkConcurrentIncrementLock(b *testing.B) {
	//LOCK
	m := NewMetrics()
	for i := 0; i < b.N; i++ {
		m.Incr("increment.call")
	}
}
