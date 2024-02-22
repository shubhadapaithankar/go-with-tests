// hello.go
package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello returns a greeting for the provided name.
// If no name is provided, it defaults to "World".
func Hello(name string) string {
    if name == "" {
        name = "World"
    }
    return englishHelloPrefix + name
}

func main() {
    fmt.Println(Hello("World"))
}
