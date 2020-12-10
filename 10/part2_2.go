package main

import (
	"fmt"
	"math/big"
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

	outsiders := make([]int, 0)

	last := 0
	for _, num := range numbers {
		diff := util.Abs(last - num)
		if diff < 4 {
			outsiders = append(outsiders, 0)
		}
		last = num
	}

	fmt.Println("len(outsiders)=", len(outsiders))
	fmt.Println("Invalid:", "1125899906842623")

	combination := make([]int, 0)
	count := 0

	/* Length of outsider bits */
	inc := big.NewInt(1)
	length := big.NewInt(2).Exp(big.NewInt(2), big.NewInt(int64(len(outsiders))), nil)

	fmt.Println("Length:", length)

	i := uint(0)
	for bits := big.NewInt(2); bits.Cmp(length) < 0; bits = bits.Add(bits, inc) {

		if i%100000 == 0 {
			fmt.Print(".")
		}
		if i%1000000 == 0 {
			fmt.Println(count)
		}

		for i, num := range outsiders {
			if bits.And(bits.Rsh(bits, uint(i)), inc).Cmp(inc) == 0 {
				combination = append(combination, num)
			}
		}
		if checkCombination(combination, numbers) {
			count++
		}
		combination = combination[:0]
		i++
	}

	fmt.Println(count)
}

func checkCombination(comb, numbers []int) bool {
	last := 0
	for _, num := range numbers {
		if InInts(num, comb) {
			continue
		}
		diff := util.Abs(last - num)
		if diff >= 4 {
			return false
		}
		last = num
	}

	return true
}

func InInts(needle int, haystack []int) bool {
	for _, num := range haystack {
		if num == needle {
			return true
		}
	}
	return false
}

func powerSetSize(set []int) *big.Int {
	sets := big.NewInt(0)
	for size, _ := range set {
		var result big.Int
		sets = sets.Add(sets, result.Binomial(int64(len(set)), int64(size)))
	}
	return sets
}
