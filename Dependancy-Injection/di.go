package main

import (
	"fmt"
	"io"
	"os"
)

// Greet sends a greeting to a writer.
func Greet(writer io.Writer, name string) {
	// fmt.Fprintf to format the greeting message and write it to the provided writer.
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie") // The os.Stdout as the writer to print the greeting to the standard output.
}
