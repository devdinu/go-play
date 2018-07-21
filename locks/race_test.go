package locks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocks(t *testing.T) {
	adders, subs := 700, 400
	res := doJobs(adders, subs)
	assert.Equal(t, adders-subs, res)
}
