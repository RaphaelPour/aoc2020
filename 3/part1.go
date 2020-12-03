package main

import (
	"fmt"

	"github.com/RaphaelPour/aoc2020/util"
)

func main() {

	trees := 0
	for i, line := range util.LoadDefault() {
		if line[i*3%len(line)] == '#' {
			trees++
		}
	}

	fmt.Println(trees)
}
