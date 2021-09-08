package main

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func secondHandPointUnitCircle(t time.Time) Point {
	sec := t.Second()
	angle := secondsInRadians(sec)
	return Point{X: math.Sin(angle), Y: math.Cos(angle)}
}

func secondHandPoint(time time.Time) Point {
	p := secondHandPointUnitCircle(time)
	return p
}

func secondsInRadians(seconds int) float64 {
	return math.Pi / (30. / float64(seconds)) // arrange to avoid arithmetic error
}

func minutesInRadians(t time.Time) float64 {
	return math.Pi/(30./float64(t.Minute())) + float64(t.Second())*math.Pi/(30.*60) // arrange to avoid arithmetic error
}
