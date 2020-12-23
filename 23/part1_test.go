package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPop_1(t *testing.T) {

	c := Circle{cups: []int{1, 2, 3}}

	require.Equal(t, 1, c.Pop(0))
	require.Equal(t, []int{2, 3}, c.cups)

	require.Equal(t, 2, c.Pop(0))
	require.Equal(t, []int{3}, c.cups)

	require.Equal(t, 3, c.Pop(0))
	require.Empty(t, c.cups)
}

func TestPop_2(t *testing.T) {

	c := Circle{cups: []int{1, 2, 3}}

	require.Equal(t, 2, c.Pop(1))
	require.Equal(t, []int{1, 3}, c.cups)

	require.Equal(t, 3, c.Pop(1))
	require.Equal(t, []int{1}, c.cups)

}

func TestPop3_1(t *testing.T) {

	c := Circle{cups: []int{1, 2, 3}}

	require.Equal(t, []int{1, 2, 3}, c.Pop3(0))
	require.Empty(t, c.cups)
}

func TestPop3_2(t *testing.T) {

	c := Circle{cups: []int{1, 2, 3, 4, 5, 6}}

	require.Equal(t, []int{1, 2, 3}, c.Pop3(0))
	require.Equal(t, []int{4, 5, 6}, c.cups)

	require.Equal(t, []int{4, 5, 6}, c.Pop3(0))
	require.Empty(t, c.cups)
}

func TestPop3_3(t *testing.T) {

	c := Circle{cups: []int{1, 2, 3, 4}}

	require.Equal(t, []int{3, 4, 1}, c.Pop3(2))
	require.Equal(t, []int{2}, c.cups)
}

func TestPop3_4(t *testing.T) {

	c := Circle{cups: []int{1, 2, 3, 4}}

	require.Equal(t, []int{4, 1, 2}, c.Pop3(3))
	require.Equal(t, []int{3}, c.cups)
}

func TestInsert(t *testing.T) {
	c := Circle{cups: []int{1, 3, 4}}
	c.Insert(2, 1)
	require.Equal(t, []int{1, 2, 3, 4}, c.cups)
}

func TestInsert3(t *testing.T) {
	c := Circle{cups: []int{1, 5}}
	c.Insert3(1, []int{2, 3, 4})
	require.Equal(t, []int{1, 2, 3, 4, 5}, c.cups)
}
