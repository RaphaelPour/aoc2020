package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBin2Dec(t *testing.T) {
	require.Equal(t, 0, Bin2Dec("0", '1', '0'))
	require.Equal(t, 1, Bin2Dec("1", '1', '0'))
	require.Equal(t, 2, Bin2Dec("10", '1', '0'))
	require.Equal(t, 3, Bin2Dec("11", '1', '0'))
	require.Equal(t, 127, Bin2Dec("1111111", '1', '0'))
}

func TestInput(t *testing.T) {
	require.Equal(t, 44, Bin2Dec("FBFBBFF", 'B', 'F'))
	require.Equal(t, 5, Bin2Dec("RLR", 'R', 'L'))

	require.Equal(t, 70, Bin2Dec("BFFFBBF", 'B', 'F'))
	require.Equal(t, 7, Bin2Dec("RRR", 'R', 'L'))
	require.Equal(t, 567, BoardingID(70, 7))

	require.Equal(t, 14, Bin2Dec("FFFBBBF", 'B', 'F'))
	require.Equal(t, 119, BoardingID(14, 7))

	require.Equal(t, 102, Bin2Dec("BBFFBBF", 'B', 'F'))
	require.Equal(t, 4, Bin2Dec("RLL", 'R', 'L'))
	require.Equal(t, 820, BoardingID(102, 4))
}
