package animals

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestAnimals(t *testing.T) {
    animalsTests := []struct {
        name string
        animal Animal
        greeting string
        food string
        hasThumbs bool
    }{
        {name: "Human", animal: Human{}, greeting: "Hello there", 
        food: "Meatballs", hasThumbs: true},
        {name: "Cat", animal: Cat{}, greeting: "Meow", 
        food: "Milk", hasThumbs: false},
        {name: "Dog", animal: Dog{}, greeting: "Woof", 
        food: "Bone", hasThumbs: false},
    }

    for _, tt := range animalsTests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equal(t, tt.greeting, tt.animal.Speak())
            assert.Equal(t, tt.food, tt.animal.Eat())
            assert.Equal(t, tt.hasThumbs, tt.animal.HasThumbs())
        })
    }
}

