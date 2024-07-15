package racer

import (
    "fmt"
    "time"
)

func Racer(a, b, timeout int) (string, error) {
    select {
    case <-race(a):
        return "first", nil
    case <-race(b):
        return "second", nil
    case <-time.After(time.Duration(timeout) * time.Second):
        return "", fmt.Errorf("Timed out before either finished")
    }
}

func race(t int) chan struct{} {
    ch := make(chan struct{})
    go func() {
        time.Sleep(time.Duration(t) * time.Second)
        close(ch)
    }()
    return ch
}
