package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func printGreeting(name string) {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(log.Ldate | log.Ltime)

	// Get a greeting message and print it
	message, err := greetings.Hello(name)

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}

func printGreetings(names []string) {
  log.SetPrefix("greetings:")
	log.SetFlags(log.Ldate | log.Ltime)

  messages, err := greetings.Hellos(names)

  if err != nil {
    log.Fatal(err)
  }

  for i, message := range messages {
    fmt.Printf("Greetings for %s: %s\n", i, message)
  }
}

func main() {
  printGreeting("Bainian")
  printGreeting("BN")
  names := []string{
    "Apple",
    "John",
    "Bob",
  }
  printGreetings(names)
  printGreeting("")
}
