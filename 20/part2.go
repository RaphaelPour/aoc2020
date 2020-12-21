package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"regexp"
	"strconv"
)

var (
	inputFile = "input2"
)

const (
	LEFT = iota
	RIGHT
	TOP
	BOTTOM
)

/*
 * TODO:
 *
 * [ ] Add image flip
 * [ ] Add image rotate
 * [ ] MatchAny: No flip/rotate, return match
 *
 */

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

func (r Row) Flip() {
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-i-1] = r[len(r)-i-1], r[i]
	}
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
	id                           int
	data                         []Row
	left, right, top, bottom     Row
	sides                        []Row
	flippedSides                 []Row
	width, height                int
	leftN, rightN, topN, bottomN int
}

func NewImage() Image {
	img := Image{}
	img.data = make([]Row, 0)
	img.leftN = -1
	img.rightN = -1
	img.topN = -1
	img.bottomN = -1
	return img
}

func (img *Image) FlipX() {
	/* Flip row-wise */
	for i, _ := range img.data {
		img.data[i].Flip()
	}
	img.Process()
}

func (img *Image) FlipY() {
	/* Swap all rows pair wise */
	for i := 0; i < len(img.data)/2; i++ {
		inverse := len(img.data) - 1 - i
		img.data[i], img.data[inverse] = img.data[inverse], img.data[i]
	}
	img.Process()
}

func (img *Image) Transpose() {
	newData := make([]Row, img.width)
	for i := 0; i < img.width; i++ {
		newData[i] = make(Row, img.height)
	}

	/* Transpose */
	for y := 0; y < img.height; y++ {
		for x := 0; x < img.width; x++ {
			newData[x][y] = img.data[y][x]
		}
	}

	img.data = newData
	img.Process()
}

func (img *Image) RotateRight() {
	img.Transpose()
	img.FlipX()
}

func (img *Image) RotateLeft() {
	img.Transpose()
	img.FlipY()
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

func (img Image) Neighbours() int {
	count := 0
	if img.leftN > -1 {
		count++
	}
	if img.rightN > -1 {
		count++
	}
	if img.topN > -1 {
		count++
	}
	if img.bottomN > -1 {
		count++
	}
	return count
}

func (img Image) IsCorner() bool {
	return img.Neighbours() == 2
}

func (img Image) MatchBorder(other Image) int {

	if img.left.Equals(other.right) {
		return LEFT
	}

	if img.right.Equals(other.left) {
		return RIGHT
	}

	if img.top.Equals(other.bottom) {
		return TOP
	}

	if img.bottom.Equals(other.top) {
		return BOTTOM
	}

	return -1
}

func (img *Image) ProcessBorder(other Image) int {

	if img.left.Equals(other.right) {
		img.leftN = other.id
		return LEFT
	}

	if img.right.Equals(other.left) {
		img.rightN = other.id
		return RIGHT
	}

	if img.top.Equals(other.bottom) {
		img.topN = other.id
		return TOP
	}

	if img.bottom.Equals(other.top) {
		img.bottomN = other.id
		return BOTTOM
	}

	return -1
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

func Resolve(images Images, width, height int) int {

	order := make([][]int, height)
	for i := range order {
		order[i] = make([]int, width)
	}

	/* Find and store border connections */
	for i, imgA := range images {
		for j, imgB := range images {
			if i == j {
				continue
			}

			imgA.ProcessBorder(imgB)
			images[i] = imgA
			fmt.Println(images[i].Neighbours())
		}
	}

	for _, img := range images {
		fmt.Println(img.id, img.Neighbours())
	}
	return 0

	/* Find first corner */
	foundCorner := false
	for id, img := range images {
		if img.IsCorner() {
			fmt.Println("Found corner")
			order[0][0] = id
			foundCorner = true
			break
		}
	}

	if !foundCorner {
		fmt.Println("Couldn't find any corner")
		return -1
	}

	/* Arrange images */
	x, y := 1, 0
	dir := LEFT
	lastID := order[0][0]
	for {
		if y >= len(order) {
			break
		}

		if x < 0 {
			x = 0
			y++
			order[y][x] = images[lastID].bottomN
			lastID = order[y][x]
			dir = RIGHT
			continue
		} else if x >= width {
			x = width - 1
			y++
			order[y][x] = images[lastID].bottomN
			lastID = order[y][x]
			dir = LEFT
			continue
		}

		if dir == RIGHT {
			order[y][x] = images[lastID].rightN
			x++
		} else if dir == LEFT {
			order[y][x] = images[lastID].leftN
			x--
		} else {
			fmt.Println("Invalid direction", dir)
		}

		lastID = order[y][x]
	}

	for i := range order {
		fmt.Println(order[i])
	}

	return 1
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
	height := images[id].height
	width := images[id].width

	/* Arrange puzzle */
	fmt.Println("Width:", width)
	fmt.Println("Height:", height)
	fmt.Println(Resolve(images, width, height))
}
