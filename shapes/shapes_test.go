package shapes

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestShapes(t *testing.T) {
    t.Run("Double Square", func(t *testing.T) {
        s := Square{4}
        assert.Equal(t, 16.0, s.Area())
        s.Double()
        assert.Equal(t, 64.0, s.Area())
    })

    t.Run("Double Circle", func(t *testing.T) {
        c := Circle{2}
        assert.InDelta(t, 12.57, c.Area(), 0.1)
        c.Double()
        assert.InDelta(t, 50.27, c.Area(), 0.1)
    })
}
