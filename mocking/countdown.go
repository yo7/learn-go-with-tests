package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	s := DefaultSleeper{}
	Countdown(os.Stdout, &s)
}

const coundDownStart = 3
const finalWord = "Go!"

func Countdown(w io.Writer, s Sleeper) {
	for i := coundDownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}
	s.Sleep()
	fmt.Fprintf(w, finalWord)
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(time.Second * 1)
}
