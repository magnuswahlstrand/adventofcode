package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = []byte(`mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`)

func TestPart1(t *testing.T) {
	input := string(testInput)
	sum, _ := calcPart1And2(input)
	require.Equal(t, 5, sum)
}

func TestPart2(t *testing.T) {
	input := string(testInput)
	_, dangerous := calcPart1And2(input)
	require.Equal(t, "mxmxvkd,sqjhc,fvjkl", dangerous)
}
