package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/RaphaelPour/aoc2020/util"
)

func main() {

	// re := regexp.MustCompile(`^$`)
	numbers := []int{0}
	for i, line := range util.LoadString("input2") {
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("line %d is not a numbers. %s\n", i, line)
			return
		}
		numbers = append(numbers, num)
	}

	sort.Ints(numbers)

	numbers = append(numbers, numbers[len(numbers)-1]+3)
	fmt.Println("Invalid:", 34012224)
	fmt.Println(findBinary(numbers, 0))
}

func findBinary(numbers []int, depth int) int {
	if len(numbers) <= 1 { //|| util.Abs(numbers[0]-numbers[len(numbers)-1]) > 3 {
		renderPartition(-1, depth, numbers)
		return 1
	}

	for i := 1; i < len(numbers)-1; i++ {
		if util.Abs(numbers[i-1]-numbers[i+1]) > 3 {
			renderPartition(i, depth, numbers)
			return findBinary(numbers[:i-1], depth+1) * findBinary(numbers[i:], depth+1)
		}
	}
	result := int(math.Pow(2, float64(len(numbers))))
	renderPartition(-1, depth, numbers)
	return result
}

func renderPartition(index, depth int, numbers []int) {

	for i := 0; i < depth; i++ {
		fmt.Print(" ")
	}

	fmt.Print("[")
	for k := range numbers {
		if k == index {
			fmt.Printf("] %d [", numbers[k])
		} else {
			fmt.Printf("%d ", numbers[k])
		}
	}
	fmt.Println("]")
}
