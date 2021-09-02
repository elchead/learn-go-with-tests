package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalk(t *testing.T) {
	x := struct {
		s string
		// i string
	}{s: "s"}

	var got []string
	fn := func(s string) { got = append(got, s) }
	walk(x, fn)

	assert.Equal(t, 1, len(got))
}
