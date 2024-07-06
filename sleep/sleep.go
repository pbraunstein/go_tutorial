package main

import (
    "fmt"
    "io"
    "os"
    "time"
)

type Sleeper interface {
    Sleep(length time.Duration)
}

type RealSleeper struct {}

func (r *RealSleeper) Sleep(length time.Duration) {
    time.Sleep(length)
}

func Counter(out io.Writer, s Sleeper) {
    for i := 1; i < 6; i++ {
        fmt.Fprintf(out, "%d\n", i)
        if  i >= 5 {
            return
        }
        s.Sleep(time.Duration(i) * time.Second)
    }
}

func main() {
    rs := RealSleeper{}
    Counter(os.Stdout, &rs)
}
