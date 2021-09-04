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
	}
	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			res := ConvertToRoman(c.Arabic)
			assert.Equal(t, c.Want, res)
		})
	}
}
