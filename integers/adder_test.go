package integers

import (
    "testing"
)

func TestAdder(t *testing.T) {
    t.Run("Two positives", func(t *testing.T) {
        sum := Add(2, 2)
        expected := 4

        AssertEqual(sum, expected, t)
    })

    t.Run("Two negatives", func(t *testing.T) {
        sum := Add(-2, -2)
        expected := -4

        AssertEqual(sum, expected, t)
    })

    t.Run("one positive one negative", func(t *testing.T) {
        sum := Add(-2, 2)
        expected := 0

        AssertEqual(sum, expected, t)
    })
}

func AssertEqual(got, want int, t testing.TB) {
    t.Helper()
    
    if got != want {
        t.Errorf("got %d want %d", got, want)
    }
}


