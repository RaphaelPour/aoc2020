package main

import (
	"fmt"
	"regexp"

	"github.com/RaphaelPour/aoc2020/util"
)

/* Reverse to util */

func Bin2Dec(input string, One, Zero rune) int {
	result := 0

	for i, char := range input {
		if char == One {
			result |= (1 << (len(input) - i - 1))
		} else if char != Zero {
			fmt.Printf("Unknown char in input %c at pos %d\n", char, i)
		}
	}
	return result
}

func BoardingID(row, col int) int {
	return row*8 + col
}

func main() {

	seatsOccupied := make(map[int]map[int]bool, 0)
	for i := 0; i < 128; i++ {
		seatsOccupied[i] = make(map[int]bool, 0)
	}

	re := regexp.MustCompile(`([FB]+)([RL]+)`)
	for i, line := range util.LoadDefaultString() {
		match := re.FindStringSubmatch(line)
		if len(match) < 3 {
			fmt.Printf("Unknown line %d: %s\n", i, line)
			return
		}

		if len(match[1]) != 7 {
			fmt.Printf("Row in line %d has length != 7: %s\n", i, match[1])
			return
		}
		if len(match[2]) != 3 {
			fmt.Printf("Col in line %d has length != 3: %s\n", i, match[2])
			return
		}
		row := Bin2Dec(match[1], 'B', 'F')
		col := Bin2Dec(match[2], 'R', 'L')
		seatsOccupied[row][col] = true
	}

	finalID := 0
	for row := 0; row < 128; row++ {
		for col := 0; col < 8; col++ {
			if seatsOccupied[row][col] {
				result := BoardingID(row, col)
				if result > finalID {
					finalID = result
				}
			}
		}
	}

	fmt.Println(finalID)
}
