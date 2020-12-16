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
	indices    []int
	index      int
}

func (c Class) Within(x int) bool {
	return ((x >= c.min1) && (x <= c.max1)) ||
		((x >= c.min2) && (x <= c.max2))
}

func (c Class) String() string {
	return fmt.Sprintf(
		"| Index: %d\n| 1st range: %d-%d\n| 2nd range: %d-%d\n\\ Indices:%v",
		c.index, c.min1, c.max1, c.min2, c.max2, c.indices,
	)
}

/* Classes
 *
 * Map from class name to class.
 */
type Classes struct {
	hash map[string]*Class
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

func (c *Classes) Resolve(positions int) error {
	round := 1
	foundAny := false
	for !c.IsResolved() {

		if round%1000 == 0 {
			fmt.Printf("Round %d\n", round)
		}
		for i := 0; i < positions; i++ {
			// fmt.Println("Resolve position", i)
			count := 0
			target := ""
			for name, class := range c.hash {
				if class.index > -1 {
					// fmt.Println("Ignore", name)
					continue
				}
				if len(class.indices) == 0 {
					return fmt.Errorf("Class %s has no indices", name)
				}
				for _, j := range class.indices {
					if j == i {
						// fmt.Println("Found", i, "in class", name)
						count++
						target = name
					}
				}
			}
			if count == 1 && target != "" {
				foundAny = true
				fmt.Println("Define index", i, "in class", target)
				if class, ok := c.hash[target]; ok {
					class.index = i
					class.indices = []int{i}
				}
			}
		}
		if !foundAny {
			return fmt.Errorf("Couldn't resolve any index")
		}
		round++
	}

	return nil
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
		name:    name,
		min1:    min1,
		max1:    max1,
		min2:    min2,
		max2:    max2,
		indices: []int{},
		index:   -1,
	}, nil
}

func ParseIntList(line string) ([]int, error) {
	list := make([]int, 0)
	numbers := strings.Split(line, ",")
	for _, number := range numbers {
		num, err := strconv.Atoi(number)
		if err != nil {
			return nil, fmt.Errorf("Error parsing own ticket: %s", line)
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
	ownTicket := make([]int, 0)
	nearbyTickets := make([][]int, 0)
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
			if line == "your ticket:" {
				continue
			}

			list, err := ParseIntList(line)
			if err != nil {
				fmt.Println(err)
				return
			}
			ownTicket = list

		case STATE_NEARBY:
			if line == "nearby tickets:" {
				continue
			}

			ticket := make([]int, 0)
			numbers := strings.Split(line, ",")
			dropTicket := false
			classRelations := make(map[string][]int)
			for i, number := range numbers {
				num, err := strconv.Atoi(number)
				if err != nil {
					fmt.Println("Error parsing nearby ticket:", line)
					return
				}

				valid := false
				for name, class := range classes.hash {
					if class.Within(num) {
						if _, ok := classRelations[name]; !ok {
							classRelations[name] = make([]int, 0)
						}
						classRelations[name] = append(
							classRelations[name],
							i,
						)
						valid = true
					}
				}
				if !valid {
					errorRate += num
					dropTicket = true
					break
				}

				ticket = append(ticket, num)
			}

			if !dropTicket {
				for name, indices := range classRelations {
					if len(indices) == 1 {
						fmt.Println("Define", name, "to have index", indices[0])
						classes.hash[name].index = indices[0]
						classes.hash[name].indices = indices
					} else {
						classes.hash[name].indices =
							Union(
								indices,
								classes.hash[name].indices,
							)
					}
				}

				nearbyTickets = append(nearbyTickets, ticket)
			}
		}
	}

	/* Deduplicate classes */
	fmt.Println(classes)
	if err := classes.Resolve(len(ownTicket)); err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(classes)
	departureClasses := make(map[string]int, 0)
	for name, class := range classes.hash {
		if !strings.HasPrefix(name, "departure") {
			continue
		}
		departureClasses[name] = class.index
	}

	fmt.Println(departureClasses)

	result := 1
	for k, v := range departureClasses {
		fmt.Println(k, ":", v)
		result *= ownTicket[v]
	}

	fmt.Println("part1:", errorRate)
	fmt.Println("part2:", result)

	fmt.Println("Invalid", "412735589081")
}

func Union(source, target []int) []int {
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
