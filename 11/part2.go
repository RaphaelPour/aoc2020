package main

import (
	"fmt"

	"github.com/RaphaelPour/aoc2020/util"
)

const (
	OCCUPIED = '#'
	EMPTY    = 'L'
	FLOOR    = '.'
)

func next(area []string) []string {

	newArea := make([]string, 0)

	for y := 0; y < len(area); y++ {
		newArea = append(newArea, "")
		for x := 0; x < len(area[0]); x++ {
			if area[y][x] == FLOOR {
				newArea[y] += string(FLOOR)
				continue
			}

			n := neighbours(x, y, area)
			if area[y][x] == EMPTY && n == 0 {
				newArea[y] += string(OCCUPIED)
			} else if area[y][x] == OCCUPIED && n >= 5 {
				newArea[y] += string(EMPTY)
			} else {
				newArea[y] += string(area[y][x])
			}
		}
	}

	return newArea
}

func countSeats(area []string) int {
	count := 0
	for y := 1; y < len(area)-1; y++ {
		for x := 1; x < len(area[0])-1; x++ {
			if area[y][x] == OCCUPIED {
				count++
			}
		}
	}
	return count
}

func neighbours(x, y int, area []string) int {
	count := 0

	seats := [][]rune{
		[]rune{FLOOR, FLOOR, FLOOR},
		[]rune{FLOOR, FLOOR, FLOOR},
		[]rune{FLOOR, FLOOR, FLOOR},
	}
	for dist := 1; dist < util.Max(len(area), len(area[0])); dist++ {
		for y1 := -1; y1 <= 1; y1++ {
			for x1 := -1; x1 <= 1; x1++ {
				if x1 == 0 && y1 == 0 {
					continue
				}
				// fmt.Printf("%d.%d => %d.%d : %c\n", x1, y1, x+x1, y+y1, area[y+y1][x+x1])
				x2 := x + x1*dist
				y2 := y + y1*dist
				if y2 >= len(area) || x2 >= len(area[0]) ||
					y2 < 0 || x2 < 0 {
					continue
				}
				if seats[y1+1][x1+1] != FLOOR {
					continue
				}
				seats[y1+1][x1+1] = rune(area[y2][x2])
			}
		}
	}

	// fmt.Printf("%d/%d:\n", x, y)
	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[0]); j++ {
			// fmt.Printf("%c", seats[i][j])
			if seats[i][j] == OCCUPIED {
				count++
			}
		}
		// fmt.Println("")
	}

	// fmt.Println(count, "-----")
	return count
}

func dump(area []string) {
	for _, line := range area {
		fmt.Println(line)
	}
}

func stopped(area1, area2 []string) bool {
	if len(area1) != len(area2) {
		fmt.Println("Lengths of area1 and 2 are different!", len(area1), len(area2))
		fmt.Println("Area1:", area1)
		fmt.Println("Area2:", area2)
		return false
	}

	for i, _ := range area1 {
		if area1[i] != area2[i] {
			return false
		}
	}

	return true
}

func main() {
	area := make([]string, 0)

	for _, line := range util.LoadString("input") {
		area = append(area, "."+line+".")
	}

	if len(area) == 0 {
		fmt.Println("Empty waiting room...")
		return
	}

	length := len(area[0])
	bar := ""
	for i := 0; i < length; i++ {
		bar += "."
	}

	area = append([]string{bar}, area...)
	area = append(area, bar)

	i := 0
	for true {
		if i%100 == 0 {
			fmt.Printf("Gen #%05d\n", i)
		}
		area2 := next(area)
		if stopped(area, area2) {
			break
		}
		area = area2
		i++
	}

	fmt.Println(countSeats(area))

}
