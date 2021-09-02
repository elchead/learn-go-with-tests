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

type SleepWriter interface {
	Sleeper
	io.Writer
}

func Countdown(writer SleepWriter) {
	for i := 3; i > 0; i-- {
		writer.Sleep()
		fmt.Fprintln(writer, i)
	}
	writer.Sleep()
	fmt.Fprint(writer, "Go!")
}

type RealWriter struct{}

func (r *RealWriter) Sleep() {
	time.Sleep(1 * time.Second)
}

func (r *RealWriter) Write(p []byte) (int, error) {
	return os.Stdout.Write(p)
}

func main() {
	writer := &RealWriter{}
	Countdown(writer)
}
