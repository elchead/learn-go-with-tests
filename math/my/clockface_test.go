package main

import (
	"math"
	"testing"
	"time"

	// "github.com/quii/learn-go-with-tests/math/v1/clockface"

	"github.com/stretchr/testify/assert"
)

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) &&
		roughlyEqualFloat64(a.Y, b.Y)
}

func TestSecondHandAtMidnight(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}
	for _, c := range cases {
		t.Run("test", func(t *testing.T) {
			got := SecondHand(c.time)
			assert.Equal(t, true, roughlyEqualPoint(got, c.point))
		})
	}
	// tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
	// want := Point{X: 150, Y: 150 - 90}
	// got := SecondHand(tm)
	// assert.Equal(t, want, got)
}

// func TestSecondHandAt30Seconds(t *testing.T) {
// 	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)
// 	want := Point{X: 150, Y: 150 + 90}
// 	got := SecondHand(tm)
// 	assert.Equal(t, want, got)
// }

func TestSecondToRadians(t *testing.T) {
	got := secondsInRadians(30)
	assert.Equal(t, math.Pi, got)
	got = secondsInRadians(0)
	assert.Equal(t, 0., got)
}
