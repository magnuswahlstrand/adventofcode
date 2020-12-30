package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSomething(t *testing.T) {
	// 41x41x41 seems like reasonable size
	const size = 21
	mid := size / 2

	inputXY := bytes.Split(testInput, []byte("\n"))
	offset := len(inputXY) / 2

	xyzWorld := initWorld(size, inputXY, mid, offset)
	printWorld(mid, mid, offset, xyzWorld)

	// Cycle 1
	prepareNext(size, xyzWorld)
	active := applyNext(size, xyzWorld)

	require.EqualValues(t, 11, active)

	printWorld(mid-1, mid, offset+1, xyzWorld)
	printWorld(mid, mid, offset+1, xyzWorld)
	printWorld(mid+1, mid, offset+1, xyzWorld)

	// Cycle 2
	prepareNext(size, xyzWorld)
	active = applyNext(size, xyzWorld)
	require.EqualValues(t, 21, active)

	cycles := 6
	for r := 0; r < cycles-2; r++ {
		prepareNext(size, xyzWorld)
		active = applyNext(size, xyzWorld)
	}
	require.EqualValues(t, 112, active)
}

func TestSomethingPart1(t *testing.T) {
	input := testInput
	const size = 21

	active := solvePart1(size, input)
	require.EqualValues(t, 112, active)
}

func TestPart2AndPrint(t *testing.T) {
	input := testInput
	const size = 21

	mid := size / 2
	inputXY := bytes.Split(input, []byte("\n"))
	offset := len(inputXY) / 2
	wxyzWorld := initWorld2(size, inputXY, mid, offset)

	printWorld2(mid, mid, mid, offset+1, wxyzWorld)

	prepareNext2(size, wxyzWorld)
	active := applyNext2(size, wxyzWorld)
	require.EqualValues(t, 29, active)

	prepareNext2(size, wxyzWorld)
	active = applyNext2(size, wxyzWorld)
	require.EqualValues(t, 60, active)
}

func TestPart2(t *testing.T) {
	input := testInput
	const size = 21

	active := solvePart2(size, input)
	require.EqualValues(t, 848, active)
}
