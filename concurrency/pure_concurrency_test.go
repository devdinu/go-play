package concurrency

import (
	"testing"

	"github.com/bcicen/grmon"
	"github.com/stretchr/testify/require"
)

func BenchmarkSumConcurrency(b *testing.B) {
	grmon.Start()
	for i := 0; i < b.N; i++ {
		SumIt(1000, 500)
	}
}

func TestSumConcurrency(t *testing.T) {
	grmon.Start()
	n := 100
	es := n * (n + 1) / 2
	require.Equal(t, es, SumIt(n, 10))
}
