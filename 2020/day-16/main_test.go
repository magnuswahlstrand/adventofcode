package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = []byte(`class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`)

var testInput2 = []byte(`class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`)

func TestSomething(t *testing.T) {
	input := testInput

	errorRate := part1ErrorRate(input)
	require.EqualValues(t, 71, errorRate)
}

func TestSomething2(t *testing.T) {
	input := testInput

	errorRate := part1ErrorRate(input)
	require.EqualValues(t, 71, errorRate)
}

func TestRange(t *testing.T) {
	pair := rangePair{
		first:  [2]int{1, 5},
		second: [2]int{7, 8},
	}
	require.False(t, pair.valInRange(-1))
	require.False(t, pair.valInRange(6))
	require.False(t, pair.valInRange(10))

	require.True(t, pair.valInRange(1))
	require.True(t, pair.valInRange(3))
	require.True(t, pair.valInRange(5))
	require.True(t, pair.valInRange(8))
}
