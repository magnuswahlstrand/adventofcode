package main

import (
	"testing"
)

func TestOverlap(t *testing.T) {
	fabricSize := 7
	input := `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`

	expected := 4

	parsedInput := parseElfInput(input)
	fabric := combineElfInput(parsedInput, fabricSize)
	output := countOverlap(fabric)

	if output != expected {
		t.Fatalf("\nExpected %d, got %d\n", expected, output)
	}
}

func TestNonOverlap(t *testing.T) {
	fabricSize := 7
	input := `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`

	expected := 3

	parsedInput := parseElfInput(input)
	fabric := combineElfInput(parsedInput, fabricSize)
	output := nonOverlapping(fabric, parsedInput)

	if output != expected {
		t.Fatalf("\nExpected %d, got %d\n", expected, output)
	}
}
