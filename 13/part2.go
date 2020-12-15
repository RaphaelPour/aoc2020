package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"github.com/fatih/color"
	"math/big"
	"strconv"
	"strings"
)

/*
 */

var (
	goodColor = color.New(color.FgGreen, color.Bold)
	badColor  = color.New(color.FgRed, color.Bold)
	inputFile = "input"
	solutions = map[string]uint64{
		"input2": 1068781,
		"input3": 3417,
		"input4": 754018,
		"input5": 779210,
		"input6": 1261476,
		"input7": 1202161486,
	}
	one = big.NewInt(1)
)

/* https://rosettacode.org/wiki/Chinese_remainder_theorem#Go */
func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

func main() {

	/* Parse input. The first number can be ingored.
	 * The second one is a coma-separated list of busses with their
	 * departure interval
	 */
	fmt.Println("Input file:", inputFile)
	lines := util.LoadString(inputFile)
	n := make([]*big.Int, 0)
	a := make([]*big.Int, 0)
	i := int64(0)
	for _, cell := range strings.Split(lines[1], ",") {

		if cell == "x" {
			i++
			continue
		}

		num, err := strconv.ParseInt(cell, 10, 64)
		if err != nil {
			fmt.Println("Cell is not a number", cell)
			return
		}
		a = append(a, big.NewInt(-i))
		n = append(n, big.NewInt(num))
		i++
	}

	fmt.Println(crt(a, n))
	return

}
