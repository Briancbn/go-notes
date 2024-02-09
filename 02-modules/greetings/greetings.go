package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person
// Capital letter export function
func Hello(name string) (string, error) {
	// If no name was given, return an error
	if name == "" {
		return "", errors.New("no name given")
	}
	// Return a greeting that embeds the name in a message
	message := fmt.Sprintf(randomFormat(), name)
	// Equivalent to
	// var message string
	message = fmt.Sprintf("Hi, %v, Welcome", name)
	return message, nil
}

// Hellos returns a map that associate each of the named people
// with a greeting message
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with message
	messages := make(map[string]string)
	// Look through the received slice of names, calling
	// the hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v, Welcome!",
		"Glad to see you, %v",
		"Ni Hao %v",
	}

	return formats[rand.Intn(len(formats))]
}
