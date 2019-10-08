package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Countdown(os.Stdout)
}

const coundDownStart = 3
const finalWord = "Go!"

func Countdown(w io.Writer) {
	for i := coundDownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
	}
	fmt.Fprintf(w, finalWord)
}
