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

func TestHistory1(t *testing.T) {
	g := NewGame()

	g.p1.cards = []int{1, 2, 3}
	g.p2.cards = []int{4, 5, 6}

	require.False(t, g.IsTurnReocurring())

	g.AddTurnToHistory()

	require.Equal(t, g.history, map[string]bool{
		"[1 2 3][4 5 6]": true,
	})
	require.True(t, g.IsTurnReocurring())

}

func TestHistory2(t *testing.T) {
	g := NewGame()

	g.p1.cards = []int{1, 2, 3}
	g.p2.cards = []int{4, 5, 6}

	require.False(t, g.IsTurnReocurring())

	g.AddTurnToHistory()

	g.p1.cards = []int{1, 2}
	g.p2.cards = []int{3, 4, 5, 6}
	require.False(t, g.IsTurnReocurring())

	g.AddTurnToHistory()
	require.Equal(t, g.history, map[string]bool{
		"[1 2 3][4 5 6]": true,
		"[1 2][3 4 5 6]": true,
	})
}

func TestPlayerFirstCardToBack1(t *testing.T) {
	p := NewPlayer()
	p.cards = []int{1, 2, 3}

	p.PutFirstCardToBack()
	require.Equal(t, p.cards, []int{2, 3, 1})

	p.PutFirstCardToBack()
	require.Equal(t, p.cards, []int{3, 1, 2})

	p.PutFirstCardToBack()
	require.Equal(t, p.cards, []int{1, 2, 3})
}

func TestPlayerFirstCardToBack2(t *testing.T) {
	p := NewPlayer()
	p.cards = []int{1, 2, 3}

	p.PutFirstCardToBack()
	require.Equal(t, p.cards, []int{2, 3, 1})

	require.Equal(t, p.RemoveFirstCard(), 2)
	require.Equal(t, p.cards, []int{3, 1})

	p.PutFirstCardToBack()
	require.Equal(t, p.cards, []int{1, 3})

	require.Equal(t, p.RemoveFirstCard(), 1)
	require.Equal(t, p.cards, []int{3})

	p.PutFirstCardToBack()
	require.Equal(t, p.cards, []int{3})

	require.Equal(t, p.RemoveFirstCard(), 3)
	require.Empty(t, p.cards)

	p.PutFirstCardToBack()
	require.Empty(t, p.cards)

	require.Equal(t, p.RemoveFirstCard(), -1)
	require.Empty(t, p.cards)
}

func TestRemoveFirstCard(t *testing.T) {
	p := NewPlayer()
	p.cards = []int{1, 2, 3}

	require.Equal(t, p.RemoveFirstCard(), 1)
	require.Equal(t, p.cards, []int{2, 3})

	require.Equal(t, p.RemoveFirstCard(), 2)
	require.Equal(t, p.cards, []int{3})

	require.Equal(t, p.RemoveFirstCard(), 3)
	require.Empty(t, p.cards)

	require.Equal(t, p.RemoveFirstCard(), -1)
	require.Empty(t, p.cards)
}

func TestSubPlayer(t *testing.T) {
	p := NewPlayer()
	p.cards = []int{1, 2, 3}

	sp := p.NewSubPlayer(2)

	require.Equal(t, p.cards, []int{1, 2, 3})
	require.Equal(t, sp.cards, []int{2, 3})

	sp.PutFirstCardToBack()
	require.Equal(t, p.cards, []int{1, 2, 3})
	require.Equal(t, sp.cards, []int{3, 2})

	sp.RemoveFirstCard()
	require.Equal(t, p.cards, []int{1, 2, 3})
	require.Equal(t, sp.cards, []int{2})
}

func TestSubGame(t *testing.T) {
	g := NewGame()

	g.p1.cards = []int{1, 2, 3}
	g.p2.cards = []int{4, 5, 6}

	sg := g.NewSubGame(2, 2)

	require.Equal(t, sg.p1.cards, []int{2, 3})
	require.Equal(t, sg.p2.cards, []int{5, 6})
}

func TestCanGoRecursive(t *testing.T) {
	g := NewGame()

	g.p1.cards = []int{1, 2, 3}
	g.p2.cards = []int{4, 5, 6}

	require.True(t, g.p1.CanGoRecursive())
	require.False(t, g.p2.CanGoRecursive())
}
