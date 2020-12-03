package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/RaphaelPour/aoc2020/util"
)

func main() {

	re := regexp.MustCompile(`^(\d+)-(\d+)\s(\w):\s(\w+)$`)

	correct1 := 0
	correct2 := 0

	/* Parse input to struct */
	for i, line := range util.LoadDefault() {
		match := re.FindStringSubmatch(line)

		min, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Printf("Min '%s' is not a number in line %d\n", match[1], i)
			return
		}

		max, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Printf("Max '%s' is not a number in line %d\n", match[2], i)
			return
		}

		char := match[3]
		pw := match[4]

		count1 := strings.Count(pw, char)

		if count1 <= max && count1 >= min {
			correct1++
		}

		count2 := 0
		if min-1 >= 0 && pw[min-1] == char[0] {
			count2 += 1
		}

		if max-1 >= 0 && pw[max-1] == char[0] {
			count2 += 1
		}

		if count2 == 1 {
			correct2++
		}
	}
	fmt.Printf("Valid part 1: %d\n", correct1)
	fmt.Printf("Valid part 2: %d\n", correct2)
}
