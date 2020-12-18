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
 * E :- F | F*F|F+F
 * F :- (E) | <int>
 * [E]xpression, [F]actor
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

func (f Formula) isNumber() bool {
	_, err := strconv.Atoi(f.current())
	return err == nil
}

func (f *Formula) consume() string {
	f.index++
	return f.input[f.index-1]
}

func (f Formula) isEmpty() bool {
	return f.index >= len(f.input)-1
}

func (f *Formula) Expression() (int, error) {
	factor, err := f.Factor()
	if err != nil {
		return -1, err
	}

	for !f.isEmpty() &&
		(f.current() == "*" || f.current() == "+" || f.isNumber()) {
		op := f.consume()
		secondFactor := 0
		var err error
		switch op {
		case "*":
			secondFactor, err = f.Factor()
			if err != nil {
				return -1, err
			}
			factor *= secondFactor
		case "+":
			secondFactor, err = f.Factor()
			if err != nil {
				return -1, err
			}
			factor += secondFactor
		default:
			return -1, fmt.Errorf("Syntax error. Unknown operator '%s'", op)
		}
	}

	return factor, nil
}

func (f *Formula) Factor() (int, error) {

	if f.current() == "(" {
		f.consume()
		num, err := f.Expression()
		if err != nil {
			return -1, err
		}
		if ch := f.consume(); ch != ")" {
			return -1, fmt.Errorf("Syntax error: Expected ')', got '%s'", ch)
		}
		return num, nil
	}

	num, err := strconv.Atoi(f.current())
	if err != nil {
		return -1, fmt.Errorf("Error converting", f.current(), "to integer")
	}
	f.consume()
	return num, nil
}

func (f *Formula) Parse() (int, error) {
	result, err := f.Expression()
	if err != nil {
		return -1, err
	}

	if !f.isEmpty() {
		return -1, fmt.Errorf(
			"Parser returned but input is not empty: %s",
			f,
		)
	}

	return result, nil
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
