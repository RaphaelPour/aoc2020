package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Setup() {
	cache = NewCache()
	cache.debug = true
}

func TestFind1(t *testing.T) {
	Setup()

	input := []int{1}
	goal := 0
	rules := Rules{0: Products{Product{rules: []int{1}}}}
	require.True(t, find(input, goal, rules))
}

func TestFind2(t *testing.T) {
	Setup()

	input := []int{1, 1}
	goal := 0
	rules := Rules{0: Products{Product{rules: []int{1, 1}}}}
	require.True(t, find(input, goal, rules))
}

func TestFind3(t *testing.T) {
	Setup()

	input := []int{2}
	goal := 0
	rules := Rules{
		0: Products{Product{rules: []int{1}}},
		1: Products{Product{rules: []int{2}}},
	}
	require.True(t, find(input, goal, rules))
}

func TestFind4(t *testing.T) {
	Setup()

	input := []int{1, 2}
	goal := 0
	rules := Rules{
		0: Products{Product{rules: []int{1, 1}}},
		1: Products{Product{rules: []int{2}}},
	}
	require.True(t, find(input, goal, rules))
}

func TestFind5(t *testing.T) {
	Setup()

	/* Test alternatives where first one leads to the goal */
	input := []int{1, 2}
	goal := 0
	rules := Rules{
		0: Products{
			Product{rules: []int{1, 2}},
			Product{rules: []int{2, 1}},
		},
	}
	require.True(t, find(input, goal, rules))
}

func TestFind6(t *testing.T) {
	Setup()

	/* Test alternatives where second one leads to the goal */
	input := []int{1, 2}
	goal := 0
	rules := Rules{
		0: Products{
			Product{rules: []int{2, 1}},
			Product{rules: []int{1, 2}},
		},
	}
	require.True(t, find(input, goal, rules))
}

func TestFind7(t *testing.T) {
	Setup()

	/* Test deeper alternative productions */
	input := []int{4, 5}
	goal := 0
	rules := Rules{
		0: Products{
			Product{rules: []int{2, 1}},
			Product{rules: []int{1, 2}},
		},
		1: Products{
			Product{rules: []int{4, 3}},
			Product{rules: []int{4}},
		},
		2: Products{
			Product{rules: []int{5, 3}},
			Product{rules: []int{5}},
		},
	}
	require.True(t, find(input, goal, rules))
}

func TestFind8(t *testing.T) {
	Setup()

	/* Test loops */
	input := []int{2, 2, 1}
	goal := 0
	rules := Rules{
		0: Products{
			Product{rules: []int{1}},
		},
		1: Products{
			Product{rules: []int{2, 1}},
			Product{rules: []int{2}},
		},
	}
	require.True(t, find(input, goal, rules))
}

func TestFindExample1(t *testing.T) {
	Setup()

	input := []int{1, 3, 1}
	goal := 0
	rules := Rules{
		0: Products{Product{rules: []int{1, 2}}},
		2: Products{
			Product{rules: []int{1, 3}},
			Product{rules: []int{3, 1}},
		},
	}
	require.True(t, find(input, goal, rules))
}

func TestFindBad1(t *testing.T) {
	Setup()

	/* Empty input */
	input := []int{}
	goal := 0
	rules := Rules{0: Products{Product{rules: []int{1}}}}
	require.False(t, find(input, goal, rules))
}

func TestFindBad2(t *testing.T) {
	Setup()

	/* Unreachable goal */
	input := []int{1}
	goal := 10
	rules := Rules{0: Products{Product{rules: []int{1}}}}
	require.False(t, find(input, goal, rules))
}

func TestFindBad3(t *testing.T) {
	Setup()

	/* Missing rules */
	input := []int{1}
	goal := 0
	rules := Rules{}
	require.False(t, find(input, goal, rules))
}

func TestFindBad4(t *testing.T) {
	Setup()

	/* Input not matching any rule */
	input := []int{2}
	goal := 0
	rules := Rules{0: Products{Product{rules: []int{1}}}}
	require.False(t, find(input, goal, rules))
}

func TestFindBad5(t *testing.T) {
	Setup()

	/* No rule reducing two variables */
	input := []int{2, 2}
	goal := 0
	rules := Rules{
		0: Products{Product{rules: []int{1}}},
		1: Products{Product{rules: []int{2}}},
	}
	require.False(t, find(input, goal, rules))
}

func TestCache1(t *testing.T) {
	Setup()

	require.False(t, cache.IsKnownBad(0, []int{1}))
}

func TestCache2(t *testing.T) {
	Setup()

	cache.AddBad(0, []int{1}, Products{
		Product{rules: []int{1}},
	})
	require.True(t, cache.IsKnownBad(0, []int{1}))
}
