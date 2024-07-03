package animals

type Animal interface {
    Speak() string
    Eat() string
    HasThumbs() bool
}

type Human struct {}
type Cat struct {}
type Dog struct {}

func (h Human) Speak() string {
    return "Hello there"
}
func (c Cat) Speak() string {
    return "Meow"
}
func (d Dog) Speak() string {
    return "Woof"
}

func (h Human) Eat() string {
    return "Meatballs"
}
func (c Cat) Eat() string {
    return "Milk"
}
func (d Dog) Eat() string {
    return "Bone"
}

func (h Human) HasThumbs() {
    return true
}
func (c Cat) HasThumbs() {
    return false
}
func (d Dog) HasThumbs() {
    return false
}



