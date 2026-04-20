package main

import "fmt" // standard library package for formatting and I/O

const(
	englishHelloPrefix = "Hello, " 
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	arabicHelloPrefix = "Marhaba, "

	spanish = "Spanish"
	french = "French"
	arabic = "Arabic"
)

func Hello(target string, language string) string {
	if target == "" {
		target = "World"
	}

	switch language {
	case spanish:
		return spanishHelloPrefix + target
	case french:
		return frenchHelloPrefix + target
	case arabic:
		return arabicHelloPrefix + target
	default:
		return englishHelloPrefix + target
	}
}

func main() {
	fmt.Println(Hello("world", "English")) 
	// Hello(): domain logic (pure function call)
	// fmt.Println: side effect (writes to stdout)
}