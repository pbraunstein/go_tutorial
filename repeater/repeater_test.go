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

func BenchmarkRepeater(b *testing.B) {
    b.Run("O(10) repeats", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            Repeater("abc", 10)
        }
    })

    b.Run("O(10000) repeats", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            Repeater("abc", 10000)
        }
    })
}
