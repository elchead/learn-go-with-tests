package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanNumerals(t *testing.T) {
	res, _ := ConvertToRoman(1)
	assert.Equal(t, "I", res)
}
