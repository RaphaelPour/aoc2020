package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/RaphaelPour/aoc2020/util"
)

func isStrSubset(set, subset []string) bool {
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

func byr(v string) bool {
	if len(v) != 4 {
		return false
	}

	num, err := strconv.Atoi(v)
	if err != nil {
		fmt.Printf("%s is NaN\n", v)
		return false
	}

	if num < 1920 || num > 2002 {
		return false
	}

	return true
}

func iyr(v string) bool {
	if len(v) != 4 {
		return false
	}

	num, err := strconv.Atoi(v)
	if err != nil {
		fmt.Printf("%s is NaN\n", v)
		return false
	}

	if num < 2010 || num > 2020 {
		return false
	}
	return true
}

func eyr(v string) bool {
	if len(v) != 4 {
		return false
	}

	num, err := strconv.Atoi(v)
	if err != nil {
		fmt.Printf("%s is NaN\n", v)
		return false
	}

	if num < 2020 || num > 2030 {
		return false
	}
	return true
}

func hgt(v string) bool {
	re := regexp.MustCompile(`^(\d+)(cm|in)$`)
	match := re.FindStringSubmatch(v)

	if len(match) < 3 {
		fmt.Println("Something is missing in hgt:", v)
		return false
	}

	num, err := strconv.Atoi(match[1])
	if err != nil {
		fmt.Printf("%s is NaN\n", v)
		return false
	}

	if len(match) < 3 {
		fmt.Println("Something is missing..")
		return false
	}

	if match[2] == "cm" {
		if num < 150 || num > 193 {
			return false
		}
	} else if match[2] == "in" {
		if num < 59 || num > 76 {
			return false
		}
	} else {
		fmt.Println("Unknown unit", match[2])
	}

	return true
}

func hcl(v string) bool {
	if len(v) != 7 {
		return false
	}
	re := regexp.MustCompile(`#[\dabcdef]`)
	return re.MatchString(v)
}

func ecl(v string) bool {
	matches := 0
	for _, color := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if v == color {
			matches++
		}
	}
	return matches == 1
}

func pid(v string) bool {
	re := regexp.MustCompile(`^[\d]{9}$`)
	match := re.FindStringSubmatch(v)
	return len(match) > 0
}

func cid(v string) bool {
	return true
}

func main() {

	re := regexp.MustCompile(`(\w+):(#?\d?\w*)`)

	currentLine := ""
	valid := 0
	for _, line := range util.Load("input") {
		/* We just read a whole "passport" */
		if line == "" {
			kv := re.FindAllStringSubmatch(currentLine, -1)

			// fmt.Println(kv)
			check := 0
			for _, match := range kv {
				switch match[1] {
				case "ecl":
					if ecl(match[2]) {
						fmt.Println("ecl")
						check++
					}
				case "pid":
					if pid(match[2]) {
						fmt.Println("pid")
						check++
					}
				case "eyr":
					if eyr(match[2]) {
						fmt.Println("eyr")
						check++
					}
				case "hcl":
					if hcl(match[2]) {
						fmt.Println("hcl")
						check++
					}
				case "byr":
					if byr(match[2]) {
						fmt.Println("byr")
						check++
					}
				case "iyr":
					if iyr(match[2]) {
						fmt.Println("iyr")
						check++
					}
				case "hgt":
					if hgt(match[2]) {
						fmt.Println("hgt")
						check++
					}
				}
			}
			if check == 7 {
				valid++
			}

			currentLine = ""
		} else {
			currentLine += line
			currentLine += " "
		}
	}

	fmt.Println("Invalid: ", 2, 151, 138)
	fmt.Println("Valid: ", valid)
}
