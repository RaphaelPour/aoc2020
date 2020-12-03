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
