package main

import (
    "fmt"
    "time"
)

func Counter() {
    for i := 1; i < 6; i++ {
        fmt.Println(i)
        if  i >= 5 {
            return
        }
        time.Sleep(time.Duration(i) * time.Second)
    }
}

func main() {
    Counter()
}
