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

type VM struct {
	program        Program
	acc, pc, Patch int
}

func (v VM) Run() bool {
	v.acc = 0
	v.pc = 0

	visited := make(map[int]bool, 0)
	for v.pc < len(v.program) {

		// Ceck if command was already executed one time
		if _, ok := visited[v.pc]; ok {
			fmt.Printf("%d already visited. Acc=%d\n", v.pc, v.acc)
			return false
		}
		visited[v.pc] = true

		command := v.program[v.pc].command
		if v.pc == v.Patch {
			if command == "jmp" {
				command = "nop"
			} else {
				command = "jmp"
			}
		}

		switch command {
		case "acc":
			v.acc += v.program[v.pc].value
		case "jmp":
			v.pc += v.program[v.pc].value
			continue
		case "nop":
		default:
			fmt.Println("Unknown instruction", command)
		}

		v.pc++
	}

	fmt.Printf("Program terminated with acc=%d\n", v.acc)
	return true
}

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

	for i, ins := range program {
		if ins.command != "acc" {
			vm := VM{program: program, Patch: i}
			if vm.Run() {
				return
			}
		}
	}

}
