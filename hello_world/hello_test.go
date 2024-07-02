package hello

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
    t.Run("English saying hello to people", func(t *testing.T) {
        got := Hello("Philip", "english")
        want := "Hello Philip!"

        assert.Equal(t, want, got)
    })

    t.Run("English Hello with empty string", func(t *testing.T) {
        got := Hello("", "english")
        want := "Hello World!"

        assert.Equal(t, want, got)
    })

    t.Run("German saying hello to people", func(t *testing.T) {
        got := Hello("Philip", "german")
        want := "Hallo Philip!"

        assert.Equal(t, want, got)
    })

    t.Run("German Hello with empty string", func(t *testing.T) {
        got := Hello("", "german")
        want := "Hallo Welt!"

        assert.Equal(t, want, got)
    })

    t.Run("Greek saying hello to people", func(t *testing.T) {
        got := Hello("Philip", "greek")
        want := "Geia Philip!"

        assert.Equal(t, want, got)
    })

    t.Run("Greek Hello with empty string", func(t *testing.T) {
        got := Hello("", "greek")
        want := "Geia Kosmos!"

        assert.Equal(t, want, got)
    })

    t.Run("Unknown language saying hello to people", func(t *testing.T) {
        got := Hello("Philip", "spanish")
        want := "Hello Philip!"

        assert.Equal(t, want, got)
    })

    t.Run("Unknown language Hello with empty string", func(t *testing.T) {
        got := Hello("", "spanish")
        want := "Hello World!"

        assert.Equal(t, want, got)
    })

}

