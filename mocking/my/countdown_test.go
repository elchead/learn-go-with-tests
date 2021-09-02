package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountdown(t *testing.T) {
	// print 3, then 1s later 2, then 1 and then Go!
	buffer := &bytes.Buffer{}
	Countdown(buffer)
	assert.Equal(t, "3", buffer.String())
}
