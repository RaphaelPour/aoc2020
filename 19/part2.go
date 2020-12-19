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
	for i, line := range util.LoadString("input4") {
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
			valid, rest, err := parse(line, 0, rules)

			if err != nil {
				fmt.Println(err)
				return
			}

			if valid && len(rest) == 0 {
				fmt.Println("[ MATCH ]", line)
				count++
			} else {
				fmt.Println("[INVALID]", line)
			}
		}
	}

	// fmt.Println(rules)

	fmt.Println("Too low: 150")
	fmt.Println(count, "matches")
}

func parse(input string, ruleID int, rules Rules) (bool, string, error) {

	if len(input) == 0 {
		return false, "", nil
	}

	rule, ok := rules[ruleID]
	if !ok {
		return false, "", fmt.Errorf("Parsing error: unknown rule %d", ruleID)
	}

	// fmt.Println(ruleID, rule)
	for _, alt := range rule {
		if alt.terminal {
			accepted := string(input[0]) == alt.symbol
			/*if accepted {
				fmt.Printf("[CONSUME] %c from %s\n", input[0], input)
			}*/
			return accepted, input[1:], nil
		}

		matchedAll := true
		rest := input
		for _, substRule := range alt.rules {

			var err error
			var valid bool
			valid, rest, err = parse(rest, substRule, rules)
			if err != nil {
				return false, input, err
			}
			if !valid {
				matchedAll = false
				break
			}
		}

		if matchedAll {
			return true, rest, nil
		}
	}

	return false, input, nil
}
