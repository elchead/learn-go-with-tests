package main

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(time time.Time) Point {
	sec := time.Second()
	len := 1. //90.
	angle := secondsInRadians(sec)
	return Point{X: len * math.Sin(angle), Y: len * math.Cos(angle)}
}

func secondsInRadians(seconds int) float64 {
	return math.Pi / (30. / float64(seconds)) // arrange to avoid arithmetic error
}
