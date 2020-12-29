package main

import (
	"fmt"
	"regexp"

	"github.com/RaphaelPour/aoc2020/util"
)

var (
	inputFile = "input"
	dir2XY    = map[string][]int{
		"e":  {2, 0},
		"se": {1, -1},
		"sw": {-1, -1},
		"w":  {-2, 0},
		"nw": {-1, 1},
		"ne": {1, 1},
	}
)

const (
	INIT = iota
	FLIPPED
)

type Tile struct {
	x, y int
}

type Tiles map[Tile]int

func (t Tiles) String() string {
	out := ""

	minX, maxX, minY, maxY := t.Boundaries()

	for y := minY; y <= maxY; y++ {
		if y%2 == 0 {
			out += " "
		}
		for x := minX; x <= maxX; x++ {
			if tile, ok := t[Tile{x, y}]; ok {
				if tile == FLIPPED {
					out += "F"
				} else {
					out += "I"
				}
			} else {
				out += "-"
			}
		}
		out += "\n"
	}

	return out
}

func (t Tiles) Boundaries() (int, int, int, int) {
	if len(t) == 0 {
		return 0, 0, 0, 0
	}
	minX, minY := 100, 100
	maxX, maxY := 0, 0

	for tile := range t {
		if tile.x > maxX {
			maxX = tile.x
		} else if tile.x < minX {
			minX = tile.x
		}

		if tile.y > maxY {
			maxY = tile.y
		} else if tile.y < minY {
			minY = tile.y
		}
	}

	return minX, maxX, minY, maxY
}

func (t Tiles) FlippedNeighbours(x, y int) int {
	flipped := 0
	for _, coord := range dir2XY {
		if tile, ok := t[Tile{x + coord[0], y + coord[1]}]; ok {
			if tile == FLIPPED {
				flipped++
			}
		}
	}
	return flipped
}

func (t Tiles) Stats() (int, int) {
	flipped := 0
	init := 0

	for _, status := range t {
		if status == FLIPPED {
			flipped++
		} else {
			init++
		}
	}

	return init, flipped
}

func (t Tiles) NextCornwayIteration() {
	changes := make(map[Tile]int, 0)

	minX, maxX, minY, maxY := t.Boundaries()

	/* Increase the boundaries to let new tiles spawn */
	minY--
	minX--
	maxX++
	maxY++

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			n := t.FlippedNeighbours(x, y)
			currentTile := Tile{x, y}
			if status, ok := t[currentTile]; ok {
				if status == FLIPPED && (n == 0 || n > 2) {
					changes[currentTile] = INIT
				} else if status == INIT && n == 2 {
					changes[currentTile] = FLIPPED
				}
			} else if n == 2 {
				changes[currentTile] = FLIPPED
			}
		}
	}

	/* Apply all changes at once */
	for tile, status := range changes {
		t[tile] = status
	}
}

func main() {

	re := regexp.MustCompile(`(e|se|sw|w|nw|ne)`)

	tiles := make(Tiles, 0)
	for i, line := range util.LoadString(inputFile) {
		match := re.FindAllString(line, -1)
		if len(match) == 0 {
			fmt.Printf("Error parsing line %d: %s\n", i, line)
			return
		}

		x, y := 0, 0

		for j := 0; j < len(match); j++ {
			direction := match[j]
			coords := dir2XY[direction]
			x += coords[0]
			y += coords[1]
		}
		if tile, ok := tiles[Tile{x, y}]; ok {
			if tile == FLIPPED {
				tiles[Tile{x, y}] = INIT
			} else {
				tiles[Tile{x, y}] = FLIPPED
			}
		} else {
			tiles[Tile{x, y}] = FLIPPED
		}
	}

	_, flippedCount := tiles.Stats()
	fmt.Println(flippedCount)

	for i := 0; i < 100; i++ {
		tiles.NextCornwayIteration()
	}

	_, flippedCount = tiles.Stats()
	fmt.Println(flippedCount)

}
