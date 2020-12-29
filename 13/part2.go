package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"math/big"
	"strconv"
	"strings"
)

var (
	inputFile = "input"
	one       = big.NewInt(1)
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

	result, _ := crt(a, n)
	fmt.Println(result)
}
