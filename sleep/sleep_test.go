package main

import (
    "bytes"
    "testing"
    "time"
    "github.com/stretchr/testify/assert"
)

type FakeSleeper struct {
    calls int
}

func (f *FakeSleeper) Sleep(t time.Duration) {
    f.calls += 1
}

func TestSleep(t *testing.T) {
    buffer := &bytes.Buffer{}
    fs := FakeSleeper{}

    Counter(buffer, &fs)

    assert.Equal(t, "1\n2\n3\n4\n5\n", buffer.String())
    assert.Equal(t, 4, fs.calls)
}
