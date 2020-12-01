package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	ex := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)

	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("%s is not a number in line %d!\n", line, index)
			return
		}

		ex = append(ex, n)

		index++
	}

	for i := 0; i < len(ex); i++ {
		for j := i + 1; j < len(ex); j++ {
			for k := j + 1; k < len(ex); k++ {
				if ex[i]+ex[j]+ex[k] == 2020 {
					fmt.Printf("Found %d + %d + %d = 2020\n", ex[i], ex[j], ex[k])
					fmt.Printf("%d x %d x %d = %d\n", ex[i], ex[j], ex[k], ex[i]*ex[j]*ex[k])
				}
			}
		}
	}
}
