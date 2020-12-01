package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(" Part I ====")
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
			if ex[i]+ex[j] == 2020 {
				fmt.Printf("Found %d + %d = 2020\n", ex[i], ex[j])
				fmt.Printf("%d x %d = %d\n", ex[i], ex[j], ex[i]*ex[j])
			}
		}
	}

}
