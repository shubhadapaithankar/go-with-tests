// hello_test.go
package main

import "testing"

func TestHello(t *testing.T) {
    assertCorrectMessage := func(t testing.TB, got, want string) {
        t.Helper() // Tells the test suite that this method is a helper.
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }

    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })

    t.Run("empty string defaults to 'World'", func(t *testing.T) {
        got := Hello("")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })
}
