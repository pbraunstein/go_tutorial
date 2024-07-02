package repeater

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestRepeater(t *testing.T) {
    t.Run("No repeats", func(t *testing.T) {
        repeated := Repeater("a", 0)
        expected := ""

        assert.Equal(t, expected, repeated)
    })

    t.Run("Multiple repeats", func(t *testing.T) {
        repeated := Repeater("a", 9)
        expected := "aaaaaaaaa"

        assert.Equal(t, expected, repeated)
    })

    t.Run("Multi Char String", func(t *testing.T) {
        repeated := Repeater("abc", 3)
        expected := "abcabcabc"

        assert.Equal(t, expected, repeated)
    })
}
