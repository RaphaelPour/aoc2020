package util

import (
	"io/ioutil"
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

	require.Equal(t, []string{input}, Load(tmpFile.Name()))
}

func TestLoadDefault(t *testing.T) {
	file, err := os.Create("input")
	require.Nil(t, err)
	defer os.Remove("input")

	input := "test"
	_, err = file.Write([]byte(input))
	require.Nil(t, err)
	require.Nil(t, file.Close())

	require.Equal(t, []string{input}, LoadDefault())
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
