package main

import "fmt"

const english = "English"
const englishHelloPrefix = "Hello, "

const spanish = "Spanish"
const spanishPrefix = "Hola, "

const french = "French"
const frenchPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Tyrone", "English"))
}
