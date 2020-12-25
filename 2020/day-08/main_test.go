package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const testInput = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

func TestFindLoop(t *testing.T) {
	n := findLoopLine(testInput)
	require.Equal(t, 5, n)
}

func TestFindAndFixLoop(t *testing.T) {
	n := findAndFixLoop(testInput)
	require.Equal(t, 5, n)
}
