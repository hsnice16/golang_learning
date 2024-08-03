package main

import (
	"fmt"

	"golang_learning/greetings"

	"rsc.io/quote"
)


func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Quote",quote.Go())

	// Get a greeting message and print it.
	message := greetings.Hello("Himanshu")
	fmt.Println(message)
}