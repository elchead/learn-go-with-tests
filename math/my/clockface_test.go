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

func TestHourHandOnUnitCircle(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := hourHandPoint(c.time)
			assert.Equal(t, true, roughlyEqualPoint(got, c.point))
		})
	}
}

func TestMinuteHandOnUnitCircle(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}
	for _, c := range cases {
		t.Run("test", func(t *testing.T) {
			got := minuteHandPoint(c.time)
			assert.Equal(t, true, roughlyEqualPoint(got, c.point))
		})
	}
}
func TestSecondHandOnUnitCircle(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}
	for _, c := range cases {
		t.Run("test", func(t *testing.T) {
			got := secondHandPoint(c.time)
			assert.Equal(t, true, roughlyEqualPoint(got, c.point))
		})
	}
}
func TestSecondToRadians(t *testing.T) {
	got := secondsInRadians(30)
	assert.Equal(t, math.Pi, got)
	got = secondsInRadians(0)
	assert.Equal(t, 0., got)
}

func TestMinuteToRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := minutesInRadians(c.time)
			assert.Equal(t, c.angle, got)
		})
	}

}
func TestHoursToRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := hoursInRadians(c.time)
			assert.Equal(t, true, roughlyEqualFloat64(c.angle, got))
		})
	}

}
