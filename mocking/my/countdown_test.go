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

const sleep = "sleep"
const write = "write"

type SpyCountdownOperations struct {
	Calls  []string
	Buffer *bytes.Buffer
}

func NewSpyCountdownOperations() *SpyCountdownOperations {
	return &SpyCountdownOperations{Buffer: &bytes.Buffer{}}
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return s.Buffer.Write(p)
}

func TestCountdown(t *testing.T) {
	// print 3, then 1s later 2, then 1 and then Go!
	spy := NewSpyCountdownOperations()
	Countdown(spy)
	want := `3
2
1
Go!`
	assert.Equal(t, want, spy.Buffer.String())
	assert.Equal(t, []string{sleep, write, sleep, write, sleep, write, sleep, write}, spy.Calls)
}
