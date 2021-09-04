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
		{"convert 3 to III", 3, "III"},
		{"convert 4 to IV", 4, "IV"},
		{"convert 5 to V", 5, "V"},
		{"convert 7 to VII", 7, "VII"},
		{"convert 9", 9, "IX"},
		{"convert 11", 11, "XI"},
		{"convert 39", 39, "XXXIX"},
	}
	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			res := ConvertToRoman(c.Arabic)
			assert.Equal(t, c.Want, res)
		})
	}
}
