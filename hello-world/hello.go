package main

import "fmt"

const englishPrefix = "Hello, "
const spanish = "Spanish"
const french = "French"
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello returns Hello world!
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return languagePrefix(language) + name
}

func languagePrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
