package main

import "fmt"

const (
	spanishHelloPrefix = "Hola,"
	frenchHelloPrefix  = "Bonjour,"
	englishHelloPrefix = "Hello,"

	spanish = "Spanish"
	french  = "French"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s %s", GreetingPrefix(language), name)
}

func GreetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
