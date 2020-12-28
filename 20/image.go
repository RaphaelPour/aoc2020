package main

import (
	"fmt"
)

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
