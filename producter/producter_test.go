package producter

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestProducter(t *testing.T) {
    t.Run("No arguments to variadic constructor", func(t *testing.T) {
        expected := make([]int, 0)
        actual := Producter()
        assert.Equal(t, expected, actual) 
    })

    t.Run("One array", func(t *testing.T) {
        expected := []int{120}
        actual := Producter([]int{1, 2, 3, 4, 5})
        assert.Equal(t, expected, actual) 
    })

    t.Run("Multiple arrays", func(t *testing.T) {
        expected := []int{120, 2700, 0}
        actual := Producter([]int{1, 2, 3, 4, 5}, []int{10, 30, 9}, []int{1, 2, 3, 4, 5, 0})
        assert.Equal(t, expected, actual) 
    })
}
