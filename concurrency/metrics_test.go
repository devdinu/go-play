package concurrency

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkIncr(b *testing.B) {

}

func TestConcurrentIncrements(t *testing.T) {
	calls := 5
	m := NewMetrics()

	var wg sync.WaitGroup
	wg.Add(calls)

	for i := 0; i < calls; i++ {
		go func(wg *sync.WaitGroup) {
			m.Incr("incr.call")
			wg.Done()
		}(&wg)
	}

	wg.Wait()

	assert.Equal(t, calls, m.Get("incr.call"))
}
