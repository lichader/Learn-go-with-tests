package main

import (
	"fmt"
)

const (
	spanish              = "Spanish"
	french               = "French"
	englishHelloPrefix   = "Hello, "
	spanishHelloPrefix   = "Hola, "
	frenchHelloPrefixabc = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	} else if language == french {
		return frenchHelloPrefixabc + name
	}

	languagePrefix := greetingPrefixadfs(language)

	return languagePrefix + name
}

func greetingPrefixadfs(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefixabc
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "english"))
}
