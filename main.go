package main

import "fmt"

var deckSize int

func main() {
	fmt.Println(newCard())
	fmt.Println(getMeSomething())
}

func newCard() string {
	return "abcd"
}
