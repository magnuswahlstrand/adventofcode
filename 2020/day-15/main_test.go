package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = []byte(`0,3,6`)

func TestSomething(t *testing.T) {
	answer := findNumberPart1(testInput, 10)
	require.Equal(t, 0, answer)
}

func TestSomethingMore(t *testing.T) {
	require.Equal(t, 1, findNumberPart1([]byte(`1,3,2`), 2020))
	require.Equal(t, 10, findNumberPart1([]byte(`2,1,3`), 2020))
	require.Equal(t, 27, findNumberPart1([]byte(`1,2,3`), 2020))
	require.Equal(t, 78, findNumberPart1([]byte(`2,3,1`), 2020))
	require.Equal(t, 438, findNumberPart1([]byte(`3,2,1`), 2020))
	require.Equal(t, 1836, findNumberPart1([]byte(`3,1,2`), 2020))
}
