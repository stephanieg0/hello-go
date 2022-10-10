package main

import (
	"fmt"

	"rsc.io/quote/v4"

	"github.com/stephanieg0/hello-go/greetings"
)

func main() {
	fmt.Println(quote.Hello())

	message := greetings.Hello("Stephanie")

	fmt.Println(message)
}
