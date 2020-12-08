package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/RaphaelPour/aoc2020/util"
)

type Instruction struct {
	command string
	value   int
}

type Program []Instruction

func main() {

	re := regexp.MustCompile(`^(jmp|acc|nop) ([\+\-]\d+)`)

	program := make(Program, 0)
	for _, line := range util.LoadString("input") {
		match := re.FindStringSubmatch(line)

		if len(match) < 3 {
			fmt.Println("Error matching line", line)
			return
		}

		num, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Println(match[2], "is not a number")
			return
		}

		program = append(program, Instruction{match[1], num})

	}

	acc := 0
	visited := make(map[int]bool, 0)
	i := 0
	for true {
		fmt.Println(program[i])
		// Ceck if command was already executed one time
		if _, ok := visited[i]; ok {
			fmt.Printf("%d already visited. Acc=%d\n", i, acc)
			return
		}
		visited[i] = true

		switch program[i].command {
		case "acc":
			acc += program[i].value
		case "jmp":
			i += program[i].value
			continue
		case "nop":
		default:
			fmt.Println("Unknown instruction", program[i].command)
		}

		i++

	}
}
