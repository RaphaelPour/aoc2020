package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"strconv"
	"strings"
)

type Entry struct {
	uint64
}

func main() {

	lines := util.LoadString("input")
	ts := make([]uint64, 0)
	pos := make([]uint64, 0)
	i := 0
	max := 0
	for _, cell := range strings.Split(lines[1], ",") {

		if cell == "x" {
			i++
			continue
		}

		num, err := strconv.Atoi(cell)
		if err != nil {
			fmt.Println("Cell is not a number", cell)
			return
		}

		fmt.Println(num, "-", i, "=", num-i)
		if num > max {
			max = num
		}
		ts = append(ts, uint64(num))
		pos = append(pos, uint64(i))
		i++
	}

	fmt.Println("Timestamps:", ts)
	fmt.Println("Positions: ", pos)
	fmt.Println("Max:", max)

	return

}
