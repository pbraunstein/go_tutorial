package main

import "testing"

func TestHello(t *testing.T) {
    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Philip")
        want := "Hello Philip!"

        assertCorrectMessage(t, got, want)
    })

    t.Run("Hello with empty string", func(t *testing.T) {
        got := Hello("")
        want := "Hello World!"

        assertCorrectMessage(t, got, want)
    })
}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
