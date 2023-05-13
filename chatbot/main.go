package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eBot := englishBot{}
	sBot := spanishBot{}
	printGreeting(eBot)
	printGreeting(sBot)
}

func (englishBot) getGreeting() string {
	// LOGIC
	return "Hello"
}

func (spanishBot) getGreeting() string {
	// LOGIC
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
