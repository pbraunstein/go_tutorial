package main

import "testing"

func BenchmarkCounter(b *testing.B) {
    for i := 0; i < b.N; i++ {
        CalculateResults()
    }
}
