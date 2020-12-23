package main

import (
	"fmt"

	"github.com/RaphaelPour/aoc2020/util"
)

var (
	inputKey = "input"
	inputs   = map[string][]int{
		"input":  []int{7, 3, 9, 8, 6, 2, 5, 4, 1},
		"input2": []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
	}
	rounds = map[string][][]int{
		"input2": {
			{3, 8, 9, 1, 2, 5, 4, 6, 7},
			{3, 2, 8, 9, 1, 5, 4, 6, 7},
			{3, 2, 5, 4, 6, 7, 8, 9, 1},
			{7, 2, 5, 8, 9, 1, 3, 4, 6},
			{3, 2, 5, 8, 4, 6, 7, 9, 1},
			{9, 2, 5, 8, 4, 1, 3, 6, 7},
			{7, 2, 5, 8, 4, 1, 9, 3, 6},
			{8, 3, 6, 7, 4, 1, 9, 2, 5},
			{7, 4, 1, 5, 8, 3, 9, 2, 6},
			{5, 7, 4, 1, 8, 3, 9, 2, 6},
			{5, 8, 3, 7, 4, 1, 9, 2, 6},
		},
	}
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
	// fmt.Println("current:", currentCup)

	// fmt.Println(c)

	picked := c.Pop3((c.currentIndex + 1) % len(c.cups))
	/* fmt.Print("pick up: ")
	for i := range picked {
		fmt.Printf("%d ", picked[i])
	}
	fmt.Println("")*/

	destIndex := c.DestinationCupIndex(currentCup)
	if destIndex == -1 {
		return fmt.Errorf("Couldn't find any destination cup... Wanted %d", c.cups[c.currentIndex])
	}
	dest := c.cups[destIndex]
	// fmt.Println("destination:", dest)
	// c.Pop(destIndex)

	for i := 0; i < len(c.cups); i++ {
		if c.cups[i] == dest {
			destIndex = i
			break
		}
	}

	c.Insert3(destIndex+1, picked)
	//c.Insert(dest, (c.currentIndex+1)%len(c.cups))

	c.AlignCurrentIndex(currentCup)
	c.currentIndex = (c.currentIndex + 1) % len(c.cups)

	// fmt.Println("")
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

func (c Circle) CupsAfterOne() int {
	oneIndex := 0
	for i := range c.cups {
		if c.cups[i] == 1 {
			oneIndex = i
			break
		}
	}

	cup1 := c.cups[oneIndex+1%len(c.cups)]
	cup2 := c.cups[oneIndex+2%len(c.cups)]

	fmt.Println("Next cups after 1:", cup1, cup2)

	return cup1 * cup2
}

func EqualSets(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {

	circle := Circle{cups: inputs[inputKey]}

	next := util.Max(circle.cups...) + 1
	for len(circle.cups) < 1000000 {
		circle.cups = append(circle.cups, next)
		next++
	}

	for i := 0; i < 10000000; i++ {
		if i%100 == 0 {
			fmt.Printf("\r%.5f%%", 100.0/float64(10000000)*float64(i))
		}
		/*if !EqualSets(circle.cups, rounds[inputKey][i]) {
			fmt.Println("Expected:", rounds[inputKey][i])
			fmt.Println("     Got:", circle.cups)
			return
		}*/
		if err := circle.Next(); err != nil {
			fmt.Println(err)
			return
		}

	}

	fmt.Println("-- final --")
	fmt.Println(circle)
	//fmt.Println(">>>", circle.Order(), "<<<")

	//fmt.Println("Wrong: 9,4,2")
	fmt.Println(">>>", circle.CupsAfterOne(), "<<<")
}
