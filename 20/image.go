package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
)

const (
	FLIP_Y_MASK     = 0b0001
	FLIP_X_MASK     = 0x0010
	ROTATE_90_MASK  = 0x0100
	ROTATE_180_MASK = 0x1000

	LEFT   = 0
	RIGHT  = 1
	TOP    = 2
	BOTTOM = 3
)

func DirString(dir int) string {
	switch dir {
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	case TOP:
		return "TOP"
	case BOTTOM:
		return "BOTTOM"
	default:
		return "UNKNOWN"
	}
}

func OppositeDirection(dir int) int {
	switch dir {
	case LEFT:
		return RIGHT
	case RIGHT:
		return LEFT
	case TOP:
		return BOTTOM
	case BOTTOM:
		return TOP
	}
	return dir
}

/*
 * ROW
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

/*
 * IMAGE
 */

type Image struct {
	data []Row
}

func NewImage() Image {
	img := Image{}
	img.data = make([]Row, 0)
	return img
}

func (img Image) Clone() Image {
	other := Image{}
	other.data = make([]Row, len(img.data))

	for i := range img.data {
		other.data[i] = img.data[i].Clone()
	}
	return other
}

func (img *Image) AddRow(row string) error {

	/*
	 * Abort if the given row doesn't match the existing ones
	 * since martices need to have rows of the same count.
	 * New matrices without any rows are unaffected.
	 */
	if img.Width() > 0 && img.Width() != len(row) {
		return fmt.Errorf(
			"Row length %d doesn't fit into images x resolution %d",
			len(row),
			img.Width(),
		)
	}

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

	img.data = append(img.data, dataRow)
	return nil
}

func (img Image) Height() int {
	return len(img.data)
}

func (img Image) Width() int {
	if img.Height() == 0 {
		return 0
	}
	return len(img.data[0])
}

func (img Image) Sides() []Row {

	left := make(Row, img.Height())
	right := make(Row, img.Height())
	top := make(Row, img.Width())
	bottom := make(Row, img.Width())

	for i, row := range img.data {
		left[i] = row[0]
		right[i] = row[img.Width()-1]
	}

	copy(top, img.data[0])
	copy(bottom, img.data[img.Height()-1])

	return []Row{left, right, top, bottom}
}

func (img Image) Left() Row {
	left := make(Row, img.Height())
	for i, row := range img.data {
		left[i] = row[0]
	}
	return left
}
func (img Image) Right() Row {
	right := make(Row, img.Height())
	for i, row := range img.data {
		right[i] = row[img.Width()-1]
	}
	return right
}

func (img Image) Top() Row {
	top := make(Row, img.Width())
	copy(top, img.data[0])
	return top
}

func (img Image) Bottom() Row {
	bottom := make(Row, img.Width())
	copy(bottom, img.data[img.Height()-1])
	return bottom
}

func (img *Image) FlipX() {
	/* Flip row-wise */
	for i, _ := range img.data {
		img.data[i].Flip()
	}
}

func (img *Image) FlipY() {
	/* Flip an image along the y-axis by swapping all rows pair wise */
	for i := 0; i < len(img.data)/2; i++ {
		inverse := len(img.data) - 1 - i
		img.data[i], img.data[inverse] = img.data[inverse], img.data[i]
	}
}

func (img *Image) Transpose() {
	/* Do a matrix transponation. This can be imagined as a flip along the diagnoal */

	/* Create fresh memory there the data can be placed in */
	newData := make([]Row, img.Width())
	for i := 0; i < img.Width(); i++ {
		newData[i] = make(Row, img.Height())
	}

	/* Transpose */
	for y := 0; y < img.Height(); y++ {
		for x := 0; x < img.Width(); x++ {
			newData[x][y] = img.data[y][x]
		}
	}

	/* Assign the image to the new transposed one */
	img.data = newData
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

func (img Image) Transform(mask int) Image {
	if !util.InRange(mask, 0, 15) {
		fmt.Printf("Transformation mask is out of bounds: %d/%04b\n", mask, mask)
	}
	imgT := img.Clone()

	if mask&FLIP_Y_MASK != 0 {
		imgT.FlipY()
	}

	if mask&FLIP_X_MASK != 0 {
		imgT.FlipX()
	}

	if mask&ROTATE_90_MASK != 0 {
		imgT.RotateLeft()
	}
	if mask&ROTATE_180_MASK != 0 {
		imgT.RotateLeft()
		imgT.RotateLeft()
	}

	return imgT
}

func (img Image) TransformOtherUntilMatch(other Image) (*Image, int) {
	for mask := 0; mask < 16; mask++ {
		otherT := other.Transform(mask)

		for _, sideB := range otherT.Sides() {
			if img.Left().Equals(sideB) {
				return &otherT, LEFT
			}
			if img.Right().Equals(sideB) {
				return &otherT, RIGHT
			}
			if img.Top().Equals(sideB) {
				return &otherT, TOP
			}
			if img.Bottom().Equals(sideB) {
				return &otherT, BOTTOM
			}
		}
	}
	return nil, -1
}
