package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	width := 10.0
	height := 2.0
	res := Perimeter(width, height)
	assert.Equal(t, 24.0, res)
}
