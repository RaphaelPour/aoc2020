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

	badNumber := 0
	for i := range cipher {
		if validate(cipher[i], preamble) {
			preamble = append(preamble, cipher[i])[1:]
		} else {
			badNumber = cipher[i]
			break
		}
	}

	fmt.Printf("%d is not valid\n", badNumber)

	crack(badNumber, numbers)
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

func crack(badNumber int, numbers []int) {

	summands := make([]int, 0)
	for size := 2; size < len(numbers); size++ {
		for i := 0; i < len(numbers); i++ {
			sum := 0
			summands = summands[:0]
			for j := i + 1; j < size; j++ {
				sum += numbers[j]
				summands = append(summands, numbers[j])
			}

			if sum == badNumber {
				min, max := util.MinMaxInts(summands)
				fmt.Printf("min=%d, max=%d, sum=%d\n", min, max, min+max)
				return
			}
		}
	}
}
