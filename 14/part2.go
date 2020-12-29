package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	memory = map[uint64]uint64{}
)

/* Assignment:
 *
 * Representation of a 'mem[address] = value' line from the input.
 */
type Assignment struct {
	address string
	value   uint64
}

/* Program:
 *
 * A program is a collection of assignments with the same bitmask.
 * This is not the same as the 'program' from the AoC description
 * which refers to the input as a whole.
 */
type Program struct {
	mask    string
	assigns []Assignment
}

func (p *Program) Run() {

	for _, as := range p.assigns {
		/* Indices of all 'floatings' from the bitmask */
		floatings := make([]int, 0)

		/* Calculate address bitmask */
		addressMask := ""
		for i, _ := range p.mask {

			/* Apply the bitmask to the current assignment's
			 * address based on the bitmask as follows:
			 *
			 * 0: Use the original value
			 * 1: Use 1
			 * X: Add the current index to the list of 'floatings'
			 *
			 * Note that the 'apply' happens by appending the
			 * appropriate 0/1/X to a new string since we can't
			 * replace a char by index in GoLang (yayyy).
			 */
			if p.mask[i] == '0' {
				addressMask += string(as.address[i])
			} else if p.mask[i] == '1' {
				addressMask += "1"
			} else if p.mask[i] == 'X' {

				/*
				 * This value doesn't really matter since it will be
				 * overwritten with the 'combinatorial iterator'
				 * below. It is set to a letter so the integer parser
				 * fails if the replacement didn't do it's job properly
				 * on replacing all 'floatings'.
				 */
				addressMask += "X"
				floatings = append(floatings, i)
			}
		}

		for k := 0; k < int(math.Pow(2, float64(len(floatings)))); k++ {
			addr := []rune(addressMask)
			for i, float := range floatings {
				if k&(1<<i) != 0 {
					addr[float] = '1'
				} else {
					addr[float] = '0'
				}

			}

			/* FINALLY, write the value to the address */
			numAddr, err := strconv.ParseUint(string(addr), 2, 64)
			if err != nil {
				fmt.Printf(
					"Error converting address %d: %s\n",
					addr,
					err,
				)
			}
			memory[numAddr] = as.value
		}
	}
}

func dec2Bin(value int64) string {
	num := strconv.FormatInt(value, 2)

	for len(num) < 36 {
		num = "0" + num
	}
	return num
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

			assign := Assignment{
				address: dec2Bin(int64(addr)),
				value:   uint64(val),
			}

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
