package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"regexp"
	"strconv"
	"strings"
)

type Product struct {
	terminal bool
	rules    []int
	symbol   string
}

type Products []Product
type Rules map[int]Products

func main() {

	re := regexp.MustCompile(`^(\d+): (.*)$`)
	terminalRe := regexp.MustCompile(`"([a-z])"`)
	rules := make(Rules, 0)
	parseRules := true
	count := 0
	for i, line := range util.LoadString("input") {
		if parseRules {
			if line == "" {
				parseRules = false
				continue
			}

			/* HOT PATCH */
			if strings.HasPrefix(line, "8") {
				line = "8: 42 | 42 8"
			} else if strings.HasPrefix(line, "11") {
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
			fmt.Printf("\r%d/%d OK", count, i-len(rules))

		}
	}

	fmt.Println(count, "matches")
}

func parse(input string, goal int, rules Rules) (bool, error) {

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

			for _, alt := range products {

				/* Find alternatives which apply */
				matched := true
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
					reducedInput := append(input[:i], ruleID)
					reducedInput = append(
						reducedInput,
						input[reduceIndex:]...,
					)*/

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
		}
	}

	return false
}
