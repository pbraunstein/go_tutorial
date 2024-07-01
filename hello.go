package main

import "fmt"

const helloPrefix = "Hello"
const emptyNameGeneric = "World"

func Hello(name string) string {
    var toPrint string
    if len(name) == 0 {
        toPrint = emptyNameGeneric
    } else {
        toPrint = name
    }
    return fmt.Sprintf("%s %s!", helloPrefix, toPrint)
}

func main() {
    fmt.Println(Hello("Phil"))
}
