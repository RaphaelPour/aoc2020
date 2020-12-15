package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"github.com/fatih/color"
	"strconv"
	"strings"
)

type Spoken struct {
	turn, preTurn uint64
}

var (
	goodColor = color.New(color.FgGreen, color.Bold)
	badColor  = color.New(color.FgRed, color.Bold)
	inputFile = "input"
	solutions = map[string][]uint64{
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
	history1 := make(map[uint64]Spoken, 0)
	history2 := make(map[uint64]Spoken, 0)
	content := util.LoadString(inputFile)[0]
	for i, el := range strings.Split(content, ",") {
		num, err := strconv.ParseUint(el, 10, 64)
		if err != nil {
			fmt.Println(el, "is not a number in at position", i)
			return
		}
		history1[num] = Spoken{turn: uint64(i + 1), preTurn: 0}
		history2[num] = Spoken{turn: uint64(i + 1), preTurn: 0}
	}

	result := play(2020, history1)
	fmt.Println("Invalid: 768, 656")
	fmt.Println("part1: >>", result, "<<")
	checkSolution(result, 1)

	fmt.Println("Too high: 2002")
	result = play(30000000, history2)
	fmt.Println("part2: >>", result, "<<")
	checkSolution(result, 2)
	return

}

func play(rounds uint64, history map[uint64]Spoken) uint64 {

	next := uint64(0)
	for turn := uint64(len(history) + 1); turn < rounds; turn++ {
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

func checkSolution(result uint64, part int) {
	solution, ok := solutions[inputFile]
	if ok {
		if solution[part-1] == result {
			goodColor.Println("OK")
		} else {
			badColor.Println("NOT OK")
		}
	}
}
