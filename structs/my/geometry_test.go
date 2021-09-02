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

func TestArea(t *testing.T) {
	res := Area(2.0, 4.0)
	assert.Equal(t, 8.0, res)
}
