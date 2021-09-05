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
	len := 90.
	p = Point{p.X * len, p.Y * len} // scale
	p = Point{p.X, -p.Y}            // flip
	p = Point{p.X + 150, p.Y + 150} // translate
	return p
}

func secondsInRadians(seconds int) float64 {
	return math.Pi / (30. / float64(seconds)) // arrange to avoid arithmetic error
}
