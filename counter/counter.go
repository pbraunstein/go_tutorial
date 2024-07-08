package main

import (
    "fmt"
    "strconv"
    "time"
)

const limit = 5

type Result struct {
    key int
    value string
}

func makeResult(key int) Result {
    time.Sleep(10 * time.Second)
    return Result{key, strconv.Itoa(key)}
}

func CalculateResults() map[int]string {
    channel := make(chan Result)
    conversion := make(map[int]string)
    for i := 1; i <= limit; i++ {
        go func () {
            channel <- makeResult(i)
        }()
    }

    for i := 0; i < limit; i++ {
        r := <- channel
        conversion[r.key] = r.value
    }
    return conversion
}

func main() {
    conversion := CalculateResults()
    fmt.Println(conversion)

}
