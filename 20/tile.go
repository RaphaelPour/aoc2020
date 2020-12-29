package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"sort"
)

type Tile struct {
	id         int
	image      Image
	neighbours []int
}

func NewTileFromImage(img *Image) Tile {
	tile := Tile{}
	tile.image = *img
	tile.neighbours = make([]int, 0, 4)
	return tile
}

func (t *Tile) AddNeighbour(id int) {
	/* Prevent adding the tile itself */
	if t.id == id {
		return
	}

	/* Skip if already included */
	for _, currentID := range t.neighbours {
		if currentID == id {
			return
		}
	}

	t.neighbours = append(t.neighbours, id)
}

func (t Tile) AreNeighbours(other Tile) bool {
	for _, sideA := range t.image.Sides() {
		for _, sideB := range other.image.Sides() {
			if sideA.Equals(sideB) {
				return true
			}
			sideB.Flip()
			if sideA.Equals(sideB) {
				return true
			}
		}
	}
	return false
}

func (t Tile) IsCorner() bool {
	/*
	 * A corner has exactly two neighbours and is otherwise the most outer
	 * tile on the two other sides.
	 */
	return len(t.neighbours) == 2
}

func (t Tile) IsEdge() bool {
	/*
	 * An edge has exactly three neighbours and is otherwise the most
	 * outer tile on one side.
	 */
	return len(t.neighbours) == 3
}

func (t Tile) IsCenterPart() bool {
	/*
	 * A center part is no outer most tile on any side and has therefore
	 * four neighbours.
	 */
	return len(t.neighbours) == 4
}

func (t Tile) HasValidNeighbourCount() bool {
	/*
	 * A tile must be either corner, edge or center part.
	 * Anything else should be treated as error.
	 */
	return util.InRange(len(t.neighbours), 2, 4)
}

type Puzzle struct {
	tiles map[int]*Tile
}

func (p Puzzle) Keys() []int {
	keys := make([]int, len(p.tiles))

	index := 0
	for id, _ := range p.tiles {
		keys[index] = id
		index++
	}
	sort.Ints(keys)
	return keys
}

func (p *Puzzle) FindNeighbours() {

	keys := p.Keys()

	/* Iterate over all unique pairs of tiles */
	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			tileA := p.tiles[keys[i]]
			tileB := p.tiles[keys[j]]
			if tileA.AreNeighbours(*tileB) {
				tileA.AddNeighbour(tileB.id)
				tileB.AddNeighbour(tileA.id)
			}
		}
	}

	for i := 0; i < len(keys); i++ {
		fmt.Println(p.tiles[keys[i]].neighbours)
	}
}
