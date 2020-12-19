package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var (
	inputFile = "input4"
	cache     = Cache{}
)

type Product struct {
	terminal bool
	rules    []int
	symbol   string
}

type Products []Product

func (p Products) MaxLength() int {
	max := 0
	for _, prod := range p {
		if len(prod.rules) > max {
			max = len(prod.rules)
		}
	}
	return max
}

func (p Products) MinLength() int {
	min := 0
	for _, prod := range p {
		if len(prod.rules) < min {
			min = len(prod.rules)
		}
	}
	return min
}

type Rules map[int]Products

type Cache struct {
	rules map[int]CacheRule
	debug bool
}

func NewCache() Cache {
	c := Cache{}
	c.rules = make(map[int]CacheRule, 0)
	return c
}

type CacheRule struct {
	badInputs           [][]int
	maxProductionLength int
}

func (c Cache) Print(action, msg string, ruleID int, input []int) {
	if !c.debug {
		return
	}
	printer := color.New(color.FgWhite)
	switch action {
	case "MISS":
		printer.Add(color.FgRed)
	case "HIT":
		printer.Add(color.FgGreen)
	case "IGN", "ADD", "DUPL":
		printer.Add(color.FgYellow)
	}
	printer.Printf("[%6s] %s %d %#v\n", action, msg, ruleID, input)
}

func (c Cache) IsKnownBad(ruleID int, input []int) bool {
	cacheRule, ok := c.rules[ruleID]

	if !ok {
		c.Print("MISS", "unknown rule", ruleID, input)
		return false
	}

	for _, badInput := range cacheRule.badInputs {
		if len(badInput) > len(input) {
			continue
		}
		hit := true
		for i, id := range badInput {
			if input[i] != id {
				hit = false
				break
			}
		}
		if hit {
			c.Print("HIT", "", ruleID, input)
			return true
		}
	}

	c.Print("MISS", "unknown input", ruleID, input)
	return false
}

func (c *Cache) AddBad(ruleID int, input []int, rule Products) {
	if c.IsKnownBad(ruleID, input) {
		c.Print("DUPL", "", ruleID, input)
		return
	}

	max := rule.MaxLength()

	if max > len(input) {
		c.Print("IGN", "input is too short", ruleID, input)
		return
	}

	if _, ok := c.rules[ruleID]; !ok {
		cr := CacheRule{}
		cr.maxProductionLength = max
		cr.badInputs = make([][]int, 1)
		cr.badInputs[0] = make([]int, max)
		copy(cr.badInputs[0], input[:max])
		c.rules[ruleID] = cr
		return
	} else {
		cr := c.rules[ruleID]
		newInput := make([]int, cr.maxProductionLength)
		copy(newInput, input[:cr.maxProductionLength])

		cr.badInputs = append(cr.badInputs, newInput)
		c.rules[ruleID] = cr
	}
	c.Print("ADD", "", ruleID, input)

}

func main() {

	re := regexp.MustCompile(`^(\d+): (.*)$`)
	terminalRe := regexp.MustCompile(`"([a-z])"`)
	rules := make(Rules, 0)
	parseRules := true
	count := 0
	for i, line := range util.LoadString(inputFile) {
		if parseRules {
			if line == "" {
				parseRules = false
				continue
			}

			/* HOT PATCH */
			if strings.HasPrefix(line, "8:") {
				line = "8: 42 | 42 8"
			} else if strings.HasPrefix(line, "11:") {
				line = "11: 42 31 | 42 11 31"
			}

			match := re.FindStringSubmatch(line)
			if len(match) < 3 {
				fmt.Println("Error parsing line", i, ":", line)
				return
			}

			ruleID, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Printf(
					"Error converting index in line  %d to number: %s\n",
					i, match[1],
				)
				return
			}

			products := make(Products, 0)

			for j, prod := range strings.Split(match[2], " | ") {
				if prod == "" {
					fmt.Println("error parsing rule", j, ". Alternative can't be empty:", match[2])
					return
				}

				product := Product{}
				/* Check if terminal or rules are present */
				if match := terminalRe.FindStringSubmatch(prod); len(match) > 0 {
					product.terminal = true
					product.symbol = string(match[1])
					products = append(products, product)
					continue
				}

				product.rules = make([]int, 0)

				for k, rule := range strings.Split(prod, " ") {
					num, err := strconv.Atoi(rule)
					if err != nil {
						fmt.Printf(
							"Error converting line %d, alternative %d, field %d to number: %s\n",
							i, j, k, prod,
						)
						return
					}

					product.rules = append(product.rules, num)
				}
				products = append(products, product)
			}

			rules[ruleID] = products
		} else {

			fmt.Println("Processing", line)
			valid, err := parse(line, 0, rules)
			if err != nil {
				fmt.Println(err)
				return
			}

			if valid {
				count++
			}
			fmt.Printf("%d/%d OK\n", count, i-len(rules))

		}
	}

	fmt.Println(count, "matches")
}

