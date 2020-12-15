package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"github.com/fatih/color"
	"strconv"
	"strings"
)

type Spoken struct {
	turn, preTurn int
}

var (
	goodColor = color.New(color.FgGreen, color.Bold)
	badColor  = color.New(color.FgRed, color.Bold)
	inputFile = "input"
	solutions = map[string][]int{
		"input":  {496, 883},
		"input2": {436, 175594},
		"input3": {1, 2578},
		"input4": {10, 3544142},
		"input5": {27, 261214},
		"input6": {78, 6895259},
		"input7": {438, 18},
		"input8": {1836, 362},
		"input9": {21, 27},
	}
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

	checkSolution(play(2020, history1), 1)
	checkSolution(play(30000000, history2), 2)

	return

}

func play(rounds int, history map[int]Spoken) int {

	next := int(0)
	for turn := int(len(history) + 1); turn < rounds; turn++ {
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

func checkSolution(result int, part int) {

	fmt.Printf("%d: ", part)

	solution, ok := solutions[inputFile]
	if ok {
		if solution[part-1] == result {
			goodColor.Println(result)
		} else {
			badColor.Println(result)
		}
	} else {
		fmt.Println(result)
	}
}
