package main

import (
    "fmt"
    "io"
    "os"
    "time"
)

func Counter(out io.Writer) {
    for i := 1; i < 6; i++ {
        fmt.Fprintf(out, "%d\n", i)
        if  i >= 5 {
            return
        }
        time.Sleep(time.Duration(i) * time.Second)
    }
}

func main() {
    Counter(os.Stdout)
}
