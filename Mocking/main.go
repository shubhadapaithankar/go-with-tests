package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper interface allows us to abstract the sleeping behavior.
type Sleeper interface {
	Sleep()
}

// Countdown prints a countdown from 3 to "Go!".
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, "Go!")
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

// DefaultSleeper is the default sleeping implementation.
type DefaultSleeper struct{}

// Sleep pauses execution for 1 second.
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
