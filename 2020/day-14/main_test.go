package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = []byte(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`)

var testInput2 = []byte(`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`)

func TestSomething(t *testing.T) {
	input := testInput

	sum := calculateSumPart1(input)
	require.EqualValues(t, 165, sum)
}

func TestGetAddresses(t *testing.T) {
	addresses := getAddresses(42, "000000000000000000000000000000X1001X")

	require.Contains(t, addresses, uint64(26))
	require.Contains(t, addresses, uint64(27))
	require.Contains(t, addresses, uint64(58))
	require.Contains(t, addresses, uint64(59))
}

func TestGetAddresses2(t *testing.T) {
	addresses := getAddresses(26, "00000000000000000000000000000000X0XX")

	require.Contains(t, addresses, uint64(16))
	require.Contains(t, addresses, uint64(17))
	require.Contains(t, addresses, uint64(18))
	require.Contains(t, addresses, uint64(19))
	require.Contains(t, addresses, uint64(24))
	require.Contains(t, addresses, uint64(25))
	require.Contains(t, addresses, uint64(26))
	require.Contains(t, addresses, uint64(27))
}

func TestSomethingPart2(t *testing.T) {
	input := testInput2

	sum := calculateSumPart2(input)
	require.EqualValues(t, uint64(208), sum)
}
