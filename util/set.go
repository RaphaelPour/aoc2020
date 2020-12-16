package util

func FindNextStringInList(needle string, haystack []string) int {
	for i, el := range haystack {
		if el == needle {
			return i
		}
	}
	return -1
}

func RemoveAllStringsFromList(needle string, haystack []string) []string {
	result := make([]string, 0)

	for _, el := range haystack {
		if el != needle {
			result = append(result, el)
		}
	}
	return result
}

func StrSubsetOf(set, subset []string) bool {
	for _, subel := range subset {
		found := false
		for _, el := range set {
			if el == subel {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func FindNextIntInList(needle int, haystack []int) int {
	for i, el := range haystack {
		if el == needle {
			return i
		}
	}
	return -1
}

func IntSubsetOf(set, subset []int) bool {
	for _, subel := range subset {
		found := false
		for _, el := range set {
			if el == subel {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}
