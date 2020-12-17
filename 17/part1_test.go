package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNeighbours1(t *testing.T) {

	lines := []string{
		"...",
		"...",
		"...",
	}

	grid, err := NewGridFromLines(lines)
	require.Nil(t, err)
	/* 1st row */
	require.Equal(t, 0, grid.Neighbours(0, 0, 0))
	require.Equal(t, 0, grid.Neighbours(1, 0, 0))
	require.Equal(t, 0, grid.Neighbours(2, 0, 0))

	/* 2nd row */
	require.Equal(t, 0, grid.Neighbours(0, 1, 0))
	require.Equal(t, 0, grid.Neighbours(1, 1, 0))
	require.Equal(t, 0, grid.Neighbours(2, 1, 0))

	/* 3rd row */
	require.Equal(t, 0, grid.Neighbours(0, 2, 0))
	require.Equal(t, 0, grid.Neighbours(1, 2, 0))
	require.Equal(t, 0, grid.Neighbours(2, 2, 0))
}

func TestNeighbours2(t *testing.T) {

	lines := []string{
		"...",
		".#.",
		"...",
	}

	grid, err := NewGridFromLines(lines)
	require.Nil(t, err)

	/* 1st row */
	require.Equal(t, 1, grid.Neighbours(0, 0, 0))
	require.Equal(t, 1, grid.Neighbours(1, 0, 0))
	require.Equal(t, 1, grid.Neighbours(2, 0, 0))

	/* 2nd row */
	require.Equal(t, 1, grid.Neighbours(0, 1, 0))
	require.Equal(t, 0, grid.Neighbours(1, 1, 0))
	require.Equal(t, 1, grid.Neighbours(2, 1, 0))

	/* 3rd row */
	require.Equal(t, 1, grid.Neighbours(0, 2, 0))
	require.Equal(t, 1, grid.Neighbours(1, 2, 0))
	require.Equal(t, 1, grid.Neighbours(2, 2, 0))
}

func TestNeighbours3(t *testing.T) {

	lines := []string{
		"...",
		"##.",
		"...",
	}

	grid, err := NewGridFromLines(lines)
	require.Nil(t, err)

	/* 1st row */
	require.Equal(t, 2, grid.Neighbours(0, 0, 0))
	require.Equal(t, 2, grid.Neighbours(1, 0, 0))
	require.Equal(t, 1, grid.Neighbours(2, 0, 0))

	/* 2nd row */
	require.Equal(t, 1, grid.Neighbours(0, 1, 0))
	require.Equal(t, 1, grid.Neighbours(1, 1, 0))
	require.Equal(t, 1, grid.Neighbours(2, 1, 0))

	/* 3rd row */
	require.Equal(t, 2, grid.Neighbours(0, 2, 0))
	require.Equal(t, 2, grid.Neighbours(1, 2, 0))
	require.Equal(t, 1, grid.Neighbours(2, 2, 0))
}

func TestNeighbours4(t *testing.T) {

	lines := []string{
		".#.",
		"##.",
		"...",
	}

	grid, err := NewGridFromLines(lines)
	require.Nil(t, err)

	/* 1st row */
	require.Equal(t, 3, grid.Neighbours(0, 0, 0))
	require.Equal(t, 2, grid.Neighbours(1, 0, 0))
	require.Equal(t, 2, grid.Neighbours(2, 0, 0))

	/* 2nd row */
	require.Equal(t, 2, grid.Neighbours(0, 1, 0))
	require.Equal(t, 2, grid.Neighbours(1, 1, 0))
	require.Equal(t, 2, grid.Neighbours(2, 1, 0))

	/* 3rd row */
	require.Equal(t, 2, grid.Neighbours(0, 2, 0))
	require.Equal(t, 2, grid.Neighbours(1, 2, 0))
	require.Equal(t, 1, grid.Neighbours(2, 2, 0))
}

func TestNeighbours5(t *testing.T) {

	lines := []string{
		".#.",
		"###",
		"...",
	}

	grid, err := NewGridFromLines(lines)
	require.Nil(t, err)

	/* 1st row */
	require.Equal(t, 3, grid.Neighbours(0, 0, 0))
	require.Equal(t, 3, grid.Neighbours(1, 0, 0))
	require.Equal(t, 3, grid.Neighbours(2, 0, 0))

	/* 2nd row */
	require.Equal(t, 2, grid.Neighbours(0, 1, 0))
	require.Equal(t, 3, grid.Neighbours(1, 1, 0))
	require.Equal(t, 2, grid.Neighbours(2, 1, 0))

	/* 3rd row */
	require.Equal(t, 2, grid.Neighbours(0, 2, 0))
	require.Equal(t, 3, grid.Neighbours(1, 2, 0))
	require.Equal(t, 2, grid.Neighbours(2, 2, 0))
}

func TestNeighboursExample(t *testing.T) {

	lines := []string{
		".#.",
		"..#",
		"###",
	}

	grid, err := NewGridFromLines(lines)
	require.Nil(t, err)

	/* 1st row */
	require.Equal(t, 1, grid.Neighbours(0, 0, 0))
	require.Equal(t, 1, grid.Neighbours(1, 0, 0))
	require.Equal(t, 2, grid.Neighbours(2, 0, 0))

	/* 2nd row */
	require.Equal(t, 3, grid.Neighbours(0, 1, 0))
	require.Equal(t, 5, grid.Neighbours(1, 1, 0))
	require.Equal(t, 3, grid.Neighbours(2, 1, 0))

	/* 3rd row */
	require.Equal(t, 1, grid.Neighbours(0, 2, 0))
	require.Equal(t, 3, grid.Neighbours(1, 2, 0))
	require.Equal(t, 2, grid.Neighbours(2, 2, 0))
}
