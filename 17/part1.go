package main

import (
	"fmt"

	"github.com/RaphaelPour/aoc2020/util"
)

const (
	ACTIVE   = '#'
	INACTIVE = '.'
)

/*
 * SLICE
 */

type Slice [][]rune

func NewSlice(length int) Slice {
	slice := make(Slice, length)

	row := make([]rune, length)
	for i := 0; i < length; i++ {
		row[i] = INACTIVE
	}

	/* Append inactive string `length` times */
	for i := 0; i < length; i++ {
		slice[i] = make([]rune, length)
		copy(slice[i], row)
	}

	return slice
}

func NewSliceFromLines(lines []string) (*Slice, error) {
	lineCount := len(lines)
	slice := make(Slice, lineCount)

	for i, line := range lines {
		if len(line) != lineCount {
			return nil, fmt.Errorf(
				"Line %d differs from line count: Expected %d, got %d",
				i, lineCount, len(line),
			)
		}
		slice[i] = []rune(line)
	}

	return &slice, nil
}

func (s Slice) String() string {
	out := ""
	for _, row := range s {
		out += string(row) + "\n"
	}
	return out
}

func (s Slice) Count() (int, int) {
	active := 0
	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[0]); x++ {
			if s[y][x] == '#' {
				active++
			}
		}
	}

	return active, len(s)*len(s) - active
}

func (s *Slice) Extend() {
	/*Extend each row by adding a leading/trailing INACTIVE cube */
	for i := range *s {
		(*s)[i] = append([]rune{INACTIVE}, (*s)[i]...)
		(*s)[i] = append((*s)[i], INACTIVE)
	}

	/* Adding leading/trailing row with only INACTIVE cubes */
	length := len(*s) + 2
	head := make([]rune, length)
	tail := make([]rune, length)
	for i := 0; i < length; i++ {
		head[i] = INACTIVE
		tail[i] = INACTIVE
	}
	/* Append inactive string `length` times */
	*s = append(*s, tail)
	*s = append(Slice{head}, *s...)
}

/*
 * GRID
 */

type Grid []Slice

func NewGrid(zLength, xyLength int) Grid {
	grid := make(Grid, zLength)
	for i := 0; i < zLength; i++ {
		grid[i] = NewSlice(xyLength)
	}
	return grid
}

func NewGridFromLines(lines []string) (*Grid, error) {
	grid := make(Grid, 1)
	sl, err := NewSliceFromLines(lines)
	if err != nil {
		return nil, err
	}
	grid[0] = *sl
	return &grid, err
}

func (g Grid) String() string {
	out := ""
	for z, slice := range g {
		out += fmt.Sprintf("z=%d\n%s\n\n", z, slice)
	}
	return out
}

func (g *Grid) Extend() {
	for i := range *g {
		(*g)[i].Extend()
	}

	/* Add leading/trailing slice with only INACTIVE cubes */
	*g = append(Grid{NewSlice(len((*g)[0][0]))}, *g...)
	*g = append(*g, NewSlice(len((*g)[0][0])))
}

func (g Grid) Neighbours(x, y, z int) int {
	// debugX, debugY, debugZ := 2, 0, 0
	neighbours := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				// Overstep current position
				if i == 0 && j == 0 && k == 0 {
					continue
				}
				zPos := z + i
				yPos := y + j
				xPos := x + k
				/* Check bounds */
				if zPos < 0 || zPos >= len(g) ||
					yPos < 0 || yPos >= len(g[zPos]) ||
					xPos < 0 || xPos >= len(g[zPos][yPos]) {
					continue
				}
				if g[zPos][yPos][xPos] == ACTIVE {
					neighbours++
				}

			}
		}
	}
	return neighbours
}

type Change struct {
	x, y, z int
	status  rune
}

func (g *Grid) NextCycle() {
	changes := make([]Change, 0)
	for z := range *g {
		for y := range (*g)[z] {
			for x := range (*g)[z][y] {
				neighbours := g.Neighbours(x, y, z)
				if (*g)[z][y][x] == ACTIVE && !util.InRange(neighbours, 2, 3) {
					changes = append(changes, Change{x, y, z, INACTIVE})
				} else if (*g)[z][y][x] == INACTIVE && neighbours == 3 {
					changes = append(changes, Change{x, y, z, ACTIVE})
				}
			}
		}
	}

	/* Apply all changes */
	for _, change := range changes {
		(*g)[change.z][change.y][change.x] = change.status
	}
}

func count(grid []Slice) (int, int) {
	active := 0
	inactive := 0
	for z := 0; z < len(grid); z++ {
		a, i := grid[z].Count()
		active += a
		inactive += i
	}
	return active, inactive
}

func main() {
	g, err := NewGridFromLines(util.LoadString("input"))
	if err != nil {
		fmt.Println(err)
		return
	}
	grid := *g

	fmt.Printf("Before any cycle:\n\n%s", grid)
	for i := 0; i < 6; i++ {
		grid.Extend()
		grid.NextCycle()
		fmt.Printf("After %d cycles:\n\n%s", i, grid)
	}

	active, _ := count(grid)
	fmt.Println(active)
}
