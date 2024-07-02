package integers

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestAdder(t *testing.T) {
    t.Run("Two positives", func(t *testing.T) {
        sum := Add(2, 2)
        expected := 4

        assert.Equal(t, expected, sum)
    })

    t.Run("Two negatives", func(t *testing.T) {
        sum := Add(-2, -2)
        expected := -4

        assert.Equal(t, expected, sum)
    })

    t.Run("one positive one negative", func(t *testing.T) {
        sum := Add(-2, 2)
        expected := 0

        assert.Equal(t, expected, sum)
    })
}

