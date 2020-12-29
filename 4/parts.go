package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/RaphaelPour/aoc2020/util"
)

func byr(v string) bool {
	num, err := strconv.Atoi(v)
	return err == nil && util.InRange(num, 1920, 2002)
}

func iyr(v string) bool {
	num, err := strconv.Atoi(v)
	return err == nil && util.InRange(num, 2010, 2020)
}

func eyr(v string) bool {
	num, err := strconv.Atoi(v)
	return err == nil && util.InRange(num, 2020, 2030)
}

func hgt(v string) bool {
	re := regexp.MustCompile(`^(\d+)(cm|in)$`)
	match := re.FindStringSubmatch(v)
	if len(match) < 3 {
		return false
	}
	num, err := strconv.Atoi(match[1])
	if err != nil {
		return false
	}

	return (match[2] == "cm" && util.InRange(num, 150, 193)) ||
		(match[2] == "in" && util.InRange(num, 59, 76))
}

func hcl(v string) bool {
	re := regexp.MustCompile(`^#[\dabcdef]{6}$`)
	return re.MatchString(v)
}

func ecl(v string) bool {
	re := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	return re.MatchString(v)
}

func pid(v string) bool {
	re := regexp.MustCompile(`^[\d]{9}$`)
	return re.MatchString(v)
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
	for _, line := range util.LoadString("input") {
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
