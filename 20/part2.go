package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"regexp"
	"strconv"
)

var (
	inputFile = "input"
)

func main() {

	re := regexp.MustCompile(`^Tile (\d+):$`)
	puzzle := NewPuzzle()
	img := NewImage()
	id := 0
	for i, line := range util.LoadString(inputFile) {
		if line == "" {
			continue
		}
		match := re.FindStringSubmatch(line)
		if len(match) == 2 {
			/* Store last buffered image if it's not the first empty one */
			if id != 0 {
				t := NewTileFromImage(&img)
				t.id = id
				puzzle.tiles[id] = &t
				img = NewImage()
			}
			num, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Printf(
					"Tile ID '%s' is not a number on index %d\n",
					match[1],
					i,
				)
				return
			}
			id = num
		} else {
			if err := img.AddRow(line); err != nil {
				fmt.Println(err)
				return
			}
		}

	}

	/* Add last image */
	t := NewTileFromImage(&img)
	t.id = id
	puzzle.tiles[id] = &t

	/* Print puzzle */
	/* PrintPuzzle(tiles, tiles.Keys(), 3) */
	if err := puzzle.Arrange(); err != nil {
		fmt.Printf("Error arranging puzzle: %s\n", err)
		return
	}
}
