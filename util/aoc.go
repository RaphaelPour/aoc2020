package util

import (
	"fmt"
	"github.com/fatih/color"
)

var (
	GoodColor = color.New(color.FgGreen, color.Bold)
	BadColor  = color.New(color.FgRed, color.Bold)
)

type Solutions map[string][2]int

func (s Solutions) check(result, part int, inputFile string) {
	fmt.Printf("%d: ", part)

	solution, ok := s[inputFile]

	if ok {
		if solution[part-1] == result {
			GoodColor.Println(result)
		} else {
			BadColor.Println(result)
		}
	} else {
		fmt.Println(result)
	}
}
