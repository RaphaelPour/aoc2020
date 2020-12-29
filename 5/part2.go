package main

import (
	"fmt"
	"regexp"

	"github.com/RaphaelPour/aoc2020/util"
)

/* Reverse to util */

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

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
	ids := make([]int, 0)

	for i := 0; i < 128; i++ {
		seatsOccupied[i] = make(map[int]bool, 0)
		for j := 0; j < 8; j++ {
			seatsOccupied[i][j] = false
		}
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
		ids = append(ids, BoardingID(row, col))
	}

	for row := 1; row < 127; row++ {
		for col := 0; col < 8; col++ {
			if !seatsOccupied[row][col] {
				id := BoardingID(row, col)

				id1found := false
				id2found := false
				for i := range ids {
					if ids[i] == id+1 {
						id1found = true
					}
					if ids[i] == id-1 {
						id2found = true
					}

					if id1found && id2found {
						fmt.Println(id)
						return
					}
				}
			}
		}
	}
}
