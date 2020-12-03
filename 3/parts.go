package main

import (
	"fmt"

	"github.com/RaphaelPour/aoc2020/util"
)

func main() {

	/* Parse input to struct */
	trees := []int{0, 0, 0, 0, 0}
	indices := []int{1, 3, 5, 7, 1}
	for i, line := range util.Load("input") {
		if i == 0 {
			continue
		}
		if line[indices[0]%len(line)] == '#' {
			trees[0]++
		}

		if line[indices[1]%len(line)] == '#' {
			trees[1]++
		}

		if line[indices[2]%len(line)] == '#' {
			trees[2]++
		}

		if line[(indices[3])%len(line)] == '#' {
			trees[3]++
		}

		if i%2 == 0 {
			if line[indices[4]%len(line)] == '#' {
				trees[4]++
			}
			indices[4]++
		}

		indices[0] += 1
		indices[1] += 3
		indices[2] += 5
		indices[3] += 7
	}
	fmt.Println("Invalid: 3049503392")
	fmt.Println("Invalid: 3716582259")
	fmt.Println("Invalid: 692311536")
	fmt.Println("Invalid: 614836656")
	fmt.Println("Invalid: 596157156")
	fmt.Println("Invalid: 634618908")
	fmt.Println("Invalid: 3907176221")

	if trees[1] != 299 {
		fmt.Printf("Slope 2 (right 3, down 1) is wrong. Got %d, expected 299\n", trees[1])
	}

	result := 1
	for i, c := range trees {
		result *= c
	}
	fmt.Printf("Part1:  %d\n", trees[1])
	fmt.Printf("Part2:  %d\n", result)
}
