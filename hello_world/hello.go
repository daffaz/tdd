package main

import "fmt"

const (
	englishHelloPrefix = "Hello"
	spanishHelloPrefix = "Hola"
	franceHelloPrefix  = "Bonjour"
)

func Hello(name, language string) string {
	if len(name) <= 0 {
		name = "World"
	}

	return fmt.Sprintf("%s, %s", greetingPrefix(language), name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "es":
		prefix = spanishHelloPrefix
	case "fr":
		prefix = franceHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Wawan", ""))
}
