package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/RaphaelPour/aoc2020/util"
	"github.com/fatih/color"
)

var (
	goodColor   = color.New(color.FgGreen, color.Bold)
	badColor    = color.New(color.FgRed, color.Bold)
	staticColor = color.New(color.FgCyan, color.Bold)
	inputFile   = "input"
	solutions   = map[string]int{
		"input2": 8,
		"input3": 19208,
	}
)

func main() {

	/* Parse input where each line represents an int defining
	 * the 'joltage' of an adapter
	 */
	numbers := []int{0}
	for i, line := range util.LoadString(inputFile) {
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("line %d is not a numbers. %s\n", i, line)
			return
		}
		numbers = append(numbers, num)
	}

	/* Since they have to be arranged in such a way, that the differences
	 * between two adapters don't overcome 3 they can be sorted. THen it is
	 * easier to find those adapters which must be at a certain position
	 * regardless of the combination. Otherwise the difference between two
	 * adapters would get to high.
	 */
	sort.Ints(numbers)

	fmt.Println("[S]plit [R]eturn [C]ombinate")

	result := findBinary(numbers, 0)
	fmt.Println("Invalid:", 34012224)
	fmt.Println(">>>", result, "<<<")

	solution, ok := solutions[inputFile]
	if ok {
		if solution == result {
			goodColor.Println("OK")
		} else {
			badColor.Println("NOT OK")
		}
	}
}

func findBinary(numbers []int, depth int) int {
	if len(numbers) < 3 {
		fmt.Printf("R:   1")
		renderPartition(-1, depth, numbers)
		return 1
	}

	for i := 1; i < len(numbers)-1; i++ {
		if util.Abs(numbers[i-1]-numbers[i+1]) > 3 {
			fmt.Printf("S: %3d", numbers[i])
			renderPartition(i, depth, numbers)
			return findBinary(numbers[:i+1], depth+1) *
				findBinary(numbers[i:], depth+1)
		}
	}
	result := int(math.Pow(2, float64(len(numbers)-2)))

	/* Remove the combination where all volatile numbers
	 * can't miss all at once since the difference would be too high.
	 */
	if numbers[len(numbers)-1]-numbers[0] > 3 {
		result--
	}
	fmt.Printf("C: %3d", result)
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
			fmt.Printf("] ")
			staticColor.Printf("%d", numbers[k])
			fmt.Printf(" [")
		} else {
			fmt.Printf("%d ", numbers[k])
		}
	}
	fmt.Println("]")
}
