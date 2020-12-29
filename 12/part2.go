package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"math"
	"regexp"
	"strconv"
)

func main() {

	sx, sy := 0, 0
	wx, wy := 10, 1

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
			wy += val
		case "S":
			wy -= val
		case "E":
			wx += val
		case "W":
			wx -= val
		case "L":
			wx, wy = Rotate(wx, wy, val)
		case "R":
			wx, wy = Rotate(wx, wy, -val)
		case "F":
			sx += wx * val
			sy += wy * val
		default:
			fmt.Printf("Unknown action in line %d: %s\n", i, action)
			return
		}
	}

	fmt.Println(util.Abs(sx) + util.Abs(sy))
}

func Rotate(x, y, deg int) (int, int) {
	/*
	 * https://en.wikipedia.org/wiki/Rotation_%28mathematics%29#Two_dimensions
	 * x' = x*cos(alpha) - y*sin(alpha)
	 * y' = y*cos(alpha) + x*sin(alpha)
	 */
	return x*int(math.Cos(util.Radians(float64(deg)))) -
			y*int(math.Sin(util.Radians(float64(deg)))),
		y*int(math.Cos(util.Radians(float64(deg)))) +
			x*int(math.Sin(util.Radians(float64(deg))))
}
