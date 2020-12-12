package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"regexp"
	"strings"
)

type Node struct {
	name string
	next []string
}

type Tree map[string]Node

func depthSearch(current, goal string, tree Tree) bool {
	if current == goal {
		return true
	}

	for _, subNode := range tree[current].next {
		if depthSearch(subNode, goal, tree) {
			return true
		}
	}
	return false
}

func main() {
	re := regexp.MustCompile(`^(\w+\s\w+) bags contain (.*).$`)
	re2 := regexp.MustCompile(`^\s?((\d) (\w+\s\w+)|no other)`)

	tree := make(Tree, 0)
	for i, line := range util.LoadString("input") {
		match := re.FindStringSubmatch(line)

		if len(match) < 3 {
			fmt.Printf("Line %d doesn't match: %s\n", i, line)
			return
		}

		// Overstep goal itself
		if match[1] == "shiny gold" {
			continue
		}

		contentList := strings.Split(match[2], ",")

		node := Node{name: match[1]}
		node.next = make([]string, 0)

		for j, content := range contentList {
			contentMatch := re2.FindStringSubmatch(content)

			if len(contentMatch) < 3 {
				fmt.Printf("Content line %d doesn't match: %s\n", j, content)
				return
			}

			/* Check if bag contains other bags */
			if len(contentMatch) > 3 {
				node.next = append(node.next, contentMatch[3])
			}

		}

		tree[match[1]] = node

	}

	found := 0
	for name, _ := range tree {
		if depthSearch(name, "shiny gold", tree) {
			found++
		}
	}

	fmt.Println(found)
}
