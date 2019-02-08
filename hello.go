package main

import "fmt"

const englishPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}
	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
