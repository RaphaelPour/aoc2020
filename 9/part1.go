package main

import (
	"fmt"
	"strconv"

	"github.com/RaphaelPour/aoc2020/util"
)

func main() {

	numbers := make([]int, 0)
	for i, line := range util.LoadString("input") {
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Line %d is not a number: %s\n", i, line)
			return
		}
		numbers = append(numbers, num)
	}

	preambleLength := 25
	preamble := numbers[:preambleLength]
	cipher := numbers[preambleLength:]

	for i := range cipher {
		if validate(cipher[i], preamble) {
			preamble = append(preamble, cipher[i])[1:]
		} else {
			fmt.Println(cipher[i])
			return
		}
	}
}

func validate(cipher int, preamble []int) bool {
	for i := 0; i < len(preamble); i++ {
		for j := i + 1; j < len(preamble); j++ {
			if preamble[i]+preamble[j] == cipher {
				return true
			}
		}
	}
	return false
}
