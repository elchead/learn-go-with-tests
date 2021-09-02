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

// ConfigurableSleeper allows to set the sleep duration
type ConfigurableSleeper struct {
	sleep    func(time.Duration)
	duration time.Duration
}

func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

func main() {
	timer := &ConfigurableSleeper{time.Sleep, 2 * time.Second} //&RealTime{}
	Countdown(os.Stdout, timer)
}
