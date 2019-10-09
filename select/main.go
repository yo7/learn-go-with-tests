package main

import (
	"fmt"
	"select/racer"
)

func main() {
	winner := racer.Racer("https://twitter.com", "https://facebook.com")
	fmt.Println(winner)
}
