package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	// set the seed for random numbers so is different each time
	rand.Seed(time.Now().UnixNano())
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name parameter can't be blank")
	}

	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}

func randomGreeting() string {
	greetings := []string{
		"Hi %v. Welcome!",
		"Great to see you, %v!",
		"Hola %v!",
	}

	return greetings[rand.Intn(len(greetings))]
}
