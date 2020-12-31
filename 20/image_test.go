package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRowEquals(t *testing.T) {
	row1 := Row{true, true, true, true}
	row2 := Row{true, true, true, true}
	row3 := Row{false, true, true, true}
	row4 := Row{true, true, true}

	require.True(t, row1.Equals(row2))
	require.False(t, row1.Equals(row3))
	require.False(t, row1.Equals(row4))
}

func TestString(t *testing.T) {
	row := Row{true, false, true}
	require.Equal(t, "#.#", row.String())
}

func TestAddRow(t *testing.T) {
	img := NewImage()

	require.Nil(t, img.AddRow("#.#"))
	require.Equal(
		t,
		"Row length 1 doesn't fit into images x resolution 3",
		img.AddRow("#").Error(),
	)
}

func TestBadPixel(t *testing.T) {
	img := NewImage()
	require.Equal(
		t,
		img.AddRow("?").Error(),
		"Error parsing row on index 0 with invalid pixel '?'\n",
	)
}

func TestImage2x2(t *testing.T) {

	expection := []Row{
		{false, false},
		{true, true},
	}

	img := NewImage()
	img.AddRow("..")
	img.AddRow("##")

	require.Equal(t, expection, img.data)
}

func TestImage3x3(t *testing.T) {

	expection := []Row{
		{false, false, false},
		{true, true, true},
		{false, false, false},
	}

	img := NewImage()
	img.AddRow("...")
	img.AddRow("###")
	img.AddRow("...")

	require.Equal(t, expection, img.data)
}

func TestImage3x2(t *testing.T) {

	expection := []Row{
		{false, false, false},
		{true, true, true},
	}

	img := NewImage()
	img.AddRow("...")
	img.AddRow("###")

	require.Equal(t, expection, img.data)

}

func TestRotateRight2x2(t *testing.T) {

	expection := []Row{
		{true, false},
		{true, false},
	}

	img := NewImage()
	img.AddRow("..")
	img.AddRow("##")
	img.RotateRight()

	require.Equal(t, expection, img.data)

}

func TestRotateRight3x3(t *testing.T) {

	expection := []Row{
		{false, false, true},
		{true, true, true},
		{false, false, true},
	}

	img := NewImage()
	img.AddRow("###")
	img.AddRow(".#.")
	img.AddRow(".#.")
	img.RotateRight()

	require.Equal(t, expection, img.data)

}

func TestRotateRight3x2(t *testing.T) {

	expection := []Row{
		{false, true},
		{true, true},
		{false, true},
	}

	img := NewImage()
	img.AddRow("###")
	img.AddRow(".#.")
	img.RotateRight()

	require.Equal(t, expection, img.data)

}

func TestRotateLeft2x2(t *testing.T) {

	expection := []Row{
		{false, true},
		{false, true},
	}

	img := NewImage()
	img.AddRow("..")
	img.AddRow("##")
	img.RotateLeft()

	require.Equal(t, expection, img.data)
}

func TestRotateLeft3x3(t *testing.T) {

	expection := []Row{
		{true, false, false},
		{true, true, true},
		{true, false, false},
	}

	img := NewImage()
	img.AddRow("###")
	img.AddRow(".#.")
	img.AddRow(".#.")
	img.RotateLeft()

	require.Equal(t, expection, img.data)

}

func TestRotateFullCircleRight(t *testing.T) {

	expection := []Row{
		{false, false},
		{true, true},
	}

	img := NewImage()
	img.AddRow("..")
	img.AddRow("##")
	img.RotateRight()
	img.RotateRight()
	img.RotateRight()
	img.RotateRight()

	require.Equal(t, expection, img.data)

}

func TestRotateFullCircleLeft(t *testing.T) {

	expection := []Row{
		{false, false},
		{true, true},
	}

	img := NewImage()
	img.AddRow("..")
	img.AddRow("##")
	img.RotateLeft()
	img.RotateLeft()
	img.RotateLeft()
	img.RotateLeft()

	require.Equal(t, expection, img.data)

}

