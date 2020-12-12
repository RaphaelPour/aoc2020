package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"math"
	"regexp"
	"strconv"
)

func main() {

	x, y := 0, 0
	angle := 0

	re := regexp.MustCompile(`^([A-Z])(\d+)$`)

	for i, line := range util.LoadString("input") {
		match := re.FindStringSubmatch(line)

		if len(match) < 3 {
			fmt.Printf("Error parsing line %d: %s\n", i, line)
			return
		}
		action := match[1]
		val, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Printf("error parsing %s: Not a number\n", match[2])
			return
		}

		switch action {
		case "N":
			y += val
		case "S":
			y -= val
		case "E":
			x += val
		case "W":
			x -= val
		case "L":
			angle += val
		case "R":
			angle -= val
		case "F":
			x += val * int(math.Cos(util.Radians(float64(angle))))
			y += val * int(math.Sin(util.Radians(float64(angle))))
		default:
			fmt.Printf("Unknown action in line %d: %s\n", i, action)
			return
		}

		fmt.Printf("%s => %d/%d %d°\n", line, x, y, angle)
	}

	fmt.Printf("|%d| + |%d| = %d\n", x, y, util.Abs(x)+util.Abs(y))
}
