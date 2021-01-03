package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testInput = []byte(`389125467`)

func TestPart1(t *testing.T) {
	input := testInput
	circle := setupCupCircle(input)
	require.Equal(t, 3, circle.current.label)
	require.Equal(t, "389125467", circle.String())
	circle.step()

	require.Equal(t, 2, circle.current.label)
	require.Equal(t, "289154673", circle.String())
	for i := 0; i < 9; i++ {
		circle.step()
	}
	require.Equal(t, "92658374", circle.StringFrom1())

	// 100 steps total
	for i := 0; i < 90; i++ {
		circle.step()
	}
	require.Equal(t, "67384529", circle.StringFrom1())
}

func TestPart2(t *testing.T) {
	input := testInput
	circle := setupCupCirclePart2(input)
	require.Equal(t, 3, circle.current.label)

	// 10 million moves steps total
	for i := 0; i < 10*1000*1000; i++ {
		circle.step()
	}

	assert.Equal(t, 934001, circle.Part2Cup1())
	assert.Equal(t, 159792, circle.Part2Cup2())
}
