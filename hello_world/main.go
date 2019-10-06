package main

import "fmt"

func main() {
	fmt.Println(Hello("world"))
}

// Hello is hello
func Hello(name string) string {
	return "Hello, " + name
}
