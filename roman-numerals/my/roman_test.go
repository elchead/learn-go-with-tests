package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanNumerals(t *testing.T) {
	t.Run("convert 1 to I", func(t *testing.T) {
		res, err := ConvertToRoman(1)
		assert.Equal(t, "I", res)
		assert.NoError(t, err)
	})
	t.Run("convert 2 to II", func(t *testing.T) {
		res, err := ConvertToRoman(2)
		assert.Equal(t, "II", res)
		assert.NoError(t, err)
	})
}
