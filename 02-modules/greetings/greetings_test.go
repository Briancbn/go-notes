package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name,
// checking for a valid return value.
func TestHelloName(t *testing.T) {
	name := "John"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) {
		t.Fatalf("Hello(%q) = %q, %v, must match for %#q, nil",
			name, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	name := ""
	_, err := Hello(name)
	if err == nil {
		t.Fatal("No error reported for empty string")
	}
}
