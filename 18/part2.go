package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"regexp"
	"strconv"
)

/*
 * Formula is a recursive decent parser evaluating
 * AoC math.
 *
 * Grammar:
 * E :- S | S*S
 * S :- F | F+F
 * F :- (E) | <int>
 * [E]xpression, [F]actor, [S]ummand
 *
 * https://en.wikipedia.org/wiki/Recursive_descent_parser
 */

type Formula struct {
	input []string
	index int
}

func (f Formula) String() string {
	return fmt.Sprintf("%3d: %v\n", f.index, f.input[f.index:])
}

func (f Formula) current() string {
	return f.input[f.index]
}

func (f Formula) currentNumber() int {
	num, err := strconv.Atoi(f.current())
	if err != nil {
		fmt.Println("Error converting", f.current(), "to integer")
		return -1
	}
	return num
}

func (f *Formula) accept(expection string) bool {
	if f.current() == expection {
		f.consume()
		return true
	}
	return false
}

func (f *Formula) consumeNumber() int {
	num := f.currentNumber()
	f.consume()
	return num
}

func (f *Formula) consume() string {
	f.index++
	return f.input[f.index-1]
}

func (f Formula) isEmpty() bool {
	return f.index >= len(f.input)-1
}

func (f *Formula) Parse() (int, error) {
	result, err := f.Expression()
	if err != nil {
		return -1, err
	}

	if !f.isEmpty() {
		return -1, fmt.Errorf("Expected input to be empty:", f)
	}
	return result, nil
}

func (f *Formula) Factor() (int, error) {
	if f.accept("(") {
		num, err := f.Expression()
		if err != nil {
			return -1, err
		}
		if !f.accept(")") {
			return -1, fmt.Errorf(
				"Syntax error: Expected ')', got '%s'",
				f.current(),
			)
		}
		return num, nil
	}

	return f.consumeNumber(), nil
}

func (f *Formula) Expression() (int, error) {
	num, err := f.Summand()
	if err != nil {
		return -1, err
	}

	for !f.isEmpty() && f.accept("*") {
		secondFactor, err := f.Summand()
		if err != nil {
			return -1, err
		}
		num *= secondFactor
	}

	return num, nil

}

func (f *Formula) Summand() (int, error) {
	num, err := f.Factor()
	if err != nil {
		return -1, err
	}

	for !f.isEmpty() && f.accept("+") {
		secondSummand, err := f.Factor()
		if err != nil {
			return -1, err
		}
		num += secondSummand
	}
	return num, nil
}

func main() {

	sum := 0
	re := regexp.MustCompile(`[\d+\(\)\+\*]`)
	for _, line := range util.LoadString("input") {
		f := Formula{input: re.FindAllString(line, -1)}
		result, err := f.Parse()
		if err != nil {
			fmt.Println(err)
			return
		}
		sum += result
	}
	fmt.Println(sum)
}
