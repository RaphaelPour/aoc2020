package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"github.com/RaphaelPour/prettybool"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	inputFile  = "input4"
	patchPart2 = true
)

type Alternative struct {
	terminal bool
	rules    []int
	symbol   rune
}

func (a Alternative) String() string {
	if a.terminal {
		return fmt.Sprintf(`"%c"`, a.symbol)
	}

	subRules := make([]string, len(a.rules))
	for i, rule := range a.rules {
		subRules[i] = strconv.Itoa(rule)
	}
	return strings.Join(subRules, " ")
}

type Rule []Alternative

func (r Rule) String() string {
	alternatives := make([]string, len(r))

	for i, alt := range r {
		alternatives[i] = alt.String()
	}

	return strings.Join(alternatives, " | ")
}

type Rules map[int]Rule

func (r Rules) String() string {
	keys := make([]int, len(r))

	i := 0
	for k, _ := range r {
		keys[i] = k
		i++
	}
	sort.Ints(keys)

	out := ""
	for _, k := range keys {
		out += fmt.Sprintf("%2d -> %s\n", k, r[k])
	}
	return out
}

func main() {

	rules, words, err := parseInput(util.LoadString(inputFile))
	if err != nil {
		fmt.Println(err)
		return
	}

	/* HOT PATCH */
	if patchPart2 {
		fmt.Println("HOT PATCH")
		rules[8] = Rule{
			{terminal: false, rules: []int{42}},
			{terminal: false, rules: []int{42, 8}},
		}

		rules[11] = Rule{
			{terminal: false, rules: []int{42, 31}},
			{terminal: false, rules: []int{42, 11, 31}},
		}
	}

	fmt.Printf(" == %s ==\n", os.Args[0])
	fmt.Println("-- rules --\n", rules)
	fmt.Println("-- words --\n", words)

	count := 0
	for _, word := range words {
		valid, rest, err := matchAny(word, 0, rules)

		if err != nil {
			fmt.Println(err)
			return
		}

		prettyBool := ""
		if valid && len(rest) == 0 {
			prettyBool = prettybool.GetPrettyBool(true, "check")
			count++
		} else {
			prettyBool = prettybool.GetPrettyBool(false, "check")
		}
		fmt.Println(prettyBool, word)
	}

	fmt.Println("Too low: 150")
	fmt.Println(count, "matches")
}

func parseInput(input []string) (Rules, []string, error) {

	ruleRe := regexp.MustCompile(`^(\d+): (.*)$`)
	terminalRe := regexp.MustCompile(`"([a-z])"`)

	rules := make(Rules, 0)
	words := make([]string, 0)
	parseRules := true

	for i, line := range input {
		if parseRules {
			if line == "" {
				parseRules = false
				continue
			}

			match := ruleRe.FindStringSubmatch(line)
			if len(match) < 3 {
				return nil, nil, fmt.Errorf("Error parsing line %d: %s", i, line)
			}

			ruleID, err := strconv.Atoi(match[1])
			if err != nil {
				return nil, nil, fmt.Errorf(
					"Error converting index in line  %d to number: %s\n",
					i,
					match[1],
				)
			}

			rule := make(Rule, 0)

			for j, prod := range strings.Split(match[2], " | ") {
				if prod == "" {
					return nil, nil, fmt.Errorf(
						"error parsing rule", j, ". Alternative can't be empty:",
						match[2],
					)
				}

				alternative := Alternative{}
				/* Check if terminal or rules are present */
				if match := terminalRe.FindStringSubmatch(prod); len(match) > 0 {
					/* Use rune over string since it is known that all terminal
					 * symbols have only one char.
					 */
					alternative.terminal = true
					alternative.symbol = rune(match[1][0])
					rule = append(rule, alternative)
					continue
				}

				alternative.rules = make([]int, 0)

				for k, rule := range strings.Split(prod, " ") {
					num, err := strconv.Atoi(rule)
					if err != nil {
						return nil, nil, fmt.Errorf(
							"Error converting line %d, alternative %d, field %d to number: %s\n",
							i, j, k, prod,
						)
					}

					alternative.rules = append(alternative.rules, num)
				}
				rule = append(rule, alternative)
			}

			rules[ruleID] = rule
		} else {
			words = append(words, line)
		}
	}

	return rules, words, nil
}

func matchAll(input string, rules []int, ruleSet Rules) (bool, string, error) {
	if len(input) == 0 {
		return false, "", nil
	}

	var err error
	var valid bool
	for _, rule := range rules {
		valid, input, err = matchAny(input, rule, ruleSet)
		if err != nil {
			return false, input, err
		}
		if !valid { //|| (len(rules)-ID-1) > len(input) {
			break
		}
	}
	return valid, input, nil
}

func matchAny(input string, ruleID int, rules Rules) (bool, string, error) {

	if len(input) == 0 {
		return false, "", nil
	}

	rule, ok := rules[ruleID]
	if !ok {
		return false, "", fmt.Errorf("Parsing error: unknown rule %d", ruleID)
	}

	for id, alt := range rule {
		if alt.terminal {
			return (rune(input[0]) == alt.symbol), input[1:], nil
		}
		fmt.Println(id, input)
		validAlternative, rest, err := matchAll(input, alt.rules, rules)
		if err != nil {
			return false, input, err
		}

		if validAlternative {
			return true, rest, nil
		}
	}

	return false, input, nil
}

func validate(input string, ruleID int, rules Rules) (bool, error) {
	rule, ok := rules[ruleID]
	if !ok {
		return false, fmt.Errorf("Validation error: unknown rule %d", ruleID)
	}

	for _, alt := range rule {
		if alt.terminal {
			/* Consume input if it matches the symbol from the rule.
			 * Otherwise skip the rule
			 */
			return rune(input[0]) == alt.symbol, nil
		}

		altInput := input
		matched := true
		for _, altID := range alt.rules {

			/* Each rule corresponds to at least one terminal symbol.
			 * In conclusion, if the count of rules after the production is
			 * higher than the input that's left, the rule is invalid.
			 */
			if len(alt.rules) > len(altInput) {
				matched = false
				break
			}
			valid, err := validate(input, altID, rules)
			if err != nil {
				return false, err
			}

			if valid {
				altInput = altInput[1:]
			}
		}
		if matched {
			return true, nil
		}
	}
	return false, nil
}
