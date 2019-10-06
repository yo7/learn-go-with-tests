package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func main() {
	fmt.Println(Hello("world", ""))
}

// Hello is hello
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	switch language {
	case french:
		return frenchHelloPrefix + name
	case spanish:
		return spanishHelloPrefix + name
	}

	return englishHelloPrefix + name
}
