package main

import (
	"github.com/RaphaelPour/aoc2020/util"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotate(t *testing.T) {
	x, y := Rotate(0, 1, 0)
	require.Equal(t, 0, x)
	require.Equal(t, 1, y)

	x, y = Rotate(0, 1, 90)
	require.Equal(t, -1, x)
	require.Equal(t, 0, y)

	x, y = Rotate(0, 1, -90)
	require.Equal(t, 1, x)
	require.Equal(t, 0, y)

	x, y = Rotate(10, 4, -90)
	require.Equal(t, 4, x)
	require.Equal(t, -10, y)
}

func Rotate(x, y, deg int) (int, int) {
	return x*int(math.Cos(util.Radians(float64(deg)))) -
			y*int(math.Sin(util.Radians(float64(deg)))),
		y*int(math.Cos(util.Radians(float64(deg)))) +
			x*int(math.Sin(util.Radians(float64(deg))))
}
