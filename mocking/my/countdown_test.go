package main

import (
	"bytes"
	"testing"
	"time"

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
	// print 3, then 1s later 2, then 1 and then Go!
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

// ConfigurableSleeper allows to set the sleep duration
type ConfigurableSleeper struct {
	sleep    func(time.Duration)
	duration time.Duration
}

func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

func TestSleepDuration(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &spyTime{}
	spy := &SpyCountdownOperations{}
	timer := &ConfigurableSleeper{spyTime.Sleep, sleepTime}
	Countdown(spy, timer)
	assert.Equal(t, 20*time.Second, spyTime.durationSlept)
}
