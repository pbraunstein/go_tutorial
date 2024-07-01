package main

import "testing"

func TestHello(t *testing.T) {
    t.Run("English saying hello to people", func(t *testing.T) {
        got := Hello("Philip", "english")
        want := "Hello Philip!"

        assertCorrectMessage(t, got, want)
    })

    t.Run("English Hello with empty string", func(t *testing.T) {
        got := Hello("", "english")
        want := "Hello World!"

        assertCorrectMessage(t, got, want)
    })

    t.Run("German saying hello to people", func(t *testing.T) {
        got := Hello("Philip", "german")
        want := "Hallo Philip!"

        assertCorrectMessage(t, got, want)
    })

    t.Run("German Hello with empty string", func(t *testing.T) {
        got := Hello("", "german")
        want := "Hallo Welt!"

        assertCorrectMessage(t, got, want)
    })

    t.Run("Greek saying hello to people", func(t *testing.T) {
        got := Hello("Philip", "greek")
        want := "Geia Philip!"

        assertCorrectMessage(t, got, want)
    })

    t.Run("Greek Hello with empty string", func(t *testing.T) {
        got := Hello("", "greek")
        want := "Geia Kosmos!"

        assertCorrectMessage(t, got, want)
    })
}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
