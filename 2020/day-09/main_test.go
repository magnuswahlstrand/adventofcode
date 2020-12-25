package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const testInput = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func TestFindLoop(t *testing.T) {
	val := findInvalid(testInput, 5)
	require.Equal(t, 127, val)
}

func TestFindRange(t *testing.T) {
	val := findInvalid(testInput, 5)
	seq := findSequence(testInput, int64(val))
	require.Equal(t, []int64{15, 25, 47, 40}, seq)
}

func TestFindEncryptionWeakness(t *testing.T) {
	val := findInvalid(testInput, 5)
	seq := findSequence(testInput, int64(val))
	weakness := encryptionWeakness(seq)
	require.Equal(t, int64(62), weakness)
}

