package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"regexp"
	"strconv"
)

var (
	inputFile = "input"
)

type Row []bool

func (r Row) Equals(other Row) bool {
	if len(r) != len(other) {
		fmt.Printf(
			"Lengths differ: %d != %d\n",
			len(r),
			len(other),
		)
		return false
	}

	for i := range r {
		if r[i] != other[i] {
			return false
		}
	}
	return true
}

func (r Row) Flip() Row {
	newRow := make(Row, len(r))

	for i := range r {
		newRow[i] = r[len(r)-i-1]
	}
	return newRow
}

func (r Row) String() string {
	out := ""
	for _, pixel := range r {
		if pixel {
			out += "#"
		} else {
			out += "."
		}
	}

	return out
}

type Image struct {
	id                       int
	data                     []Row
	left, right, top, bottom Row
	sides                    []Row
	flippedSides             []Row
	width, height            int
}

func NewImage() Image {
	img := Image{}
	img.data = make([]Row, 0)
	return img
}

func (img *Image) AddRow(row string) error {
	if img.width > 0 && img.width != len(row) {
		return fmt.Errorf(
			"Row length %d doesn't fit into images x resolution %d",
			len(row),
			img.width,
		)
	}
	img.width = len(row)
	dataRow := make(Row, len(row))
	for i, pixel := range row {
		if pixel != '#' && pixel != '.' {
			return fmt.Errorf(
				"Error parsing row on index %d with invalid pixel '%c'\n",
				i,
				pixel,
			)
		}

		dataRow[i] = (pixel == '#')
	}

	img.data = append(img.data, dataRow)
	img.height++
	return nil
}

func (img Image) MatchAny(other Image) bool {

	for _, sideA := range img.sides {
		for _, sideB := range other.sides {
			if sideA.Equals(sideB) {
				return true
			}
		}
	}

	for _, sideA := range img.flippedSides {
		for _, sideB := range other.flippedSides {
			if sideA.Equals(sideB) {
				return true
			}
		}
	}

	for _, sideA := range img.sides {
		for _, sideB := range other.flippedSides {
			if sideA.Equals(sideB) {
				return true
			}
		}
	}

	for _, sideA := range img.flippedSides {
		for _, sideB := range other.sides {
			if sideA.Equals(sideB) {
				return true
			}
		}
	}

	return false
}

func (img *Image) Process() error {
	if img.height != img.width {
		return fmt.Errorf(
			"Image is no square. height=%d, width=%d",
			img.height,
			img.width,
		)
	}
	img.left = make(Row, img.height)
	img.right = make(Row, img.height)
	img.top = make(Row, img.width)
	img.bottom = make(Row, img.width)

	copy(img.top, img.data[0])
	copy(img.bottom, img.data[img.height-1])

	for i, row := range img.data {
		img.left[i] = row[0]
		img.right[i] = row[img.width-1]
	}

	img.sides = make([]Row, 4)
	img.sides[0] = img.left
	img.sides[1] = img.right
	img.sides[2] = img.bottom
	img.sides[3] = img.top

	img.flippedSides = make([]Row, 4)
	img.flippedSides[0] = img.left.Flip()
	img.flippedSides[1] = img.right.Flip()
	img.flippedSides[2] = img.bottom.Flip()
	img.flippedSides[3] = img.top.Flip()

	return nil
}

func (img Image) String() string {
	out := fmt.Sprintf("= [%4d] =\n", img.id)

	for _, row := range img.data {
		out += fmt.Sprintf("%s\n", row)
	}

	out += fmt.Sprintf("L: %s\n", img.left)
	out += fmt.Sprintf("R: %s\n", img.right)
	out += fmt.Sprintf("T: %s\n", img.top)
	out += fmt.Sprintf("B: %s\n", img.bottom)
	return out
}

type Images map[int]Image

func (imgs Images) String() string {
	out := ""
	for _, img := range imgs {
		out += fmt.Sprintf("%s\n", img)
	}
	return out
}

func FindEdges(imgs Images) []int {

	edges := make([]int, 0)
	for _, imgA := range imgs {
		sidesMatched := 0
		for _, imgB := range imgs {
			/* Skip image itself */
			if imgA.id == imgB.id {
				continue
			}
			if imgA.MatchAny(imgB) {
				sidesMatched++
			}
		}
		fmt.Printf("%d %d\n", sidesMatched, imgA.id)
		if sidesMatched == 2 {
			edges = append(edges, imgA.id)
		}

	}
	return edges
}

func main() {

	re := regexp.MustCompile(`^Tile (\d+):$`)
	images := make(Images, 0)
	img := NewImage()
	id := 0
	for i, line := range util.LoadString(inputFile) {
		if line == "" {
			continue
		}
		match := re.FindStringSubmatch(line)
		if len(match) == 2 {
			/* Store last buffered image if it's not the first empty one */
			if id != 0 {

				if err := img.Process(); err != nil {
					fmt.Println(err)
					return
				}
				images[id] = img
				img = NewImage()
			}
			num, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Printf(
					"Tile ID '%s' is not a number on index %d\n",
					match[1],
					i,
				)
				return
			}
			id = num
			img.id = num
		} else {
			if err := img.AddRow(line); err != nil {
				fmt.Println(err)
				return
			}
		}

	}

	/* Add last image */
	if err := img.Process(); err != nil {
		fmt.Println(err)
		return
	}
	images[id] = img

	result := 1
	for _, edge := range FindEdges(images) {
		fmt.Println(edge)
		result *= edge
	}
	fmt.Println(result)
}
