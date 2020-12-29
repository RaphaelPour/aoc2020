package main

import (
	"fmt"

	"github.com/RaphaelPour/aoc2020/util"
)

var (
	input = []int{7, 3, 9, 8, 6, 2, 5, 4, 1}
)

type Circle struct {
	cups         []int
	currentIndex int
}

func (c Circle) String() string {
	out := "cups: "
	for i := range c.cups {
		if i == c.currentIndex {
			out += fmt.Sprintf("(%d)", c.cups[i])
		} else {
			out += fmt.Sprintf(" %d ", c.cups[i])
		}
	}
	return out
}

func (c *Circle) Pop(i int) int {
	cup := c.cups[i]

	if i+1 == len(c.cups) {
		c.cups = c.cups[:i]
	} else {
		c.cups = append(c.cups[:i], c.cups[i+1:]...)
	}
	return cup
}

func (c *Circle) Pop3(index int) []int {
	popped := make([]int, 3)
	for i := 0; i < 3; i++ {
		if index >= len(c.cups) {
			index = 0
		}
		popped[i] = c.Pop(index % len(c.cups))
	}
	return popped
}

func (c *Circle) Insert(cup, index int) {
	cups := make([]int, len(c.cups)+1)
	copy(cups, c.cups[:index])
	cups[index] = cup
	copy(cups[index+1:], c.cups[index:])
	c.cups = cups
}

func (c *Circle) Insert3(index int, cups []int) {
	for i := range cups {
		c.Insert(cups[i], index+i)
	}
}

func (c *Circle) AlignCurrentIndex(cup int) {
	cupIndex := 0
	for i := 0; i < len(c.cups); i++ {
		if c.cups[i] == cup {
			cupIndex = i
			break
		}
	}

	c.currentIndex = cupIndex
}

func (c Circle) DestinationCupIndex(goal int) int {
	for {
		goal--
		if goal < 0 {
			goal = util.Max(c.cups...)
		}

		for i := 0; i < len(c.cups); i++ {
			if c.cups[i] == goal {
				return i
			}
		}
	}
}

func (c *Circle) Next() error {

	currentCup := c.cups[c.currentIndex]

	picked := c.Pop3((c.currentIndex + 1) % len(c.cups))

	destIndex := c.DestinationCupIndex(currentCup)
	if destIndex == -1 {
		return fmt.Errorf("Couldn't find any destination cup... Wanted %d", c.cups[c.currentIndex])
	}
	dest := c.cups[destIndex]

	for i := 0; i < len(c.cups); i++ {
		if c.cups[i] == dest {
			destIndex = i
			break
		}
	}

	c.Insert3(destIndex+1, picked)

	c.AlignCurrentIndex(currentCup)
	c.currentIndex = (c.currentIndex + 1) % len(c.cups)

	return nil
}

func (c Circle) Order() string {
	oneIndex := 0
	for i := range c.cups {
		if c.cups[i] == 1 {
			oneIndex = i
			break
		}
	}

	result := ""
	for i := 1; i < len(c.cups); i++ {
		index := (oneIndex + i) % len(c.cups)
		result += fmt.Sprintf("%d", c.cups[index])
	}
	return result
}

func main() {
	circle := Circle{cups: input}
	for i := 0; i < 100; i++ {
		if err := circle.Next(); err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(circle.Order())
}
