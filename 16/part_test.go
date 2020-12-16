package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnion(t *testing.T) {
	require.Equal(t, []int{1, 2, 3}, Union([]int{2, 3}, []int{1, 2}))
	require.Equal(t, []int{3, 2, 1}, Union([]int{1, 3}, []int{3, 2}))
	require.Equal(t, []int{1, 2, 3}, Union([]int{1, 2, 3}, []int{}))
	require.Equal(t, []int{1, 2, 3}, Union([]int{}, []int{1, 2, 3}))

}
