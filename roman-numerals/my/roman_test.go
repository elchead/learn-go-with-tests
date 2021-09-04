package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"convert 1 to I", 1, "I"},
		{"convert 2 to II", 2, "II"},
		{"convert 3 to III", 3, "III"},
		{"convert 4 to IV", 4, "IV"},
		{"convert 5 to V", 5, "V"},
		{"convert 6 to VI", 6, "VI"},
		{"convert 7 to VII", 7, "VII"},
		{"convert 11", 11, "XI"},
	}
	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			res := ConvertToRoman(c.Arabic)
			assert.Equal(t, c.Want, res)
		})
	}
}
