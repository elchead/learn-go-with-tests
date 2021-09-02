package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	// print 3, then 1s later 2, then 1 and then Go!
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	Countdown(buffer, spySleeper)
	want := `3
2
1
Go!`
	assert.Equal(t, want, buffer.String())
	assert.Equal(t, 4, spySleeper.Calls)
}
