package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func LoadDefaultString() []string {
	return LoadString("input")
}

func LoadString(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Error loading file '%s': %s\n", filename, err))
	}

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, string(scanner.Text()))
	}

	return lines
}

func LoadDefaultInt() []int {
	return LoadInt("input")
}

func LoadInt(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Error loading file '%s': %s\n", filename, err))
	}

	scanner := bufio.NewScanner(file)
	numbers := make([]int, 0)
	for scanner.Scan() {
		num, err := strconv.Atoi(string(scanner.Text()))
		if err != nil {
			panic(fmt.Sprintf("'%s' is not a number\n", scanner.Text()))
		}
		numbers = append(numbers, num)
	}

	return numbers
}