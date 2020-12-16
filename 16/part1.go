package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"regexp"
	"strconv"
	"strings"
)

const (
	STATE_CLASS = iota
	STATE_TICKET
	STATE_NEARBY
	STATE_INVALID
)

var (
	reFields  = regexp.MustCompile(`^([a-z\s]+): (\d+)-(\d+) or (\d+)-(\d+)$`)
	inputFile = "input"
)

/* Class:
 *
 * Field on the ticket including validation parameters, name and possible
 * indices.
 */
type Class struct {
	name       string
	min1, max1 int
	min2, max2 int
}

func (c Class) Within(n int) bool {
	return ((n >= c.min1) && (n <= c.max1)) ||
		((n >= c.min2) && (n <= c.max2))
}

/* Classes
 *
 * Map from class name to class.
 */
type Classes struct {
	hash map[string]*Class
}

func ParseField(line string) (string, *Class, error) {

	match := reFields.FindStringSubmatch(line)
	if len(match) < 6 {
		return "", nil, fmt.Errorf("error on parsing class: %s", line)
	}

	name := match[1]

	min1, err := strconv.Atoi(match[2])
	if err != nil {
		return name, nil, fmt.Errorf("Error parsing first number: %s", line)
	}
	max1, err := strconv.Atoi(match[3])
	if err != nil {
		return name, nil, fmt.Errorf("Error parsing second number: %s", line)
	}

	min2, err := strconv.Atoi(match[4])
	if err != nil {
		return name, nil, fmt.Errorf("Error parsing third number: %s", line)
	}
	max2, err := strconv.Atoi(match[5])
	if err != nil {
		return name, nil, fmt.Errorf("Error parsing fourth number: %s", line)
	}

	return name, &Class{
		name: name,
		min1: min1,
		max1: max1,
		min2: min2,
		max2: max2,
	}, nil
}

func ParseIntList(line string) ([]int, error) {
	list := make([]int, 0)
	numbers := strings.Split(line, ",")
	for _, number := range numbers {
		num, err := strconv.Atoi(number)
		if err != nil {
			return nil, fmt.Errorf("Error parsing ticket: %s", line)
		}
		list = append(list, num)
	}
	return list, nil
}

func main() {

	state := STATE_CLASS

	classes := Classes{
		hash: map[string]*Class{},
	}
	errorRate := 0

	for _, line := range util.LoadString(inputFile) {

		if line == "" {
			state++
			if state == STATE_INVALID {
				fmt.Println("Reached invalid state ", state)
				return
			}
			continue
		}

		switch state {
		case STATE_CLASS:
			name, class, err := ParseField(line)
			if err != nil {
				fmt.Println(name, ":", err)
				return
			}
			classes.hash[name] = class
		case STATE_TICKET:
			/* Ignore own ticket in the first part */
		case STATE_NEARBY:
			if line == "nearby tickets:" {
				continue
			}

			fields, err := ParseIntList(line)
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, field := range fields {
				valid := false
				for _, class := range classes.hash {
					if class.Within(field) {
						valid = true
						break
					}
				}
				if !valid {
					errorRate += field
					break
				}

			}
		}
	}

	fmt.Println(errorRate)
}
