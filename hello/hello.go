package main

import (
	"fmt"
	"log"

	"rsc.io/quote"
	"vimday.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// Request a greeting message.
	//message, err := greetings.Hello("")
	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}
	//message,err:= greetings.Hello("Gladysa")
	message, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
	fmt.Println(quote.Go())
}
