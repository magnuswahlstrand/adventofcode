package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

var testInput = []byte(`939
7,13,x,x,59,x,31,19`)

func TestSomething(t *testing.T) {
	input := string(testInput)
	min, minBus := calcAnswer(input)
	fmt.Println(min * minBus)
}

func TestSomething2(t *testing.T) {
	input := string(testInput)
	val := findTimestamp(input)
	require.Equal(t, int64(1068781), val)
}
