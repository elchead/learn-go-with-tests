package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

func Countdown(writer io.Writer, time Sleeper) {
	for i := 3; i > 0; i-- {
		time.Sleep()
		fmt.Fprintln(writer, i)
	}
	time.Sleep()
	fmt.Fprint(writer, "Go!")
}

type RealTime struct{}

func (r *RealTime) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	timer := &RealTime{}
	Countdown(os.Stdout, timer)
}
