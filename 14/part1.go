package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"regexp"
	"strconv"
	"strings"
)

var (
	binMap  = map[bool]uint64{false: 0, true: 1}
	runeMap = map[rune]uint64{'0': 0, '1': 1}
	memory  = map[uint64]uint64{}
)

type Assignment struct {
	address uint64
	value   uint64
}

type Program struct {
	mask    string
	assigns []Assignment
}

func (p *Program) Run() {
	for _, assign := range p.assigns {
		binValue := dec2Bin(assign.value)
		memory[assign.address] =
			AndX(
				binValue,
				p.mask,
				memory[assign.address],
			)
	}

}

func dec2Bin(value uint64) string {
	num := strconv.FormatUint(value, 2)

	for len(num) < 36 {
		num = "0" + num
	}
	return num
}

func AndX(value, mask string, acc uint64) uint64 {
	if len(value) > len(mask) {
		fmt.Println("Mask must be longer than the value")
	}

	for i := range mask {
		if mask[i] == 'X' {

			currentMask := uint64(1 << (36 - i - 1))
			if value[i] == '1' {
				acc |= currentMask
			} else {
				acc &= ^currentMask
			}
			continue
		}

		currentMask := uint64(1 << (36 - i - 1))
		if mask[i] == '1' {
			acc |= currentMask
		} else {
			acc &= ^currentMask
		}
		acc |= (binMap[mask[i] == '1'] << (36 - i - 1))
	}

	return acc
}

func main() {

	reMask := regexp.MustCompile(`^mask = ([X01]{36})$`)
	reMem := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

	programs := make([]Program, 0)
	for _, line := range util.LoadString("input") {
		if strings.HasPrefix(line, "mask") {
			match := reMask.FindStringSubmatch(line)
			if len(match) < 2 {
				fmt.Println("Error matching mask:", line)
				return
			}
			mask := match[1]
			programs = append(programs, Program{mask: mask})
		} else {
			match := reMem.FindStringSubmatch(line)
			if len(match) < 3 {
				fmt.Println("Error matching assignment:", "'", line, "'")
				fmt.Println("Got", match, "instead.")
				return
			}

			addr, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Println("Address is not a number:", match[1])
				return
			}

			val, err := strconv.Atoi(match[2])
			if err != nil {
				fmt.Println("Value is not a number:", match[2])
				return
			}

			assign := Assignment{address: uint64(addr), value: uint64(val)}

			programs[len(programs)-1].assigns = append(
				programs[len(programs)-1].assigns,
				assign,
			)
		}

	}

	/* Run the programs */
	result := uint64(0)
	for _, program := range programs {
		program.Run()
	}

	/* Sum up all left values */
	for _, value := range memory {
		result += value
	}

	fmt.Println(result)
}
