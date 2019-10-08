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

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(time.Second * 1)
}

type CountdownOperationSpy struct {
	Calls []string
}

var sleep = "sleep"
var write = "write"

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}
