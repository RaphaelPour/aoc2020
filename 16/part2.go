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
	name            string
	min1, max1      int
	min2, max2      int
	possibleIndices []bool
	index           int
}

func (c Class) Within(n int) bool {
	return ((n >= c.min1) && (n <= c.max1)) ||
		((n >= c.min2) && (n <= c.max2))
}

func (c Class) ValidIndices() []int {
	result := make([]int, 0)
	for i, val := range c.possibleIndices {
		if val {
			result = append(result, i)
		}
	}
	return result
}

func (c Class) String() string {
	trueCount := len(c.ValidIndices())

	return fmt.Sprintf(
		"| Index: %d\n| 1st range: %d-%d\n| 2nd range: %d-%d\n\\ Indices:%v (%d)",
		c.index, c.min1, c.max1, c.min2, c.max2, c.possibleIndices, trueCount,
	)
}

func (c *Class) Resolve() bool {
	/* Check if there is only one possible index and
	 * set the index to this one
	 */
	if c.index != -1 {
		fmt.Println("index is > -1", c.index)
		return true
	}
	indices := c.ValidIndices()
	if len(indices) == 1 {
		c.index = indices[0]
		return true
	}
	return false
}

/* Classes
 *
 * Map from class name to class.
 */
type Classes struct {
	hash map[string]*Class
}

func (c *Classes) InitIndices(fieldCount int) {
	defaultIndices := make([]bool, fieldCount)
	for i, _ := range defaultIndices {
		defaultIndices[i] = true
	}

	for _, class := range c.hash {
		class.possibleIndices = make([]bool, fieldCount)
		copy(class.possibleIndices, defaultIndices)
	}
}

func (c Classes) FindAllInverse(n int) []string {
	result := make([]string, 0)
	for name, class := range c.hash {
		if !class.Within(n) {
			result = append(result, name)
		}
	}
	return result
}

func (c Classes) FindAll(n int) []string {
	result := make([]string, 0)
	for name, class := range c.hash {
		if class.Within(n) {
			result = append(result, name)
		}
	}

	return result
}

func (c Classes) WithinAny(n int) bool {
	for _, class := range c.hash {
		if class.Within(n) {
			return true
		}
	}

	return false
}

func (c Classes) String() string {
	out := ""
	for name, class := range c.hash {
		out += fmt.Sprintf("%s:\n%s\n\n", name, class)
	}
	return out
}

func (c Classes) IsResolved() bool {
	for _, class := range c.hash {
		if class.index == -1 {
			return false
		}
	}
	return true
}

func (c *Classes) Resolve(fieldCount int) {
	round := 0
	for !c.IsResolved() {
		for name, class := range c.hash {
			/* Skip all already resolved classes */
			if class.index != -1 {
				continue
			}

			/*
			 * Check if the class has already only one possible index
			 * and set this index at all other classes to false.
			 */
			if class.Resolve() {
				for nameOther, classOther := range c.hash {
					/* Skip resolved class */
					if name == nameOther {
						continue
					}

					classOther.possibleIndices[class.index] = false
				}
			}
		}
		round++
	}

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
		name:            name,
		min1:            min1,
		max1:            max1,
		min2:            min2,
		max2:            max2,
		possibleIndices: []bool{},
		index:           -1,
	}, nil
}

func ParseTicket(line string) ([]int, error) {
	ticket := make([]int, 0)
	for i, rawField := range strings.Split(line, ",") {
		field, err := strconv.Atoi(rawField)
		if err != nil {
			return nil, fmt.Errorf(
				"Error parsing ticket in line %d: %s",
				i,
				line,
			)
		}
		ticket = append(ticket, field)
	}
	return ticket, nil
}

func main() {

	state := STATE_CLASS

	classes := Classes{
		hash: map[string]*Class{},
	}
	ownTicket := make([]int, 0)
	nearbyTickets := make([][]int, 0)
	ticketLength := 0

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
			if line == "your ticket:" {
				continue
			}

			ticket, err := ParseTicket(line)
			if err != nil {
				fmt.Println(err)
				return
			}
			ticketLength = len(ticket)
			ownTicket = ticket

		case STATE_NEARBY:
			if line == "nearby tickets:" {
				continue
			}

			ticket, err := ParseTicket(line)
			if err != nil {
				fmt.Println(err)
				return
			}

			/* Drop all invalid tickets */
			validTicket := true
			for _, field := range ticket {
				if !classes.WithinAny(field) {
					validTicket = false
					break
				}
			}
			if validTicket {
				nearbyTickets = append(nearbyTickets, ticket)
			}

		}
	}

	/* Initialize the possible indices of all classes with true. They will
	 * get false sequentially
	 */
	classes.InitIndices(ticketLength)

	for _, ticket := range nearbyTickets {
		/* Find which classes can't be at this index */
		for i, field := range ticket {
			for _, name := range classes.FindAllInverse(field) {
				classes.hash[name].possibleIndices[i] = false
			}
		}
	}

	/* Resolve indice ambiguity */
	classes.Resolve(ticketLength)

	departureClasses := make(map[string]int, 0)
	for name, class := range classes.hash {
		if !strings.HasPrefix(name, "departure") {
			continue
		}
		departureClasses[name] = class.index
	}

	result := 1
	for _, v := range departureClasses {
		result *= ownTicket[v]
	}

	fmt.Println(result)
}

func Union(target, source []int) []int {
	for _, src := range source {
		found := false
		for _, tgt := range target {
			if tgt == src {
				found = true
				break
			}
		}
		if !found {
			target = append(target, src)
		}
	}

	return target
}

func intersectAll(sets [][]int) int {
	if len(sets) == 0 {
		fmt.Println("No set given")
		return -1
	}

	result := sets[0]
	for i, set := range sets {
		if i == 0 {
			continue
		}
		result = intersect(result, set)
	}

	if len(result) != 1 {
		fmt.Println("Result set has unexpected length", len(result))
		fmt.Println(sets)
		return -1
	}

	return result[0]
}

func intersect(a, b []int) []int {
	results := make([]int, 0)

	for _, valueA := range a {
		for _, valueB := range b {
			if valueA == valueB {
				results = append(results, valueA)
			}
		}
	}
	return results
}
