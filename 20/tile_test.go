package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNeighbour1(t *testing.T) {
	img1 := NewImage()
	img1.AddRow("#..#")
	img1.AddRow("#...")
	img1.AddRow("#..#")
	img1.AddRow("#...")

	img2 := NewImage()
	img2.AddRow("####")
	img2.AddRow(".###")
	img2.AddRow("####")
	img2.AddRow("###.")

	t1 := NewTileFromImage(&img1)
	t2 := NewTileFromImage(&img2)

	require.True(t, t1.AreNeighbours(t2))
}

func TestNeighbour2(t *testing.T) {
	img1 := NewImage()
	img1.AddRow("....")
	img1.AddRow("....")
	img1.AddRow("....")
	img1.AddRow("....")

	img2 := NewImage()
	img2.AddRow("####")
	img2.AddRow("####")
	img2.AddRow("####")
	img2.AddRow("####")

	t1 := NewTileFromImage(&img1)
	t2 := NewTileFromImage(&img2)

	require.False(t, t1.AreNeighbours(t2))
}

func TestNeighbour3(t *testing.T) {
	img1 := NewImage()
	img1.AddRow("####")
	img1.AddRow("....")
	img1.AddRow("....")
	img1.AddRow("....")

	img2 := NewImage()
	img2.AddRow("....")
	img2.AddRow("....")
	img2.AddRow("....")
	img2.AddRow("####")

	t1 := NewTileFromImage(&img1)
	t2 := NewTileFromImage(&img2)

	require.True(t, t1.AreNeighbours(t2))
}

func TestFindNeighbours(t *testing.T) {

	img1 := NewImage()
	img1.AddRow("##########")
	img1.AddRow("#........#")
	img1.AddRow("..........")
	img1.AddRow("..........")
	img1.AddRow("..........")
	img1.AddRow("..........")
	img1.AddRow("..........")
	img1.AddRow("..........")
	img1.AddRow("..........")
	img1.AddRow(".........#")

	img2 := NewImage()
	img2.AddRow("##########")
	img2.AddRow("#........#")
	img2.AddRow(".........#")
	img2.AddRow(".........#")
	img2.AddRow(".........#")
	img2.AddRow(".........#")
	img2.AddRow(".........#")
	img2.AddRow(".........#")
	img2.AddRow("..........")
	img2.AddRow("###......#")

	img3 := NewImage()
	img3.AddRow(".........#")
	img3.AddRow("#........#")
	img3.AddRow("#.........")
	img3.AddRow(".........#")
	img3.AddRow("#.........")
	img3.AddRow("#........#")
	img3.AddRow("#.........")
	img3.AddRow(".........#")
	img3.AddRow("..........")
	img3.AddRow(".###...###")

	img4 := NewImage()
	img4.AddRow("###......#")
	img4.AddRow("#.........")
	img4.AddRow("..........")
	img4.AddRow("#.........")
	img4.AddRow("..........")
	img4.AddRow("#........#")
	img4.AddRow(".........#")
	img4.AddRow("#........#")
	img4.AddRow(".........#")
	img4.AddRow("####.....#")

	t1 := NewTileFromImage(&img1)
	t1.id = 1
	t2 := NewTileFromImage(&img2)
	t2.id = 2
	t3 := NewTileFromImage(&img3)
	t3.id = 3
	t4 := NewTileFromImage(&img4)
	t4.id = 4

	puzzle := Puzzle{
		tiles: map[int]*Tile{
			1: &t1,
			2: &t2,
			3: &t3,
			4: &t4,
		},
	}

	require.Equal(t, []int{1, 2, 3, 4}, puzzle.Keys())

	puzzle.FindNeighbours()
	require.True(t, puzzle.tiles[1].IsCorner())
	require.True(t, puzzle.tiles[2].IsCorner())
	require.True(t, puzzle.tiles[3].IsCorner())
	require.True(t, puzzle.tiles[4].IsCorner())

	require.Equal(t, []int{2, 3}, puzzle.tiles[1].neighbours)
	require.Equal(t, []int{1, 4}, puzzle.tiles[2].neighbours)
	require.Equal(t, []int{1, 4}, puzzle.tiles[3].neighbours)
	require.Equal(t, []int{2, 3}, puzzle.tiles[4].neighbours)

	require.False(t, puzzle.tiles[1].IsEdge())
	require.False(t, puzzle.tiles[2].IsEdge())
	require.False(t, puzzle.tiles[3].IsEdge())
	require.False(t, puzzle.tiles[4].IsEdge())

	require.True(t, puzzle.tiles[1].HasValidNeighbourCount())
	require.True(t, puzzle.tiles[2].HasValidNeighbourCount())
	require.True(t, puzzle.tiles[3].HasValidNeighbourCount())
	require.True(t, puzzle.tiles[4].HasValidNeighbourCount())

	require.False(t, puzzle.tiles[1].IsCenterPart())
	require.False(t, puzzle.tiles[2].IsCenterPart())
	require.False(t, puzzle.tiles[3].IsCenterPart())
	require.False(t, puzzle.tiles[4].IsCenterPart())
}

func TestAddNeighbour(t *testing.T) {
	tile := Tile{id: 0}

	/* Neighbours should be initialized empty */
	require.Empty(t, tile.neighbours)

	/* Neighbours should be still empty on trying add the tile itself */
	tile.AddNeighbour(0)
	require.Empty(t, tile.neighbours)

	/* Add a valid neighbour */
	tile.AddNeighbour(1)
	require.Equal(t, []int{1}, tile.neighbours)

	/* Duplicate neighbours should be avoided */
	tile.AddNeighbour(1)
	require.Equal(t, []int{1}, tile.neighbours)
}
