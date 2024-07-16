package reflection

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

type Empty struct {}

type Dog struct{
    bark string
    dig int
    run string
}

type Cat struct {
    meow string
    zooomie float64
}

func TestReflection(t *testing.T) {
    t.Run("Empty struct", func(t *testing.T) {
        empty := Empty{}
        result := Info(empty)
        assert.Equal(t, 0, len(result))
    })

    t.Run("Dog struct", func(t *testing.T) {
        dog := Dog{"hi", 1, "zoom"}
        result := Info(dog)
        assert.Equal(t, 3, len(result))
        assert.ElementsMatch(t, []string{"bark", "dig", "run"}, result)
    })

    t.Run("Cat struct", func(t *testing.T) {
        cat := Cat{"hi", 1}
        result := Info(cat)
        assert.Equal(t, 2, len(result))
        assert.ElementsMatch(t, []string{"meow", "zooomie"}, result)
    })
}
