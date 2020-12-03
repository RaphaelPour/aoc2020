package main

import (
	"fmt"

	"github.com/RaphaelPour/aoc2020/util"
)

type Slope struct {
	trees, index, offX, offY int
}

func (s Slope) String() string {
	return fmt.Sprintf("trees=%d, index=%d, offX=%d, offY=%d",
		s.trees, s.index, s.offX, s.offY,
	)
}

func main() {

	slopes := []Slope{
		Slope{offX: 1, offY: 1, index: 1},
		Slope{offX: 3, offY: 1, index: 1},
		Slope{offX: 5, offY: 1, index: 1},
		Slope{offX: 7, offY: 1, index: 1},
		Slope{offX: 1, offY: 2, index: 1},
	}

	for i, line := range util.LoadDefault() {
		/*if i == 0 {
			continue
		}*/
		for _, slope := range slopes {
			if i%slope.offY == 0 {
				if line[slope.index%len(line)] == '#' {
					slope.trees++
				}
				slope.index += slope.offX
			}
		}
	}

	result := 1
	for i, slope := range slopes {
		fmt.Printf("Slope %d:%s\n", i+1, slope)
		result *= slope.trees
	}

	fmt.Printf("%d\n", result)
}
