package main

import (
	"fmt"

	"github.com/RaphaelPour/aoc2020/util"
)

func main() {

	result := 0
	group := make(map[rune]bool, 0)
	for _, line := range util.LoadString("input") {
		if line == "" {
			result += len(group)
			group = make(map[rune]bool, 0)
		} else {
			for _, char := range line {
				group[char] = true
			}
		}
	}

	fmt.Println(result)
}
