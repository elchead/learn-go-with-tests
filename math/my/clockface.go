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
	len := 90.
	angle := secondsInRadians(sec)
	return Point{X: 150 + len*math.Sin(angle), Y: 150 - len*math.Cos(angle)}
}

func secondsInRadians(seconds int) float64 {
	return math.Pi / (30. / float64(seconds)) // arrange to avoid arithmetic error
}
