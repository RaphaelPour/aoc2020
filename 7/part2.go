package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"regexp"
	"strconv"
)

type Node map[string]int
type Tree map[string]Node

func (t Tree) DepthSearch(current string) int {
	total := 1
	for subNode, count := range t[current] {
		total += count * t.DepthSearch(subNode)
	}
	return total
}

func main() {
	lineRegex := regexp.MustCompile(`^(\w+\s\w+) bags contain (.*).$`)
	bagRegex := regexp.MustCompile(`((\d) (\w+\s\w+)|no other)`)

	tree := make(Tree, 0)
	for i, line := range util.LoadString("input") {
		lineMatch := lineRegex.FindStringSubmatch(line)

		if len(lineMatch) < 3 {
			fmt.Printf("Line %d doesn't match: %s\n", i, line)
			return
		}

		node := make(Node, 0)

		bagMatches := bagRegex.FindAllStringSubmatch(lineMatch[2], -1)

		for j, bagMatch := range bagMatches {
			if len(bagMatch) < 3 {
				fmt.Printf(
					"Content line %d doesn't match: %s\n",
					j,
					bagMatch,
				)
				return
			}

			/* Check if bag contains other bags */
			if bagMatch[1] != "no other" {
				count, err := strconv.Atoi(bagMatch[2])
				if err != nil {
					fmt.Printf(
						"%s is not a number from line '%s'\n",
						bagMatch[2],
						bagMatch,
					)
					return
				}
				node[bagMatch[3]] = count
			}

		}

		tree[lineMatch[1]] = node

	}

	/*
	 * Count all bags within the shiny gold and subtract one which is
	 * the shiny one itself
	 */
	fmt.Println(tree.DepthSearch("shiny gold") - 1)
}
