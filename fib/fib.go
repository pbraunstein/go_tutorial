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

func CollectFibSerial() map[int]int {
    toReturn := make(map[int]int)
    for i := 0; i <= MaxCalc; i++ {
        toReturn[i] = SlowFib(i)
    }
    return toReturn
}

func CollectFibFast() map[int]int {
    toReturn := make(map[int]int)
    toReturn[0] = 0
    toReturn[1] = 1
    for i := 2; i <= MaxCalc; i++ {
        toReturn[i] = toReturn[i - 1] + toReturn[i - 2]
    }
    return toReturn
}

func CollectFibParallel() map[int]int {
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
    fibMap := CollectFibFast()
    fmt.Println(fibMap)
}
