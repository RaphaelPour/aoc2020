package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/RaphaelPour/aoc2020/util"
)

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

var validators = map[string]func(v string) bool{
	"ecl": ecl,
	"pid": pid,
	"eyr": eyr,
	"hcl": hcl,
	"byr": byr,
	"iyr": iyr,
	"hgt": hgt,
}

func main() {

	re := regexp.MustCompile(`(\w+):(#?\d?\w*)`)

	currentLine := ""
	valid := 0
	for _, line := range util.Load("input") {
		/* We just read a whole "passport" */
		if line == "" {
			kv := re.FindAllStringSubmatch(currentLine, -1)

			check := 0
			for _, match := range kv {
				validator, ok := validators[match[1]]
				if !ok {
					continue
				}
				if validator(match[2]) {
					check++
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
