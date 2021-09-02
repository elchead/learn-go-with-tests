package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(writer io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintf(writer, "%d\n", i)
	}
	fmt.Fprint(writer, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
