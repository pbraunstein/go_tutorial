package main

import "testing"

func BenchmarkFib(b *testing.B) {
    b.Run("SerialFib", func(b *testing.B) {
        for i := 0; i< b.N; i++ {
            _ = CollectFibSerial()
        }
    })

    b.Run("ParallelFib", func(b *testing.B) {
        for i := 0; i< b.N; i++ {
            _ = CollectFibParallel()
        }
    })

    b.Run("FastFib", func(b *testing.B) {
        for i := 0; i< b.N; i++ {
            _ = CollectFibFast()
        }
    })
}
