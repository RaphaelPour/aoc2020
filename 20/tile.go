package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"math"
	"sort"
	"strings"
)

type Tile struct {
	id         int
	image      Image
	neighbours []int
}

func NewTileFromImage(img *Image) Tile {
	tile := Tile{}
	tile.image = *img
	tile.neighbours = []int{-1, -1, -1, -1}
	return tile
}

func (t *Tile) AddNeighbour(id, direction int) error {
	/* Prevent adding the tile itself */
	if t.id == id {
		return fmt.Errorf(
			"Can't add the tile itself (%d)",
			id,
		)
	}

	if !util.InRange(direction, 0, 3) {
		return fmt.Errorf(
			"Error adding neighbour. Direction %d is invalid",
			direction,
		)
	}

	if t.neighbours[direction] != -1 {
		return fmt.Errorf(
			"Neighbour %s is already set to %d",
			DirString(direction),
			t.neighbours[direction],
		)
	}
	t.neighbours[direction] = id
	return nil
}

func (t Tile) AreNeighbours(other Tile) int {
	for i, sideA := range t.image.Sides() {
		for _, sideB := range other.image.Sides() {
			if sideA.Equals(sideB) {
				return i
			}
			sideB.Flip()
			if sideA.Equals(sideB) {
				return i
			}
		}
	}
	return -1
}

func (t Tile) NeighbourCount() int {
	count := 0
	for _, ID := range t.neighbours {
		if ID != -1 {
			count++
		}
	}
	return count
}

func (t Tile) IsCorner() bool {
	/*
	 * A corner has exactly two neighbours and is otherwise the most outer
	 * tile on the two other sides.
	 */
	return t.NeighbourCount() == 2
}

func (t Tile) IsEdge() bool {
	/*
	 * An edge has exactly three neighbours and is otherwise the most
	 * outer tile on one side.
	 */
	return t.NeighbourCount() == 3
}

func (t Tile) IsCenterPart() bool {
	/*
	 * A center part is no outer most tile on any side and has therefore
	 * four neighbours.
	 */
	return t.NeighbourCount() == 4
}

func (t Tile) HasValidNeighbourCount() bool {
	/*
	 * A tile must be either corner, edge or center part.
	 * Anything else should be treated as error.
	 */
	return util.InRange(t.NeighbourCount(), 2, 4)
}

type Puzzle struct {
	tiles       map[int]*Tile
	arrangement [][]*Tile
	corners     map[int]*Tile
	edges       map[int]*Tile
	centerParts map[int]*Tile
}

func NewPuzzle() Puzzle {
	p := Puzzle{}
	p.tiles = make(map[int]*Tile, 0)
	p.corners = make(map[int]*Tile, 0)
	p.edges = make(map[int]*Tile, 0)
	p.centerParts = make(map[int]*Tile, 0)
	return p
}

