package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const testInput = `abc

a
b
c

ab
ac

a
a
a
a

b`

func TestGroups(t *testing.T) {
	nGroups, _ := countGroupsAndYes(testInput)
	require.Equal(t, 5, nGroups)
}

func TestCountYes(t *testing.T) {
	nGroups, nYes := countGroupsAndYes(testInput)
	require.Equal(t, 5, nGroups)
	require.Equal(t, 11, nYes)
}

func TestCountYes2(t *testing.T) {
	nGroups, nYes := countGroupsAndYes2(testInput)
	require.Equal(t, 5, nGroups)
	require.Equal(t, 6, nYes)
}
