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
	return angleToPoint(secondsInRadians(t.Second()))
}

func angleToPoint(angle float64) Point {
	return Point{X: math.Sin(angle), Y: math.Cos(angle)}
}

func secondHandPoint(time time.Time) Point {
	p := secondHandPointUnitCircle(time)
	return p
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func secondsInRadians(seconds int) float64 {
	return math.Pi / (30. / float64(seconds)) // arrange to avoid arithmetic error
}

func minutesInRadians(t time.Time) float64 {
	return math.Pi/(30./float64(t.Minute())) + secondsInRadians(t.Second())/60 // arrange to avoid arithmetic error
}

func hoursInRadians(t time.Time) float64 {
	return math.Pi/(6./float64(t.Hour()%12)) + minutesInRadians(t)/12 //secondsInRadians(t.Second())/60 // arrange to avoid arithmetic error
}
