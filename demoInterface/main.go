package main

import "fmt"

type bot interface {
	// Remember function name AND return etc must be same
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	engBot := englishBot{}
	spanBot := spanishBot{}

	printGreeting(engBot)
	printGreeting(spanBot)
}

// Print greeting of a bot
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Get englishBot greeting text
func (englishBot) getGreeting() string {
	return "hello"
}

// Get spanishBot greeting text
func (spanishBot) getGreeting() string {
	return "hola"
}
