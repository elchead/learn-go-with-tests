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
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{1, 2}, 2.0},
		{Circle{1.0}, 3.14},
	}

	for _, test := range areaTests {
		assert.Equal(t, test.want, test.shape.Area())
	}
}
