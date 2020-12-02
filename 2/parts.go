package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Entry struct {
	min, max int
	needle   string
	password string
}

func (e Entry) String() string {
	return fmt.Sprintf("min=%d, max=%d, needle=%s, password=%s",
		e.min, e.max, e.needle, e.password,
	)
}

func main() {

	re := regexp.MustCompile(`^(\d+)-(\d+)\s(\w):\s(\w+)$`)
	scanner := bufio.NewScanner(os.Stdin)

	index := 0
	correct1 := 0
	correct2 := 0
	/* Parse input to struct */
	for scanner.Scan() {
		line := scanner.Text()

		match := re.FindStringSubmatch(line)

		min, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Printf("Min '%s' is not a number in line %d\n", match[1], index)
			return
		}

		max, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Printf("Max '%s' is not a number in line %d\n", match[1], index)
			return
		}

		e := Entry{
			min:      min,
			max:      max,
			needle:   match[3],
			password: match[4],
		}

		count1 := strings.Count(match[4], match[3])

		if count1 <= max && count1 >= min {
			correct1++
		}
		count := 0
		if min-1 >= 0 && e.password[min-1] == e.needle[0] {
			count += 1
		}

		if max-1 >= 0 && e.password[max-1] == e.needle[0] {
			count += 1
		}

		if count == 1 {
			correct2++
		}

		index++
	}
	fmt.Printf("Valid part 1: %d\n", correct1)
	fmt.Printf("Valid part 2: %d\n", correct2)
}
