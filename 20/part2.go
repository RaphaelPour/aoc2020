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
	DIRECTION_COUNT
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

func DirectionString(direction int) string {
	switch direction {
	case LEFT:
		return "L"
	case RIGHT:
		return "R"
	case TOP:
		return "T"
	case BOTTOM:
		return "B"
	default:
		return "?"
	}
}

type Tile struct {
	id         int
	image      Image
	neighbours [DIRECTION_COUNT]int
}

func NewTile() Tile {
	return Tile{
		neighbours: [DIRECTION_COUNT]int{-1, -1, -1, -1},
	}
}

func NewTileFromImage(img *Image) Tile {
	tile := NewTile()
	tile.image = *img
	return tile
}

func (t Tile) Neighbours() int {
	/* Counts the neighbours by checking all neighbour ids.
	 * Those contain
	 *  - positive number with the neighbours id
	 *  - -1 if the tile itself the most outer one on this side
	 */

	count := 0
	for i := 0; i < DIRECTION_COUNT; i++ {
		if t.neighbours[i] > -1 {
			count++
		}
	}
	return count
}

func (t Tile) IsCorner() bool {
	/*
	 * A corner has exactly two neighbours and is otherwise the most outer
	 * tile on the two other sides.
	 */
	return t.Neighbours() == 2
}

func (t Tile) IsEdge() bool {
	/*
	 * An edge has exactly three neighbours and is otherwise the most
	 * outer tile on one side.
	 */
	return t.Neighbours() == 3
}

func (t Tile) IsCenterPart() bool {
	/*
	 * A center part is no outer most tile on any side and has therefore
	 * four neighbours.
	 */
	return t.Neighbours() == 4
}

func (t Tile) HasValidNeighbourCount() bool {
	/*
	 * A tile must be either corner, edge or center part.
	 * Anything else should be treated as error.
	 */
	return util.InRange(t.Neighbours(), 2, 4)
}

func (t Tile) MissingSides() string {
	out := ""
	for i := 0; i < DIRECTION_COUNT; i++ {
		if t.neighbours[i] == -1 {
			out += DirectionString(i)
		}
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
