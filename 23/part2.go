package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/RaphaelPour/aoc2020/util"
)

var (
	inputKey = "input"
	inputs   = map[string][]int{
		"input":  []int{7, 3, 9, 8, 6, 2, 5, 4, 1},
		"input2": []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
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
	/*cups := make([]int, len(c.cups)+1)
	copy(cups, c.cups[:index])
	cups[index] = cup
	copy(cups[index+1:], c.cups[index:])
	c.cups = cups
	*/
	c.cups = append(c.cups, -1)
	copy(c.cups[index+1:], c.cups[index:])
	c.cups[index] = cup

}

func (c *Circle) Insert3(index int, cups []int) {
	c.Insert(cups[0], index+0)
	c.Insert(cups[1], index+1)
	c.Insert(cups[2], index+2)
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

func (c Circle) CupsAfterOne() int {
	oneIndex := 0
	for i := range c.cups {
		if c.cups[i] == 1 {
			oneIndex = i
			break
		}
	}

	cup1 := c.cups[(oneIndex+1)%len(c.cups)]
	cup2 := c.cups[(oneIndex+2)%len(c.cups)]

	fmt.Println("Next cups after 1:", cup1, cup2)

	return cup1 * cup2
}
func main() {
	f, err := os.Create(fmt.Sprintf("part2_%s.profile", time.Now().Format(time.RFC3339)))
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	cups := make([]int, 1000000)
	copy(cups, inputs[inputKey])

	next := util.Max(inputs[inputKey]...) + 1
	for i := len(inputs[inputKey]); i < len(cups); i++ {
		cups[i] = next
		next++
	}

	circle := Circle{cups: cups}

	for i := 0; i < 10000000; i++ {
		if i%100000 == 0 {
			fmt.Print(".")
		}
		if err := circle.Next(); err != nil {
			fmt.Println(err)
			return
		}
	}
	//fmt.Println(circle.Order())
	fmt.Println(">>>", circle.CupsAfterOne(), "<<<")
}
