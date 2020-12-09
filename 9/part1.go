package main

import (
	"fmt"
	"strconv"

	"github.com/RaphaelPour/aoc2020/util"
)

func main() {

	// re := regexp.MustCompile(`^$`)
	numbers := make([]int, 0)
	for i, line := range util.LoadString("input") {
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Line %d is not a number: %s\n", i, line)
			return
		}
		numbers = append(numbers, num)
	}

	fmt.Printf("%d numbers found\n", len(numbers))

	preambleLength := 25
	preamble := numbers[:preambleLength]
	cipher := numbers[preambleLength:]

	fmt.Printf("len(Preamble)=%d\n", len(preamble))
	fmt.Printf("len(Cipher)=%d\n", len(cipher))

	for i := range cipher {
		if validate(cipher[i], preamble) {
			preamble = append(preamble, cipher[i])[1:]
		} else {
			fmt.Printf("%d is not valid\n", cipher[i])
			return
		}
	}

	fmt.Println("Invalid:", 24)
}

func validate(cipher int, preamble []int) bool {
	for i := 0; i < len(preamble); i++ {
		for j := i + 1; j < len(preamble); j++ {
			if preamble[i]+preamble[j] == cipher {
				fmt.Println("Found", preamble[i], "+", preamble[j], "=", cipher)
				return true
			}
		}
	}
	return false
}
