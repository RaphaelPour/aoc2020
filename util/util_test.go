package util

import (
	"io/ioutil"
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "input-")
	require.Nil(t, err)
	defer os.Remove(tmpFile.Name())

	input := "test"
	_, err = tmpFile.Write([]byte(input))
	require.Nil(t, err)
	require.Nil(t, tmpFile.Close())

	require.Equal(t, []string{input}, LoadString(tmpFile.Name()))
}

func TestLoadDefault(t *testing.T) {
	file, err := os.Create("input")
	require.Nil(t, err)
	defer os.Remove("input")

	input := "test"
	_, err = file.Write([]byte(input))
	require.Nil(t, err)
	require.Nil(t, file.Close())

	require.Equal(t, []string{input}, LoadDefaultString())
}

func TestLoadInt(t *testing.T) {
	file, err := os.Create("input")
	require.Nil(t, err)
	defer os.Remove("input")

	input := "1\n2"
	_, err = file.Write([]byte(input))
	require.Nil(t, err)
	require.Nil(t, file.Close())

	require.Equal(t, []int{1, 2}, LoadDefaultInt())
}

func TestProduct(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := 120

	require.Equal(t, expected, Product(input))
}

func TestSum(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := 15

	require.Equal(t, expected, Sum(input))
}

func TestReverse(t *testing.T) {
	require.Equal(t, "a", Reverse("a"))
	require.Equal(t, "ab", Reverse("ba"))
	require.Equal(t, "aba", Reverse("aba"))
	require.Equal(t, "", Reverse(""))
	require.Equal(t, "Aba", Reverse("abA"))
}

func TestMath(t *testing.T) {
	require.Equal(t, 1, Abs(1))
	require.Equal(t, 1, Abs(-1))
	require.Equal(t, 0, Abs(0))

	require.Equal(t, 0, Min(0))
	require.Equal(t, 1, Min(1))
	require.Equal(t, 1, Min(2, 1))
	require.Equal(t, -1, Min(-1, 0))

	require.Equal(t, 0, Max(0))
	require.Equal(t, 1, Max(1))
	require.Equal(t, 1, Max(-1, 1))
	require.Equal(t, 5, Max(-1, -3, -5, 3, 4, 5))

	require.True(t, InRange(0, -1, 1))
	require.True(t, InRange(1, -1, 1))
	require.True(t, InRange(-1, -1, 1))
	require.False(t, InRange(10, -1, 1))
	require.False(t, InRange(-10, -1, 1))
}

func TestMinMax(t *testing.T) {

	min, max := MinMaxInts([]int{1, 2, 3, 4})
	require.Equal(t, 1, min)
	require.Equal(t, 4, max)

	require.Equal(t, 1, MinInts([]int{1, 2, 3, 4}))
	require.Equal(t, 4, MaxInts([]int{1, 2, 3, 4}))
}

func TestSign(t *testing.T) {
	require.Equal(t, 1, Sign(1))
	require.Equal(t, -1, Sign(-1))
	require.Equal(t, 1, Sign(0))
	require.Equal(t, -1, Sign(-10000))
	require.Equal(t, 1, Sign(10000))
}

func TestPow(t *testing.T) {
	require.Equal(t, 1, Pow(1, 1))
	require.Equal(t, 2, Pow(2, 1))
	require.Equal(t, 4, Pow(2, 2))
	require.Equal(t, 9, Pow(3, 2))
	require.Equal(t, 27, Pow(3, 3))
	require.Equal(t, 1, Pow(1, 0))
}

func TestRadians(t *testing.T) {
	require.EqualValues(t, 0, Radians(0))
	require.EqualValues(t, math.Pi/2, Radians(90))
	require.EqualValues(t, math.Pi, Radians(180))
	require.EqualValues(t, 3*math.Pi/2, Radians(270))
	require.EqualValues(t, 2*math.Pi, Radians(360))
}
