package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func (eb englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hi there"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (sp spanishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hola"
}
