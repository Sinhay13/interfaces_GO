package main

import "fmt"

// call types

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string // way to call interface, we are not limited to one type o value
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

// call with type bot
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hi There! "
}

func (sb spanishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hola ! "
}
