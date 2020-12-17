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

func (g Grid) Count() (int, int) {
	active := 0
	inactive := 0
	for _, slice := range g {
		a, i := slice.Count()
		active += a
		inactive += i
	}
	return active, inactive
}

type Cluster []Grid

func NewClusterFromLines(lines []string) (*Cluster, error) {
	cluster := make(Cluster, 1)
	g, err := NewGridFromLines(lines)
	if err != nil {
		return nil, err
	}
	cluster[0] = *g
	return &cluster, err
}

func (c Cluster) String() string {
	out := ""
	for w, grid := range c {
		out += fmt.Sprintf("w=%d\n%s", w, grid)
	}
	return out
}

func (c Cluster) Count() (int, int) {
	active := 0
	inactive := 0
	for _, grid := range c {
		a, i := grid.Count()
		active += a
		inactive += i
	}
	return active, inactive
}
func (c *Cluster) Extend() {
	for i := range *c {
		(*c)[i].Extend()
	}

	/* Add leading/trailing slice with only INACTIVE cubes */
	gridLength := len((*c)[0])
	sliceLength := len((*c)[0][0])
	*c = append(Cluster{NewGrid(gridLength, sliceLength)}, *c...)
	*c = append(*c, NewGrid(gridLength, sliceLength))
}

func (c Cluster) Neighbours(x, y, z, w int) int {
	neighbours := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					// Overstep current position
					if i == 0 && j == 0 && k == 0 && l == 0 {
						continue
					}
					wPos := w + i
					zPos := z + j
					yPos := y + k
					xPos := x + l
					/* Check bounds */
					if wPos < 0 || wPos >= len(c) ||
						zPos < 0 || zPos >= len(c[wPos]) ||
						yPos < 0 || yPos >= len(c[wPos][zPos]) ||
						xPos < 0 || xPos >= len(c[wPos][zPos][yPos]) {
						continue
					}
					if c[wPos][zPos][yPos][xPos] == ACTIVE {
						neighbours++
					}

				}
			}
		}
	}
	return neighbours

}

type Change struct {
	x, y, z, w int
	status     rune
}

func (c *Cluster) NextCycle() {
	changes := make([]Change, 0)
	for w := range *c {
		for z := range (*c)[w] {
			for y := range (*c)[w][z] {
				for x := range (*c)[w][z][y] {
					neighbours := c.Neighbours(x, y, z, w)
					if (*c)[w][z][y][x] == ACTIVE && !util.InRange(neighbours, 2, 3) {
						changes = append(changes, Change{x, y, z, w, INACTIVE})
					} else if (*c)[w][z][y][x] == INACTIVE && neighbours == 3 {
						changes = append(changes, Change{x, y, z, w, ACTIVE})
					}
				}
			}
		}
	}

	/* Apply all changes */
	for _, change := range changes {
		(*c)[change.w][change.z][change.y][change.x] = change.status
	}
}

func main() {
	c, err := NewClusterFromLines(util.LoadString("input"))
	if err != nil {
		fmt.Println(err)
		return
	}
	cluster := *c

	fmt.Printf("Before any cycle:\n\n%s", cluster)
	for i := 0; i < 6; i++ {
		cluster.Extend()
		cluster.NextCycle()
		fmt.Printf("After %d cycles:\n\n%s", i, cluster)
	}

	active, _ := cluster.Count()
	fmt.Println(active)
}
