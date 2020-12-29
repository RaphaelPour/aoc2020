package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/RaphaelPour/aoc2020/util"
)

var (
	inputFile = "input"
)

/* Food:
 *
 * Corresponds to a line from the input.
 */
type Food struct {
	ingredients []string
	allergens   []string
}

func (f Food) Infer(other Food) (string, string, bool) {

	in := Intersect(f.ingredients, other.ingredients)
	if len(in) != 1 {
		return "", "", false
	}
	al := Intersect(f.allergens, other.allergens)
	if len(al) != 1 {
		return "", "", false
	}

	return in[0], al[0], true
}

func (f *Food) RemovePair(in, al string) {
	f.ingredients = util.RemoveAllStringsFromList(in, f.ingredients)
	f.allergens = util.RemoveAllStringsFromList(al, f.allergens)
}

func (f Food) String() string {
	return fmt.Sprintf(
		"%s contains %s",
		strings.Join(f.ingredients, ", "),
		strings.Join(f.allergens, ", "),
	)
}

/* Foods:
 *
 * Corresponds to the whole input. Includes a map of all found
 * allergen to food associations.
 */

type Foods struct {
	list                   []Food
	allergizingIngredients map[string]string
}

func NewFoodList() Foods {
	f := Foods{}
	f.list = make([]Food, 0)
	f.allergizingIngredients = make(map[string]string, 0)

	return f
}

func (f Foods) String() string {
	out := ""
	for i, item := range f.list {
		out += fmt.Sprintf("%2d %s\n", i+1, item)
	}
	return out
}

func (f *Foods) RemovePair(in, al string) {
	for i := range f.list {
		f.list[i].RemovePair(in, al)
	}
}

func (f Foods) IngredientCount() int {
	count := 0
	for _, food := range f.list {
		count += len(food.ingredients)
	}
	return count
}

func (f Foods) AllergenCount() int {
	count := 0
	for _, food := range f.list {
		count += len(food.allergens)
	}
	return count
}

func (f Foods) RecursiveInference(depth int, intersected []int, ins, als []string) (string, string, bool) {

	if len(ins) == 1 && len(als) == 1 {
		return ins[0], als[0], true
	} else if depth == 0 || len(ins) == 0 || len(als) == 0 {
		return "", "", false
	}

	for i := 0; i < len(f.list); i++ {
		alreadyIntersected := false
		for _, j := range intersected {
			if i == j {
				alreadyIntersected = true
			}
		}
		if alreadyIntersected {
			continue
		}
		if in, al, ok := f.RecursiveInference(
			depth-1,
			append(intersected, i),
			Intersect(ins, f.list[i].ingredients),
			Intersect(als, f.list[i].allergens),
		); ok {
			return in, al, true
		}
	}

	return "", "", false
}

func (f *Foods) MapAllergens() {

	/*
	 * All allergens have been mapped if there are
	 * no allergens left
	 */
	for f.AllergenCount() > 0 {
		inferedSomething := false

		for depth := 0; depth < len(f.list); depth++ {
			for i := 0; i < len(f.list); i++ {
				if in, al, ok := f.RecursiveInference(depth, []int{i}, f.list[i].ingredients, f.list[i].allergens); ok {
					f.RemovePair(in, al)
					f.allergizingIngredients[al] = in
					fmt.Printf("%2d: '%s' -> '%s'\n", depth, in, al)
					inferedSomething = true
					break
				}
			}
			if inferedSomething {
				break
			}
		}
		if !inferedSomething {
			fmt.Println("Couldn't find more")
			break
		}

	}
}

func (f Foods) CanoncialDangerousIngredients() string {

	keys := make([]string, 0, len(f.allergizingIngredients))
	for k, _ := range f.allergizingIngredients {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	list := make([]string, len(keys))
	for i, key := range keys {
		list[i] = f.allergizingIngredients[key]
	}

	return strings.Join(list, ",")
}

func Intersect(a, b []string) []string {
	mem := make(map[string]bool, 0)

	for _, item := range a {
		mem[item] = true
	}

	c := make([]string, 0)
	for _, item := range b {
		if _, ok := mem[item]; ok {
			c = append(c, item)
		}
	}

	return c
}

func main() {

	re := regexp.MustCompile(`^([a-z\s]+) \(contains ([a-z\s,]+)\)$`)

	foods := NewFoodList()

	for i, line := range util.LoadString(inputFile) {
		match := re.FindStringSubmatch(line)

		if len(match) != 3 {
			fmt.Printf("Error matching line %d: %s\n", i, line)
		}

		foods.list = append(foods.list, Food{
			ingredients: strings.Split(match[1], " "),
			allergens:   strings.Split(match[2], ", "),
		})
	}

	foods.MapAllergens()

	fmt.Println(foods.IngredientCount())
	fmt.Println(foods.CanoncialDangerousIngredients())
}
