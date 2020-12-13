package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"strconv"
	"strings"
)

func main() {

	// re := regexp.MustCompile(`^([A-Z])(\d+)$`)

	lines := util.LoadString("input")
	earliestTs, err := strconv.Atoi(lines[0])
	if err != nil {
		fmt.Println("Earliest ts is not a number:", lines[0])
		return
	}

	timestamps := make([]int, 0)
	for _, cell := range strings.Split(lines[1], ",") {
		if cell == "x" {
			continue
		}

		num, err := strconv.Atoi(cell)
		if err != nil {
			fmt.Println("Cell is not a number", cell)
			return
		}

		timestamps = append(timestamps, num)
	}

	for time := earliestTs; true; time++ {
		for _, busTs := range timestamps {

			if time%busTs == 0 {
				fmt.Printf(
					"BusID=%d, time=%d, waiting=%d\n",
					busTs, time, time-earliestTs,
				)
				fmt.Println(">>", busTs*(time-earliestTs), "<<")
				return
			}
		}
	}

}
