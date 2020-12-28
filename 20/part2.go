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
 * ## Problems
 *
 * By finding connected border pairwise all images have two matches. It appears
 * that there are duplicate sides.
 *
 * The neighbour count doesn't seem correct. There might be some flip/rotate
 * missing.
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

func (r Row) Clone() Row {
	other := make(Row, len(r))
	copy(other, r)
	return r
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
	/* Flip an image along the y-axis by swapping all rows pair wise */
	for i := 0; i < len(img.data)/2; i++ {
		inverse := len(img.data) - 1 - i
		img.data[i], img.data[inverse] = img.data[inverse], img.data[i]
	}
	img.Process()
}

func (img *Image) Transpose() {
	/* Do a matrix transponation. This can be imagined as a flip along the diagnoal */

	/* Create fresh memory there the data can be placed in */
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

	/* Assign the image to the new transposed one */
	img.data = newData

	/* Do reprocessing in order to update all convenience data */
	img.Process()
}

func (img *Image) RotateRight() {
	/* A 90° right rotation of a matrix is a transponation+flip along the x-axis */
	img.Transpose()
	img.FlipX()
}

func (img *Image) RotateLeft() {
	/* A 90° left rotation of a matrix is a transponation+flip along the y-axis */
	img.Transpose()
	img.FlipY()
}

func (img *Image) AddRow(row string) error {

	/*
	 * Abort if the given row doesn't match the existing ones
	 * since martices need to have rows of the same count.
	 * New matrices without any rows are unaffected.
	 */
	if img.width > 0 && img.width != len(row) {
		return fmt.Errorf(
			"Row length %d doesn't fit into images x resolution %d",
			len(row),
			img.width,
		)
	}

	/* Set the length of the image. Only necessary for the first row. For any
	* further row the width will be the same. */
	img.width = len(row)

	/* Transform the string to a bool array to lower memory footprint (since
	* pixels can be only black or white. */
	dataRow := make(Row, len(row))
	for i, pixel := range row {

		/* Abort if the given row has an unexpected pixel. */
		if pixel != '#' && pixel != '.' {
			return fmt.Errorf(
				"Error parsing row on index %d with invalid pixel '%c'\n",
				i,
				pixel,
			)
		}

		dataRow[i] = (pixel == '#')
	}

	/* Append the current row and correct the height */
	img.data = append(img.data, dataRow)
	img.height++
	return nil
}

func (img Image) Neighbours() int {
	/* Counts the neighbours by checking all neighbour ids.
	 * Those contain
	 *  - positive number with the neighbours id
	 *  - -1 if the tile itself the most outer one on this side
	 */
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
	/*
	 * A corner has exactly two neighbours and is otherwise the most outer
	 * tile on the two other sides.
	 */
	return img.Neighbours() == 2
}

func (img Image) IsEdge() bool {
	/*
	 * An edge has exactly three neighbours and is otherwise the most
	 * outer tile on one side.
	 */
	return img.Neighbours() == 3
}

func (img Image) IsCenterPart() bool {
	/*
	 * A center part is no outer most tile on any side and has therefore
	 * four neighbours.
	 */
	return img.Neighbours() == 4
}

func (img Image) HasValidNeighbourCount() bool {
	/*
	 * A tile must be either corner, edge or center part.
	 * Anything else should be treated as error.
	 */
	return util.InRange(img.Neighbours(), 2, 4)
}

func (img Image) MissingSides() string {
	out := ""
	if img.leftN == -1 {
		out += "L "
	}
	if img.rightN == -1 {
		out += "R "
	}
	if img.bottomN == -1 {
		out += "B "
	}
	if img.topN == -1 {
		out += "T "
	}
	return out
}

func (img Image) MatchBorder(other Image) int {
	/*
	 * Check if two images with the current orientation would
	 * fit together and return the side on which this is true.
	 *
	 * This early-return method works because two images can
	 * only be neighbours on one side simlutaneously.
	 */
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
	if len(img.sides) != len(other.sides) {
		return -2
	}
	matches := 0
	for i := 0; i < len(img.sides); i++ {
		if img.left.Equals(other.sides[i]) {
			img.leftN = other.id
			matches++
			//			return LEFT
		}

		if img.right.Equals(other.sides[i]) {
			img.rightN = other.id
			matches++
			//			return RIGHT
		}

		if img.top.Equals(other.sides[i]) {
			img.topN = other.id
			matches++
			//			return TOP
		}

		if img.bottom.Equals(other.sides[i]) {
			img.bottomN = other.id
			matches++
			///			return BOTTOM
		}

	}

	if matches > 1 {
		fmt.Printf("%d with %d matched %d times\n", img.id, other.id, matches)
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

	img.sides = make([]Row, 8)
	img.sides[0] = img.left
	img.sides[1] = img.right
	img.sides[2] = img.bottom
	img.sides[3] = img.top

	/* Store flipped sides */
	leftFlipped := img.left.Clone()
	leftFlipped.Flip()
	img.sides[4] = leftFlipped

	rightFlipped := img.right.Clone()
	rightFlipped.Flip()
	img.sides[5] = rightFlipped

	bottomFlipped := img.bottom.Clone()
	bottomFlipped.Flip()
	img.sides[6] = bottomFlipped

	topFlipped := img.top.Clone()
	topFlipped.Flip()
	img.sides[7] = topFlipped

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
			ret := imgA.ProcessBorder(imgB)
			if ret == -2 {
				fmt.Printf("Error processing %d with %d\n", imgA.id, imgB.id)
				return -1
			}
			images[i] = imgA
		}
	}

	fmt.Println("ID   |Neighbours | Missing ")
	fmt.Println("-----+-----------+----------")
	for _, img := range images {
		fmt.Printf("%4d | %9d | %s\n", img.id, img.Neighbours(), img.MissingSides())
	}
	return 0

	/* Find first corner */
	foundCorner := false
	for id, img := range images {
		if img.IsCorner() {
			fmt.Printf("Found corner %d\n", id)
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
