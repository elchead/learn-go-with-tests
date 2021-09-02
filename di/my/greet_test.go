package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	writer := bytes.Buffer{}
	Greet(&writer, "Frank")
	assert.Equal(t, "Hello, Frank", writer.String())
}
