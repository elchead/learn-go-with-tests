package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Arabic int
		Want   string
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
	for _, c := range cases {
		t.Run("convert", func(t *testing.T) {
			res := ConvertToRoman(c.Arabic)
			assert.Equal(t, c.Want, res)
		})
	}
}
