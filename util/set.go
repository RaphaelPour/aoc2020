package util

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
