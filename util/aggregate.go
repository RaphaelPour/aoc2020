package util

func Product(vec []int) int {
	result := 1
	for _, p := range vec {
		result *= p
	}
	return result
}

func Sum(vec []int) int {
	result := 0
	for _, p := range vec {
		result += p
	}
	return result
}

func MaxInts(vec []int) int {
	max := 0
	for _, x := range vec {
		if x > max {
			max = x
		}
	}

	return max
}

func MinInts(vec []int) int {
	min := int(^uint(0) >> 1)
	for _, x := range vec {
		if x < min {
			min = x
		}
	}

	return min
}

func MinMaxInts(vec []int) (int, int) {
	min := int(^uint(0) >> 1)
	max := 0
	for _, x := range vec {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}

	return min, max

}
