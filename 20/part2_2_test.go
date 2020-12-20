package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
	img.RotateLeft()
	img.RotateLeft()
	img.RotateRight()
	img.RotateRight()

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
