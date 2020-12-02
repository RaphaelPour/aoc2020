package main

import (
	"fmt"
	"strconv"

	"github.com/RaphaelPour/aoc2020/util"
)

func main() {
	ex := make([]int, 0)

	lines := util.LoadDefault()

	for i, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("%s is not a number in line %d!\n", line, i)
			return
		}

		ex = append(ex, n)

	}

	for i := 0; i < len(ex); i++ {
		for j := i + 1; j < len(ex); j++ {
			if ex[i]+ex[j] == 2020 {
				fmt.Printf("Found %d + %d = 2020\n", ex[i], ex[j])
				fmt.Printf("%d x %d = %d\n", ex[i], ex[j], ex[i]*ex[j])
			}
		}
	}

}
