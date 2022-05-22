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

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())

}

func (eb englishBot) getGreeting() string {
	// Very custom logic for English
	return "Hi there"

}

func (sb spanishBot) getGreeting() string {
	// Very custom logic for Spanish
	return "Hola que tal"

}
