package main

import (
	"bytes"
	"testing"
)

// TestGreet checks if the Greet function outputs the correct greeting message.
func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	// Pass a buffer and the name "Chris" to the Greet function.
	Greet(&buffer, "Chris")

	got := buffer.String()
	// Expect the output to be "Hello, Chris".
	want := "Hello, Chris"

	if got != want {
		// Report an error if the output doesn't match the expectation.
		t.Errorf("got %q want %q", got, want)
	}
}
