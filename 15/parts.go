package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"strconv"
	"strings"
)

type Spoken struct {
	turn, preTurn int
}

var (
	inputFile = "input"
)

func main() {
	history1 := make(map[int]Spoken, 0)
	history2 := make(map[int]Spoken, 0)
	content := util.LoadString(inputFile)[0]
	for i, el := range strings.Split(content, ",") {
		num, err := strconv.Atoi(el)
		if err != nil {
			fmt.Println(el, "is not a number in at position", i)
			return
		}
		history1[num] = Spoken{turn: int(i + 1), preTurn: 0}
		history2[num] = Spoken{turn: int(i + 1), preTurn: 0}
	}

	fmt.Println(play(2020, history1))
	fmt.Println(play(30000000, history2))
}

func play(rounds int, history map[int]Spoken) int {
	next := 0
	for turn := len(history) + 1; turn < rounds; turn++ {
		if _, ok := history[next]; !ok {
			history[next] = Spoken{turn: turn, preTurn: 0}
			next = 0
			continue
		}
		history[next] = Spoken{
			preTurn: history[next].turn,
			turn:    turn,
		}
		next = history[next].turn - history[next].preTurn
	}

	return next
}
