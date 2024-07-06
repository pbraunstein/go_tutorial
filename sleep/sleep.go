package sleep

import (
    "fmt"
    "time"
)

func Counter() {
    for i := 1; i < 11; i++ {
        fmt.Println(1)
        time.Sleep(1)
    }
}
