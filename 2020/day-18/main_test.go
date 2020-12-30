package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = []byte(``)

//5 + (8 * 3 + 9 + 3 * 4 * 3) becomes 437.
//5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4)) becomes 12240.
//((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2 becomes 13632.

func TestPart1(t *testing.T) {
	require.Equal(t, 133, evaluate("1 + 2 * (3 * 4 + (5*6)) + 7"))
	require.Equal(t, 26, evaluate("2 * 3 + (4 * 5)"))
	require.Equal(t, 54, evaluate("((2 + 4 * 9))"))
	require.Equal(t, 13632, evaluate("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"))
}

func TestPart2(t *testing.T) {
	// More tests!
	//require.Equal(t,, evaluate2(""))
	//require.Equal(t,, evaluate2(""))
	//require.Equal(t,, evaluate2(""))
	//require.Equal(t,, evaluate2(""))
	require.Equal(t, 13, evaluate2("((1 + 2 * 3) + 4)"))
	//require.Equal(t, 23004, evaluate2("2 + 7 + 6 + 3 * ((3 + 5 * 4 * 4) + 6 + 8 * 9)"))
	require.Equal(t, 750, evaluate2("2 * 5 * 2 + 5 + 2 + ((9 + 7 + 6) + 7 + 2 + 2 * 2)"))
	require.Equal(t, 20, evaluate2("5 * 4"))
	require.Equal(t, 11, evaluate2("4 + 7"))

	require.Equal(t, 200, evaluate2("(2 * 3 + 2) + (2 * 3 + 2) * (2 * 3 + 2)"))
	require.Equal(t, 46, evaluate2("2 * 3 + (4 * 5)"))
	require.Equal(t, 70, evaluate2("2 * 3 + 4 * 5"))
	require.Equal(t, 231, evaluate2("1 + 2 * 3 + 4 * 5 + 6"))
	require.Equal(t, 44, evaluate2("(4 * (5 + 6))"))
	require.Equal(t, 51, evaluate2("1 + (2 * 3) + (4 * (5 + 6))"))
	require.Equal(t, 1445, evaluate2("5 + (8 * 3 + 9 + 3 * 4 * 3)"))
	require.Equal(t, 669060, evaluate2("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"))
	require.Equal(t, 23340, evaluate2("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"))
	require.Equal(t, 1320, evaluate2("10 * 11 + (11 * 11)"))
}