func (p Puzzle) AreAllTilesArranged() bool {
	return len(p.corners)+len(p.edges)+len(p.centerParts) == 0
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
	width, height := 0, 0
	for _, tile := range p.tiles {
		width = tile.image.Width()
		height = tile.image.Height()
		break
	}
	for y := 0; y < len(p.arrangement); y++ {
		for row := 0; row < height; row++ {
			for x := 0; x < len(p.arrangement[y]); x++ {

				/* If tile is empty, fill space with spaces */
				if p.arrangement[y][x] == nil {
					fmt.Print(
						strings.Repeat(
							" ",
							width,
						),
					)
				} else {
					fmt.Print(p.arrangement[y][x].image.data[row])
				}
				fmt.Print(" ")
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
}

func (p *Puzzle) FindNeighbours() error {

	keys := p.Keys()

	/* Iterate over all unique pairs of tiles */
	for i := 0; i < len(keys); i++ {
		for j := 0; j < len(keys); j++ {
			if i == j {
				continue
			}
			tileA := p.tiles[keys[i]]
			tileB := p.tiles[keys[j]]
			if dir := tileA.AreNeighbours(*tileB); dir != -1 {
				fmt.Printf(
					"Adding %d %s to %d\n",
					tileA.id,
					DirString(dir),
					tileB.id,
				)
				if err := tileA.AddNeighbour(tileB.id, dir); err != nil {
					return fmt.Errorf(
						"Error adding neighbour %d %s of %d",
						tileB.id,
						DirString(dir),
						tileA.id,
					)
				}
			}
		}
	}

	for _, tile := range p.tiles {
		if tile.IsCorner() {
			p.corners[tile.id] = tile
		} else if tile.IsEdge() {
			p.edges[tile.id] = tile
		} else if tile.IsCenterPart() {
			p.centerParts[tile.id] = tile
		}
	}
	return nil
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
			fmt.Printf("Found bad tile %d: %v\n", tile.id, tile.neighbours)
			bad = append(bad, tile.id)
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

	p.arrangement = make([][]*Tile, resolution)

	for i := range p.arrangement {
		p.arrangement[i] = make([]*Tile, resolution)
	}

	/* Find the neighbours of each tile and validate the result afterwards.*/
	p.FindNeighbours()
	if err := p.ValidateNeighbourCount(); err != nil {
		return fmt.Errorf("Neighbour validation failed: %s", err)
	}

	/* Start with an arbitrary corner */
	for key, tile := range p.corners {
		p.arrangement[0][0] = tile
		fmt.Printf("Arrange %d to 0/0\n", tile.id)
		delete(p.corners, key)
		break
	}

	/* Set origin of first corner so neighbours are Right/Top */
	for i := 0; i < 16; i++ {
		newImg := p.arrangement[0][0].image.Transform(i)
		if p.arrangement[0][0].neighbours[RIGHT] != -1 &&
			p.arrangement[0][0].neighbours[BOTTOM] != -1 {
			p.arrangement[0][0].image = newImg
			break
		}
	}

	/* Do edges/corners first */
	x, y := 0, 0
	for len(p.corners)+len(p.edges) > 0 {

		/* Reset x if end of a row has been reached */
		if x == resolution {
			x = 0
			y++
		}

		referenceTile := p.arrangement[y][x]
		for neighbourDir, neighbourID := range referenceTile.neighbours {
			if neighbourID == -1 {
				continue
			}

			neighbourTile, ok := p.tiles[neighbourID]
			if !ok {
				return fmt.Errorf("Neighbour tile %d doesn't exist", neighbourID)
			}
			/* Skip center parts and only consider corners and edges */
			if neighbourTile.IsCenterPart() {
				continue
			}
			transformedImg, dir := referenceTile.image.TransformOtherUntilMatch(neighbourTile.image)
			if transformedImg == nil {
				return fmt.Errorf(
					"Error transforming %d until it matches %d",
					referenceTile.id,
					neighbourID,
				)
			}

			neighbourTile.image = *transformedImg

			switch dir {
			case LEFT:
				x--
			case RIGHT:
				x++
			case TOP:
				y++
			case BOTTOM:
				y--
			}

			fmt.Printf("Arrange %d to %d/%d\n", neighbourID, x, y)
			p.arrangement[y][x] = neighbourTile

			/* Remove corner/edge from its source map */
			if neighbourTile.IsCorner() {
				delete(p.corners, neighbourID)
			} else if neighbourTile.IsEdge() {
				delete(p.edges, neighbourID)
			} else {
				return fmt.Errorf(
					"Can't remove tile %d from source map."+
						"It's neither a corner nor an edge",
					neighbourID,
				)
			}

			/* Remove reference/neighbour tile from each others neighbours */
			referenceTile.neighbours[neighbourDir] = -1
			neighbourTile.neighbours[OppositeDirection(neighbourDir)] = -1
			break
		}
	}
	p.PrintArrangement()

	fmt.Println("DONE")

	return nil
}
