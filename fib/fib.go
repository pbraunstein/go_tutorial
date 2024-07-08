package main

import "fmt"

const MaxCalc = 50

type Result struct {
    key int
    val int
}

func SlowFib(n int) int {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    return SlowFib(n - 1) + SlowFib(n - 2)
}

func CollectFib() map[int]int {
    toReturn := make(map[int]int)
    channel := make(chan Result)
    for i := 0; i <= MaxCalc; i++ {
        go func() {
            channel <- Result{i, SlowFib(i)}
        }()
    }
    for i := 0; i <= MaxCalc; i++ {
        r := <- channel
        toReturn[r.key] = r.val
    }
    return toReturn
}


func main() {
    fibMap := CollectFib()
    fmt.Println(fibMap)
}
