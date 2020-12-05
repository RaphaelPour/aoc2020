package util

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Min(nums ...int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func Max(nums ...int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func InRange(num, min, max int) bool {
	return num >= min && num <= max
}
