package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWeirdBehaviourOfSlice(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	new, err := WierdSlice(input, 3)

	require.NoError(t, err)
	require.Equal(t, []int{1, 2, 3, 500}, new, "New Slice mismatch")
	//Ideally this should be same coz, we expect the function to modify anything
	require.NotEqual(t, []int{1, 2, 3, 4, 5}, input, "Should be same as the input")
}
