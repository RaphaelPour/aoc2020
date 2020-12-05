package util

func Reverse(in string) string {
	result := ""
	for _, char := range in {
		result = string(char) + result
	}
	return result
}
