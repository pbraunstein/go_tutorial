package dictionary

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "os"
)

var d Dictionary

func TestSearch(t *testing.T) {
    t.Run("Key found", func(t *testing.T) {
        v, e := d.Search("geia")
        assert.Equal(t, v, "hi")
        assert.Nil(t, e)
        
    })

    t.Run("Key not found", func(t *testing.T) {
        v, e := d.Search("skylo")
        assert.Equal(t, v, "")
        assert.Error(t, e)
        assert.Equal(t, ErrNotFound, e)
    })

}

func TestAdd(t *testing.T) {
    t.Run("New value", func(t *testing.T) {
        e := d.Add("skylo", "dog")
        assert.Nil(t, e)
        v, e := d.Search("skylo")
        assert.Equal(t, v, "dog")
        assert.Nil(t, e)
    })

    t.Run("Existing value", func(t *testing.T) {
        e := d.Add("geia", "hello")
        assert.Error(t, e)
        assert.Equal(t, ErrWordExists, e)
    })
}

func setup() {
    d = Dictionary{"geia": "hi",
         "mhlo": "apple",
         "gata": "cat",
     }
}

func TestMain(m *testing.M) {
    setup()
    os.Exit(m.Run())
}
