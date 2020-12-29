package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestByr(t *testing.T) {
	require.True(t, byr("1920"))
	require.True(t, byr("2002"))
	require.True(t, byr("2000"))
	require.False(t, byr(""))
	require.False(t, byr("2003"))
	require.False(t, byr("200"))
	require.False(t, byr("20"))
	require.False(t, byr("2"))
	require.False(t, byr("20000"))
	require.False(t, byr("1919"))
	require.False(t, byr("a"))
}

func TestIyr(t *testing.T) {
	require.True(t, iyr("2010"))
	require.True(t, iyr("2020"))
	require.True(t, iyr("2015"))
	require.False(t, iyr("2009"))
	require.False(t, iyr("2021"))
	require.False(t, iyr("20210"))
	require.False(t, iyr("202"))
	require.False(t, iyr("20"))
	require.False(t, iyr("2"))
	require.False(t, iyr(""))
	require.False(t, iyr("a"))
	require.False(t, iyr("#1234"))
}

func TestEyr(t *testing.T) {
	require.True(t, eyr("2020"))
	require.True(t, eyr("2030"))
	require.True(t, eyr("2025"))
	require.False(t, eyr("2019"))
	require.False(t, eyr("2031"))
	require.False(t, eyr("20190"))
	require.False(t, eyr("203"))
	require.False(t, eyr("20"))
	require.False(t, eyr("2"))
	require.False(t, eyr(""))
	require.False(t, eyr("a"))
}

func TestHgt(t *testing.T) {
	require.True(t, hgt("150cm"))
	require.True(t, hgt("193cm"))
	require.True(t, hgt("170cm"))
	require.True(t, hgt("59in"))
	require.True(t, hgt("76in"))
	require.True(t, hgt("70in"))
	require.False(t, hgt("150cmm"))
	require.False(t, hgt("150ccm"))
	require.False(t, hgt("150c"))
	require.False(t, hgt("15m"))
	require.False(t, hgt("15cmm"))
	require.False(t, hgt("15ccm"))
	require.False(t, hgt("15c"))
	require.False(t, hgt("15m"))

	require.False(t, hgt("59inn"))
	require.False(t, hgt("59iin"))
	require.False(t, hgt("59i"))
	require.False(t, hgt("59n"))
	require.False(t, hgt("5inn"))
	require.False(t, hgt("5iin"))
	require.False(t, hgt("5i"))
	require.False(t, hgt("5n"))
	require.False(t, hgt("190in"))
	require.False(t, hgt("190"))

}

func TestHcl(t *testing.T) {
	require.True(t, hcl("#000000"))
	require.True(t, hcl("#abcdef"))
	require.True(t, hcl("#a1c3e4"))
	require.False(t, hcl("#00000"))
	require.False(t, hcl("#0000000"))
	require.False(t, hcl("#0000"))
	require.False(t, hcl("#000"))
	require.False(t, hcl("#00"))
	require.False(t, hcl("#0"))
	require.False(t, hcl("#"))
	require.False(t, hcl("0#"))
	require.False(t, hcl("a#"))
	require.False(t, hcl("#a"))
	require.False(t, hcl("##"))
	require.False(t, hcl(""))
	require.False(t, hcl("123abc"))
}

func TestEcl(t *testing.T) {
	require.True(t, ecl("amb"))
	require.True(t, ecl("blu"))
	require.True(t, ecl("brn"))
	require.True(t, ecl("gry"))
	require.True(t, ecl("grn"))
	require.True(t, ecl("hzl"))
	require.True(t, ecl("oth"))

	require.False(t, ecl("ammb"))
	require.False(t, ecl("ab"))
	require.False(t, ecl("a"))
	require.False(t, ecl("am"))
	require.False(t, ecl(""))
	require.False(t, ecl("bluu"))
	require.False(t, ecl("brnn"))
}

func TestPid(t *testing.T) {
	require.True(t, pid("000000000"))
	require.True(t, pid("000000001"))
	require.True(t, pid("101010101"))
	require.True(t, pid("123412341"))

	require.False(t, pid("00000000"))
	require.False(t, pid("0000000"))
	require.False(t, pid("000000"))
	require.False(t, pid("00000"))
	require.False(t, pid("0000"))
	require.False(t, pid("000"))
	require.False(t, pid("00"))
	require.False(t, pid("0"))
	require.False(t, pid(""))
	require.False(t, pid("a"))
	require.False(t, pid("a1"))
	require.False(t, pid("1a"))
}
