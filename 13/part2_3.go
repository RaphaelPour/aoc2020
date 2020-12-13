package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"strconv"
	"strings"
)

func main() {

	lines := util.LoadString("input2")
	ts := make([]uint64, 0)
	i := 0
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

		fmt.Println(num, i, num-i)
		ts = append(ts, uint64(num-i))

		i++
	}

	result := LCM(ts)
	fmt.Println("\r>>", result, "<<")
	return

}

func LCM(num []uint64) uint64 {
	a, b := num[0], num[1]
	result := a * b / GCD(a, b)

	for i := 0; i < len(num[2:]); i++ {
		result = LCM([]uint64{result, num[i+2]})
	}

	return result
}

func GCD(a, b uint64) uint64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}
