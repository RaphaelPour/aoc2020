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
	for i, line := range util.LoadString("input") {
		if parseRules {
			if line == "" {
				parseRules = false
				continue
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
		}
	}

	_, err := generate(0, rules)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(strings.Join(generated, "\n"))
}

func generate(ruleID int, rules Rules) ([]string, error) {
	result := make([]string, 0)

	rule, ok := rules[ruleID]
	if !ok {
		return nil, fmt.Errorf("Parsing error: unknown rule %d", ruleID)
	}

	for _, alt := range rule {
		if alt.terminal {
			result = append(result, alt.symbol)
			continue
		}

		subResult := make([]string, 0)
		for _, altID := range alt.rules {
			generated, err := generate(altID, rules)
			if err != nil {
				return nil, err
			}

			if len(subResult) == 0 {
				subResult = generated
			} else {
				newResult := make([]string, len(subResult)*len(generated))
				index := 0
				for _, a := range subResult {
					for _, b := range generated {
						newResult[index] = a + b
						index++
					}
				}
				subResult = newResult
			}
		}

		result = append(result, subResult...)
	}
	return result, nil
}
