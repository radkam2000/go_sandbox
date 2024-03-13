package main

import "fmt"

const englishHelloPrefix = "Hello, "
const englishHelloWorld = "Hello world!"

func Hello(name string) string {
	if name == "" {
		return englishHelloWorld
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello(""))
}
