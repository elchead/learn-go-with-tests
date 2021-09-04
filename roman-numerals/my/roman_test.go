package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	Arabic int
	Roman  string
}{
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{7, "VII"},
	{9, "IX"},
	{11, "XI"},
	{39, "XXXIX"},
	{56, "LVI"},
	{1984, "MCMLXXXIV"},
}

func TestAllValueOf(t *testing.T) {
	res := allLiterals.ValueOf("X")
	assert.Equal(t, 10, res)
}
func TestConvertingToRoman(t *testing.T) {
	for _, c := range cases {
		t.Run("convert", func(t *testing.T) {
			res := ConvertToRoman(c.Arabic)
			assert.Equal(t, c.Roman, res)
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, c := range cases[:2] {
		t.Run("convert", func(t *testing.T) {
			res := ConvertToArabic(c.Roman)
			assert.Equal(t, c.Arabic, res)
		})
	}
}
