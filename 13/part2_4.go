package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"strconv"
	"strings"
)

/*
 * If all input numbers are primes, why can't we just multiply
 * them all together? They can't share any divider (since
 * they're prime).
 *
 * But if we try to solve the position problem from 2_2,2_3,
 * we'll get non-primes, which will might have common dividers
 * again.
 */

func main() {

	lines := util.LoadString("input")
	ts := make([]uint64, 0)
	i := 0
	max := uint64(0)
	prod := uint64(1)
	for _, cell := range strings.Split(lines[1], ",") {

		if cell == "x" {
			i++
			continue
		}

		num, err := strconv.Atoi(cell)
		if err != nil {
			fmt.Println("Cell is not a number", cell)
			return
		}

		fmt.Println(num, "-", i, "=", num-i)
		n := uint64(num - i)
		if n > max {
			max = n
		}
		ts = append(ts, n)

		prod *= n
		i++
	}

	fmt.Println("Product:", prod)

	fmt.Println(ts)
	divider := make([]uint64, 0)
	for n := uint64(2); n < max/2; n = nextPrime(n + 1) {
		if dividesAll(n, ts) {
			divider = append(divider, n)
		}
	}

	fmt.Println(divider)

	//fmt.Println("\r>>", result, "<<")
	return

}

func nextPrime(offset uint64) uint64 {
	var p uint64
	for p = offset; true; p++ {
		isPrime := true
		for n := uint64(2); n <= p/2; n++ {
			if p%n == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			break
		}
	}
	return p
}

func dividesAll(divider uint64, nums []uint64) bool {
	for _, num := range nums {
		if num%divider != 0 {
			return false
		}
	}
	return true
}
