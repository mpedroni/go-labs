package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanish = "Spanish"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return "Hola, " + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
