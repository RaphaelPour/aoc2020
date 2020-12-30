package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"math"
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
	tiles       map[int]*Tile
	arrangement [][]Tile
}

func NewPuzzle() Puzzle {
	p := Puzzle{}
	p.tiles = make(map[int]*Tile, 0)
	return p
}

func (p Puzzle) Resolution() (int, error) {
	/* It can be assumed that all puzzles are squares. The resolution is the
	 * square root of the total count of tiles. The result must have no
	 * decimals. Otherwise making a square out of the tiles is not possible
	 */
	rawResolution := math.Sqrt(float64(len(p.tiles)))
	if rawResolution != float64(int64(rawResolution)) {
		return 0, fmt.Errorf(
			"Bad resolution. Expected square, got sqrt(%d)=%f != %f",
			len(p.tiles),
			rawResolution,
			float64(int64(rawResolution)),
		)
	}
	return int(rawResolution), nil
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

func (p Puzzle) PrintPuzzle() error {
	resolution, err := p.Resolution()
	if err != nil {
		return fmt.Errorf("Error printing puzzle: %s", err)
	}

	/* Render multiple tiles on the same line depending on the width
	 * parameter to make the puzzle observable
	 */

	keys := p.Keys()
	/* Go through each key and overstep width-many tiles */
	for i := 0; i < len(keys); i += resolution {
		/* Go through the full height of a tile */
		for y := 0; y < p.tiles[keys[i]].image.Height(); y++ {
			/* Go through all tiles which should be on the same line */
			for j := i; j < i+resolution; j++ {
				currentImage := p.tiles[keys[j]].image
				fmt.Print(currentImage.data[y])
				fmt.Print(" ")
			}
			fmt.Println("")
		}
		fmt.Println("")
	}

	return nil
}

func (p Puzzle) PrintArrangement() {
	for y := 0; y < len(p.arrangement); y++ {
		for row := 0; row < p.arrangement[y][0].image.Height(); row++ {
			for x := 0; x < len(p.arrangement[y]); x++ {
				fmt.Print(p.arrangement[y][x].image.data[row])
				fmt.Print(" ")
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
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

func (p Puzzle) ValidateNeighbourCount() error {
	corners, edges, centerParts := 0, 0, 0
	bad := make([]int, 0)
	for _, tile := range p.tiles {
		if tile.IsCorner() {
			corners++
		} else if tile.IsEdge() {
			edges++
		} else if tile.IsCenterPart() {
			centerParts++
		} else {
			bad = append(bad, len(tile.neighbours))
		}
	}

	if len(bad) > 0 {
		return fmt.Errorf(
			"%d tiles with bad neighbour count:%v (c=%d,e=%d,cp=%d,b=%d)",
			len(bad),
			bad,
			corners,
			edges,
			centerParts,
			len(bad),
		)
	}

	if corners != 4 {
		return fmt.Errorf(
			"Bad corner count. Expected 4, got %d (c=%d,e=%d,cp=%d,b=%d)",
			corners,
			corners,
			edges,
			centerParts,
			len(bad),
		)
	}

	resolution, err := p.Resolution()
	if err != nil {
		return fmt.Errorf("Error getting resolution: %s", err)
	}

	optimalEdgeCount := 4 * (resolution - 2)
	if edges != optimalEdgeCount {
		return fmt.Errorf(
			"Bad edge count. Expected %d, got %d (c=%d,e=%d,cp=%d,b=%d)",
			optimalEdgeCount,
			edges,
			corners,
			edges,
			centerParts,
			len(bad),
		)
	}

	/*
	 *|CenterParts| = (resolution-2)^2
	 * This calculates the count of tiles without the corners/edges.
	 */
	optimalCenterPartsCount := util.Pow(resolution-2, 2)
	if centerParts != optimalCenterPartsCount {
		return fmt.Errorf(
			"Bad center part count. Expected %d, got %d (c=%d,e=%d,cp=%d,b=%d)",
			optimalCenterPartsCount,
			centerParts,
			corners,
			edges,
			centerParts,
			len(bad),
		)
	}

	return nil
}

func (p *Puzzle) Arrange() error {

	/* Initialize final arrangement by allocating the 2D array where all
	 * tiles should be placed in in the right order and orientation.
	 */
	resolution, err := p.Resolution()
	if err != nil {
		return fmt.Errorf("Error getting resolution: %s", err)
	}

	arrangement := make([][]int, resolution)

	for i := range arrangement {
		arrangement[i] = make([]int, resolution)
	}

	/* Find the neighbours of each tile and validate the result afterwards.*/
	p.FindNeighbours()
	if err := p.ValidateNeighbourCount(); err != nil {
		return fmt.Errorf("Neighbour validation failed: %s", err)
	}

	fmt.Println("OK")

	return nil
}
