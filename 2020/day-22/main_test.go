package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = []byte(`Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`)

func TestPart1(t *testing.T) {
	sum := calcPart1(testInput)
	require.Equal(t, 306, sum)
}

var testInputPart2Simple = []byte(`Player 1:
6
4
9
8
5
2

Player 2:
7
3
10
1`)

func TestPart2(t *testing.T) {
	p1, p2 := setupPlayers(testInputPart2Simple)
	runMatchRecursive2(p1, p2)
	sum := calculateResult(p1, p2)
	require.Equal(t, 291, sum)
}

var testInputPart2Infinite = []byte(`Player 1:
43
19

Player 2:
2
29
14`)

func TestPart2StopsInfinite(t *testing.T) {
	require.Equal(t, 119, calcPart2(testInputPart2Infinite))
}
