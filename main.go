package main

import (
	"fmt"
	"log"
	"os"

	"github.com/stephanieg0/hello-go/greetings"
	"rsc.io/quote/v4"
)

func main() {
	// generic greeting
	fmt.Println(quote.Hello())

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("Greetings error: ")
	log.SetFlags(0)

	var name_arg string

	if len(os.Args) >= 2 {
		name_arg = os.Args[1]
	}

	// custom greeting
	greeting_message, greeting_error := greetings.Hello(name_arg)

	if greeting_error != nil {
		log.Fatal(greeting_error)
	}

	fmt.Println(greeting_message)
}
