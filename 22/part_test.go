package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayer1(t *testing.T) {
	p1 := NewPlayer()
	p1.cards = []int{1, 2, 3}
	p1.PutFirstCardToBack()

	require.Equal(t, []int{2, 3, 1}, p1.cards)
}

func TestPlayer2(t *testing.T) {
	p1 := NewPlayer()
	p1.cards = []int{1, 2, 3}

	require.Equal(t, 1, p1.RemoveFirstCard())
	require.Equal(t, []int{2, 3}, p1.cards)
}
