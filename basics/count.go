package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep(time.Duration)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep(duration time.Duration) {
	time.Sleep(duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}
