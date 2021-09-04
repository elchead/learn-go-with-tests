package main

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	Arabic uint16
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
	res := allLiterals.ValueOf('X')
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
	for _, c := range cases {
		t.Run("convert", func(t *testing.T) {
			res := ConvertToArabic(c.Roman)
			assert.Equal(t, c.Arabic, res)
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
