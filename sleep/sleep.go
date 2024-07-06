package main

import (
    "fmt"
    "time"
)

func Counter() {
    for i := 1; i < 11; i++ {
        fmt.Println(i)
        if  i >= 10 {
            return
        }
        time.Sleep(1 * time.Second)
    }
}

func main() {
    Counter()
}
