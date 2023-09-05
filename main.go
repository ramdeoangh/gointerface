package main

import "fmt"

type bot interface {
	getGreetings(string)
	//getBotVersion() float64
	//respondToUser(user) string
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
	b.getGreetings("Ramdeo")
	fmt.Println()
}

func (englishBot) getGreetings(s string) string {
	return "Hello " + s
}

func (spanishBot) getGreetings(s string) string {
	return "Hola " + s
}
