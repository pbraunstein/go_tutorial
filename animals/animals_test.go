package animals

import (
    "testing"
    "fmt"
    "github.com/stretchr/testify/assert"
)

func TestAnimals(t *testing.T) {
    h := Human{}
    fmt.Printf("%T", h)
    assert.Equal(t, 1, 1)
}

