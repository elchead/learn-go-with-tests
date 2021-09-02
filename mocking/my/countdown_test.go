package main

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const sleep = "sleep"
const write = "write"

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy := &SpyCountdownOperations{}
		Countdown(buffer, spy)
		want := `3
2
1
Go!`
		assert.Equal(t, want, buffer.String())
	},
	)
	t.Run("sleep before every print", func(t *testing.T) {
		spy := &SpyCountdownOperations{}
		Countdown(spy, spy)
		assert.Equal(t, []string{sleep, write, sleep, write, sleep, write, sleep, write}, spy.Calls)
	})
}

type spyTime struct {
	durationSlept time.Duration
}

func (s *spyTime) Sleep(duration time.Duration) {
	s.durationSlept = s.durationSlept + duration
}

func TestSleepDuration(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &spyTime{}
	timer := &ConfigurableSleeper{spyTime.Sleep, sleepTime}
	timer.Sleep()
	assert.Equal(t, sleepTime, spyTime.durationSlept)
	timer.Sleep()
	assert.Equal(t, 2*sleepTime, spyTime.durationSlept)
}
