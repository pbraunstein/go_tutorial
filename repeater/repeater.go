package repeater

func Repeater(val string, times int) string {
    repeated := ""

    for i := 0; i < times; i++ {
        repeated += val
    }

    return repeated
}
