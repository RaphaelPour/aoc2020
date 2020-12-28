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

type Side []bool

func (r Side) Equals(other Side) bool {
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

func (r Side) Flip() Side {
	newSide := make(Side, len(r))

	for i := range r {
		newSide[i] = r[len(r)-i-1]
	}
	return newSide
}

func (r Side) String() string {
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

type Img struct {
	id    int
	data  []Side
	sides []Side
}

func NewImg() Img {
	img := Img{}
	img.data = make([]Side, 0)
	return img
}

func (img Img) Height() int {
	return len(img.data)
}

func (img Img) Width() int {
	if img.Height() == 0 {
		return 0
	}
	return len(img.data[0])
}

func (img *Img) AddSide(row string) error {
	if img.Width() > 0 && img.Width() != len(row) {
		return fmt.Errorf(
			"Side length %d doesn't fit into images x resolution %d",
			len(row),
			img.Width(),
		)
	}
	dataSide := make(Side, len(row))
	for i, pixel := range row {
		if pixel != '#' && pixel != '.' {
			return fmt.Errorf(
				"Error parsing row on index %d with invalid pixel '%c'\n",
				i,
				pixel,
			)
		}

		dataSide[i] = (pixel == '#')
	}

	img.data = append(img.data, dataSide)
	return nil
}

func (img Img) MatchAny(other Img) bool {
	for _, sideA := range img.sides {
		for _, sideB := range other.sides {
			if sideA.Equals(sideB) {
				return true
			}
		}
	}
	return false
}

func (img *Img) Process() error {
	if img.Height() != img.Width() {
		return fmt.Errorf(
			"Img is no square. height=%d, width=%d",
			img.Height(),
			img.Width(),
		)
	}

	left := make(Side, img.Height())
	right := make(Side, img.Height())
	for i, row := range img.data {
		left[i] = row[0]
		right[i] = row[img.Width()-1]
	}

	img.sides = make([]Side, 8)
	img.sides[0] = left
	img.sides[1] = right
	img.sides[2] = img.data[img.Height()-1]
	img.sides[3] = img.data[0]
	img.sides[4] = left.Flip()
	img.sides[5] = right.Flip()
	img.sides[6] = img.data[img.Height()-1].Flip()
	img.sides[7] = img.data[0].Flip()

	return nil
}

func (img Img) String() string {
	out := fmt.Sprintf("= [%4d] =\n", img.id)

	for _, row := range img.data {
		out += fmt.Sprintf("%s\n", row)
	}
	return out
}

type Imgs map[int]Img

func (imgs Imgs) String() string {
	out := ""
	for _, img := range imgs {
		out += fmt.Sprintf("%s\n", img)
	}
	return out
}

func FindCorners(imgs Imgs) []int {

	corners := make([]int, 0)
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
		if sidesMatched == 2 {
			corners = append(corners, imgA.id)
		}

	}
	return corners
}

func main() {

	re := regexp.MustCompile(`^Tile (\d+):$`)
	images := make(Imgs, 0)
	img := NewImg()
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
				img = NewImg()
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
			if err := img.AddSide(line); err != nil {
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
	for _, corner := range FindCorners(images) {
		fmt.Println(corner)
		result *= corner
	}
	fmt.Println(result)
}
