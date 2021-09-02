package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRectangle(t *testing.T) {
	rect := Rectangle{Width: 10.0, Height: 2.0}
	assert.Equal(t, 24.0, rect.Perimeter())
}

func TestArea(t *testing.T) {
	rect := Rectangle{Width: 2.0, Height: 4.0}
	assert.Equal(t, 8.0, rect.Area())
}
