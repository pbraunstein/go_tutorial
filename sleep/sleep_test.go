package main

import (
    "bytes"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSleep(t *testing.T) {
    buffer := &bytes.Buffer{}

    Counter(buffer)

    assert.Equal(t, "1\n2\n3\n4\n5\n", buffer.String())
}
