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
	staticColor = color.New(color.FgCyan, color.Bold)
	inputFile   = "input"
)

func main() {

	/* Parse input. Each line is the 'joltage' (int) of an adapter */
	numbers := []int{0}
	for i, line := range util.LoadString(inputFile) {
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("line %d is not a numbers. %s\n", i, line)
			return
		}
		numbers = append(numbers, num)
	}

	/* Based from sorted adapters, the problem can be solved by only
	 * using or not using an adapter without the need of rearranging the
	 * adapter any further.
	 */
	sort.Ints(numbers)

	fmt.Println("[S]plit [R]eturn [C]ombinate")
	fmt.Println(">>>", findBinary(numbers, 0), "<<<")
}

/* findBinary:
 *
 * Some adapters can't be removed since they are needed to keep
 * the difference of the surrounding adapters lower than four.
 *
 * This function splits the adapters at those mandatory elements
 * recursively with a new list of numbers having the previous and
 * next mandatory as 'boundary' until:
 * - only two are left: No combination possible, only mandatory left
 * - No further mandatory left (except the first and last) -> calculate
 *   combinations. Decrement result if the difference between first and
 *   last is higher than 3.
 */
func findBinary(numbers []int, depth int) int {
	if len(numbers) < 3 {
		fmt.Printf("R:   1")
		printPartition(-1, depth, numbers)
		return 1
	}

	for i := 1; i < len(numbers)-1; i++ {
		if util.Abs(numbers[i-1]-numbers[i+1]) > 3 {
			fmt.Printf("S: %3d", numbers[i])
			printPartition(i, depth, numbers)
			return findBinary(numbers[:i+1], depth+1) *
				findBinary(numbers[i:], depth+1)
		}
	}

	/* Calculate all possible combinations of the numbers between the
	 * first+last (mandatory) adapters: 2^|volatile adapters|
	 * with: |volatile adapters| = |adapters|-2
	 */
	result := int(math.Pow(2, float64(len(numbers)-2)))

	/* Remove one combination if all volatile numbers can't miss all
	 * at once when the difference would be too high.
	 */
	if numbers[len(numbers)-1]-numbers[0] > 3 {
		result--
	}
	fmt.Printf("C: %3d", result)
	printPartition(-1, depth, numbers)
	return result
}

func printPartition(index, depth int, numbers []int) {

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
