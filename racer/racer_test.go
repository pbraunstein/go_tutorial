package racer

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestRacer(t *testing.T) {
    t.Run("first wins", func(t *testing.T) {
        r, e := Racer(1, 2, 5)
        assert.Equal(t, "first", r)
        assert.Nil(t, e)
    })

    t.Run("second wins", func(t *testing.T) {
        r, e := Racer(2, 1, 5)
        assert.Equal(t, "second", r)
        assert.Nil(t, e)
    })

    t.Run("timeout", func(t *testing.T) {
        r, e := Racer(2, 2, 1)
        assert.Equal(t, "", r)
        assert.Error(t, e)
    })
}
