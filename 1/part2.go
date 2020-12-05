package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
)

func main() {
	ex := util.LoadDefaultInt()

	for i := 0; i < len(ex); i++ {
		for j := i + 1; j < len(ex); j++ {
			for k := j + 1; k < len(ex); k++ {
				if ex[i]+ex[j]+ex[k] == 2020 {
					fmt.Println( ex[i]*ex[j]*ex[k])
				}
			}
		}
	}
}