func parse(input string, goal int, rules Rules) (bool, error) {

	cache = NewCache()
	convertedInput := make([]int, 0)

	/* Find rule for each terminal symbol */
	for _, ch := range input {
		matchedAny := false
		for ruleID, rule := range rules {
			if len(rule) > 1 {
				continue
			}
			prod := rule[0]
			if !prod.terminal {
				continue
			}

			if prod.symbol == string(ch) {
				convertedInput = append(convertedInput, ruleID)
				matchedAny = true
				break
			}
		}

		if !matchedAny {
			return false, fmt.Errorf(
				"Error finding rule for terminal symbol '%c' in '%s'",
				ch,
				input,
			)
		}
	}

	return find(convertedInput, goal, rules), nil
}

func find(input []int, goal int, rules Rules) bool {

	for i, _ := range input {
		for ruleID, products := range rules {
			/* Skip all rules with terminal symbols */
			if len(products) == 1 && products[0].terminal {
				continue
			}

			/* Ask cache if input+rule is a dead end */
			if cache.IsKnownBad(ruleID, input[i:]) {
				continue
			}

			for _, alt := range products {

				/* Find alternatives which apply */
				matched := true
				if len(alt.rules) > len(input[i:]) {
					continue
				}
				for j, altRuleID := range alt.rules {
					if j+i >= len(input) {
						matched = false
						break
					}

					/* Skip alternative if production doesn't match */
					if altRuleID != input[i+j] {
						/*if j > 0 {
							fmt.Println("Input differs from alternative",
								input, alt.rules,
							)
						}*/
						matched = false
						break
					}

				}

				if !matched {
					continue
				}

				/*
				 * Check if goal has been reached and get out of here.
				 * Not only has at least one alternative to match but
				 * also the input has to be empty if another reduction would
				 * take place.
				 */
				if ruleID == goal && len(alt.rules) == len(input) {
					return true
				}

				reduceIndex := len(alt.rules) + i

				/* Reduce the input using the matched alternative */
				// fmt.Println("Input before reduction", input)
				reducedInput := make(
					[]int,
					len(input[:i])+1+len(input[reduceIndex:]),
				)

				copy(reducedInput, input[:i])
				reducedInput[i] = ruleID
				copy(reducedInput[i+1:], input[reduceIndex:])

				/*
					fmt.Printf(
						"reduced=%#v, input=%#v, input[:%d]=%#v, ruleID=%d, input[%d:]=%#v\n",
						reducedInput,
						input,
						i,
						input[:i],
						ruleID,
						reduceIndex,
						input[reduceIndex:],
					)
				*/

				/* Go up and search for the light */
				if find(reducedInput, goal, rules) {
					// fmt.Println("[  OK   ]", input, "->", reducedInput)
					return true
				}
				// fmt.Println("[NOT OK]", input, "-/->", reducedInput)
			}

			/*
			 * No alternative of the rule matched. Add it to the
			 * cache of bad rules.
			 */
			cache.AddBad(ruleID, input[i:], products)
		}
	}

	return false
}
