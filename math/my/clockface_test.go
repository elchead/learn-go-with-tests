package main

import (
	"math"
	"testing"
	"time"

	// "github.com/quii/learn-go-with-tests/math/v1/clockface"

	"github.com/stretchr/testify/assert"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
	want := Point{X: 150, Y: 150 - 90}
	got := SecondHand(tm)
	assert.Equal(t, want, got)
}

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)
	want := Point{X: 150, Y: 150 + 90}
	got := SecondHand(tm)
	assert.Equal(t, want, got)
}

func TestSecondToRadians(t *testing.T) {
	got := secondsInRadians(30)
	assert.Equal(t, math.Pi, got)
}