func TestRotateLeftRight(t *testing.T) {

	expection := []Row{
		{false, false},
		{true, true},
	}

	img := NewImage()
	img.AddRow("..")
	img.AddRow("##")
	img.RotateLeft()
	img.RotateRight()

	require.Equal(t, expection, img.data)
}

func TestRotateLeftRightTwice(t *testing.T) {

	expection := []Row{
		{false, false},
		{true, true},
	}

	img := NewImage()
	img.AddRow("..")
	img.AddRow("##")
	img.RotateLeft()
	img.RotateLeft()
	img.RotateRight()
	img.RotateRight()

	require.Equal(t, expection, img.data)
}

func TestRotateRightLeftTwice(t *testing.T) {

	expection := []Row{
		{false, false},
		{true, true},
	}

	img := NewImage()
	img.AddRow("..")
	img.AddRow("##")
	img.RotateRight()
	img.RotateRight()
	img.RotateLeft()
	img.RotateLeft()

	require.Equal(t, expection, img.data)
}

func TestFlipX(t *testing.T) {

	expection := []Row{
		{false, false, true},
		{true, true, true},
		{false, false, true},
	}

	img := NewImage()
	img.AddRow("#..")
	img.AddRow("###")
	img.AddRow("#..")
	img.FlipX()

	require.Equal(t, expection, img.data)
}

func TestFlipY(t *testing.T) {

	expection := []Row{
		{false, true, false},
		{false, true, false},
		{true, true, true},
	}

	img := NewImage()
	img.AddRow("###")
	img.AddRow(".#.")
	img.AddRow(".#.")
	img.FlipY()

	require.Equal(t, expection, img.data)
}

func TestTranspose3x3(t *testing.T) {

	expection := []Row{
		{false, false, false},
		{true, false, false},
		{true, false, false},
	}
	img := NewImage()
	img.AddRow(".##")
	img.AddRow("...")
	img.AddRow("...")
	img.Transpose()

	require.Equal(t, expection, img.data)
}

func TestTranspose10x1(t *testing.T) {

	expection := []Row{
		{false, false, false, false, false, false, false, false, false, true},
	}
	img := NewImage()
	img.AddRow(".")
	img.AddRow(".")
	img.AddRow(".")
	img.AddRow(".")
	img.AddRow(".")
	img.AddRow(".")
	img.AddRow(".")
	img.AddRow(".")
	img.AddRow(".")
	img.AddRow("#")
	img.Transpose()

	require.Equal(t, expection, img.data)
}

func TestClone(t *testing.T) {
	img := NewImage()
	img.AddRow("..")
	img.AddRow("#.")

	other := img.Clone()

	require.Equal(t, other.data, img.data)

	other.data[0] = Row{true, true}
	require.NotEqual(t, other.data, img.data)
}

func TestDirection(t *testing.T) {
	dirs := []int{1, 2, 3, 4}
	require.Equal(t, 1, dirs[LEFT])
	require.Equal(t, 2, dirs[RIGHT])
	require.Equal(t, 3, dirs[TOP])
	require.Equal(t, 4, dirs[BOTTOM])

	require.Equal(t, 2, dirs[OppositeDirection(LEFT)])
	require.Equal(t, 1, dirs[OppositeDirection(RIGHT)])
	require.Equal(t, 4, dirs[OppositeDirection(TOP)])
	require.Equal(t, 3, dirs[OppositeDirection(BOTTOM)])
	require.Equal(t, -1, OppositeDirection(-1))
}

func TestSides(t *testing.T) {
	img := NewImage()
	img.AddRow("..")
	img.AddRow("#.")

	sides := img.Sides()
	require.True(t, sides[0].Equals(img.Left()))
	require.True(t, sides[1].Equals(img.Right()))
	require.True(t, sides[2].Equals(img.Top()))
	require.True(t, sides[3].Equals(img.Bottom()))
}
