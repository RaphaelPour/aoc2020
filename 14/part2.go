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
	binMap = map[bool]string{true: "1", false: "0"}
)

type Assignment struct {
	address string
	value   uint64
}

type Program struct {
	mask    string
	assigns []Assignment
}

func (p *Program) Run() {

	for _, as := range p.assigns {
		floatings := make([]int, 0)

		/* Calculate address bitmask */
		addressMask := ""
		for i, _ := range p.mask {
			if p.mask[i] == '0' {
				addressMask += string(as.address[i])
			} else if p.mask[i] == '1' {
				addressMask += "1"
			} else if p.mask[i] == 'X' {
				addressMask += "X"
				floatings = append(floatings, i)
			}
		}

		/*
			fmt.Println("     Address:", as.address)
			fmt.Println("        Mask:", p.mask)
			fmt.Println("Address mask:", addressMask)
			fmt.Println("   Floatings:", floatings)
		*/

		for k := 0; k < int(math.Pow(2, float64(len(floatings)))); k++ {
			addr := []rune(addressMask)
			for i, float := range floatings {
				if k&(1<<i) != 0 {
					// fmt.Println("Change index", float, ":", addr[float], "-> 1")
					addr[float] = '1'
				} else {
					// fmt.Println("Change index", float, ":", addr[float], "-> 0")
					addr[float] = '0'
				}

				/*
					// Unset bit in any case
					addr &= uint64(^(1 << (i)))

					// Set bit again (if any)
					addr |= uint64(index << (i))
				*/
			}
			/*
				fmt.Print("Final address:")
				dumpRunes(addr)
				fmt.Println("")
			*/

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

func dumpRunes(runes []rune) {
	for _, r := range runes {
		fmt.Printf("%c", r)
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

	fmt.Println("== [MEMORY] ==")
	for k, v := range memory {
		fmt.Printf("%6d: %d\n", k, v)
	}

	/* Sum up all left values */
	for _, value := range memory {
		result += value
	}

	/* TODO: Make util vor result, valid, invalid values */
	fmt.Println("Too low:", "87681422310")
	fmt.Println(">>", result, "<<")
}
