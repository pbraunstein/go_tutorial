package main

import "fmt"

const MaxCalc = 50

func SlowFib(n int) int {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    return SlowFib(n - 1) + SlowFib(n - 2)
}


func main() {
    fmt.Println(SlowFib(MaxCalc))
}
