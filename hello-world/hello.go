package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	} else if language == french {
		return frenchHelloPrefix + name
	}
	languagePrefix := englishHelloPrefix

	switch language {
	case spanish:
		languagePrefix = spanishHelloPrefix
	case french:
		languagePrefix = frenchHelloPrefix
	}

	return languagePrefix + name
}

func greetingPrefix(language string) (prefix string) {
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

func main() {
	fmt.Println(Hello("world", "english"))
}
