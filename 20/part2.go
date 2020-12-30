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

func PrintPuzzle(p Puzzle, keys []int, width int) {
	if len(p.tiles) != len(keys) {
		fmt.Println(
			"Count of keys doesn't match count of tiles. Got",
			len(p.tiles),
			"tiles and",
			len(keys),
			"keys",
		)
		return
	}

	/* Render multiple tiles on the same line depending on the width
	 * parameter to make the puzzle observable
	 */

	/* Go through each key and overstep width-many tiles */
	for i := 0; i < len(keys); i += width {
		/* Go through the full height of a tile */
		for y := 0; y < p.tiles[keys[i]].image.Height(); y++ {
			/* Go through all tiles which should be on the same line */
			for j := i; j < i+width; j++ {
				currentImage := p.tiles[keys[j]].image
				fmt.Print(currentImage.data[y])
				fmt.Print(" ")
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
}
