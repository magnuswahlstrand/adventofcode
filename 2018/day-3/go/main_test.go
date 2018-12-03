package main

import (
	"fmt"
	"testing"
)

func TestRegex(t *testing.T) {
	tcs := []struct {
		in       string
		expected [5]int
	}{
		{"#1 @ 1,3: 4x4", [5]int{1, 1, 3, 4, 4}},
		{"#2 @ 3,1: 4x4", [5]int{2, 3, 1, 4, 4}},
		{"#3 @ 5,5: 2x2", [5]int{3, 5, 5, 2, 2}},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {

			output := parseElfData(tc.in)
			if tc.expected != output {
				t.Fatalf("\nExpected %v, got %v\n", tc.expected, output)
			}
		})
	}
}

func TestOverlap(t *testing.T) {
	fabricSize := 7
	input := `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`

	expected := 4

	fabric := combineElfInput(input, fabricSize)
	output := countOverlap(fabric)
	if output != expected {
		t.Fatalf("\nExpected %d, got %d\n", expected, output)
	}
}
