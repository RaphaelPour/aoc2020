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

	hardcoreOutsiders := make([]int, 0)
	ultimateOutsiders := make([]int, 0)
	sharewareOutsiders := make([]int, 0)
	last := 0
	postLast := 0
	postPostLast := 0
	next := 0
	for i, num := range numbers {
		if i+1 == len(numbers) {
			break
		}
		next = numbers[i+1]
		diff1 := util.Abs(last - next)
		diff2 := util.Abs(postLast - next)
		diff3 := util.Abs(postPostLast - next)
		if diff1 < 4 && diff2 < 4 && diff3 < 4 {
			hardcoreOutsiders = append(hardcoreOutsiders, num)
		} else if diff1 < 4 && diff2 < 4 {
			ultimateOutsiders = append(ultimateOutsiders, num)
		} else if diff1 < 4 && diff2 >= 4 {
			sharewareOutsiders = append(sharewareOutsiders, num)
		}
		last = num
		if i > 0 {
			postLast = numbers[i-1]
		}
		if i > 1 {
			postPostLast = numbers[i-2]
		}
	}

	fmt.Println("len(ultimateOutsiders)=", len(ultimateOutsiders))
	fmt.Println("len(sharewareOutsiders)=", len(sharewareOutsiders))
	fmt.Println("len(hardcoreOutsiders)=", len(hardcoreOutsiders))

	fmt.Println("Invalid:", "1125899906842623", "too low: 537395200")
	fmt.Println("|P(ultimateOutsiders)| =", powerSetSize(ultimateOutsiders))
	fmt.Println("|P(shrewareOutsiders)| =", powerSetSize(sharewareOutsiders))
	fmt.Println("|P(hardcoreOutsiders)| =", powerSetSize(hardcoreOutsiders))

}

func powerSetSize(set []int) *big.Int {
	sets := big.NewInt(0)
	for size, _ := range set {
		var result big.Int
		sets = sets.Add(sets, result.Binomial(int64(len(set)), int64(size)))
	}
	return sets
}
