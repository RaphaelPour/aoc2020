package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/RaphaelPour/aoc2020/util"
)

func main() {

	// re := regexp.MustCompile(`^$`)
	numbers := make([]int, 0)
	for i, line := range util.LoadString("input") {
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("line %d is not a numbers. %s\n", i, line)
			return
		}
		numbers = append(numbers, num)
	}

	sort.Ints(numbers)

	numbers = append(numbers, numbers[len(numbers)-1]+3)

	histogram := make(map[int]int, 0)
	last := 0
	for _, num := range numbers {
		diff := util.Abs(last - num)
		fmt.Printf("| %d - %d | = %d\n", last, num, diff)
		if _, ok := histogram[diff]; !ok {
			histogram[diff] = 1
		} else {
			histogram[diff]++
		}

		last = num
	}
	fmt.Println(histogram)
	fmt.Println(histogram[1] * histogram[3])
}
