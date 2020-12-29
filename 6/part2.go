package main

import (
	"fmt"

	"github.com/RaphaelPour/aoc2020/util"
)

func main() {

	result := 0
	group := make(map[rune]int, 0)
	persons := 0
	for _, line := range util.LoadString("input") {
		if line == "" {
			for _, count := range group {
				if count == persons {
					result++
				}
			}
			persons = 0
			group = make(map[rune]int, 0)
		} else {
			for _, char := range line {
				if _, ok := group[char]; !ok {
					group[char] = 1
				} else {
					group[char]++
				}
			}
			persons++
		}
	}

	fmt.Println(result)
}
