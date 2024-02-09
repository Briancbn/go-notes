package main

import "fmt"

import "rsc.io/quote"

func main() {
  fmt.Println("Hello World")
  fmt.Println("Quotes")
  fmt.Printf("Glass - %s\n", quote.Glass())
  fmt.Printf("Go - %s\n", quote.Go())
  fmt.Printf("Hello - %s\n", quote.Hello())
  fmt.Printf("Opt - %s\n", quote.Opt())
}
