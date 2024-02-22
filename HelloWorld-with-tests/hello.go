// hello.go
package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

// Hello returns a greeting for the provided name and language.
// If no name is provided, it defaults to "World".
// If no language is provided, it defaults to English.
func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }
    switch language {
    case "Spanish":
        return spanishHelloPrefix + name
    default:
        return englishHelloPrefix + name
    }
}

func main() {
    fmt.Println(Hello("World", ""))
}
