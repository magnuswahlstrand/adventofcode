package main

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

const testInput = `16
10
15
5
1
11
7
19
6
12
4`

const testInput2 = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

func TestFindDiffs(t *testing.T) {
	lines := strings.Split(testInput, "\n")

	n1Diff, n3Diff := diffs(lines)
	require.Equal(t, 7, n1Diff)
	require.Equal(t, 5, n3Diff)
}

func TestFindDiffs2(t *testing.T) {
	lines := strings.Split(testInput2, "\n")

	n1Diff, n3Diff := diffs(lines)
	require.Equal(t, 22, n1Diff)
	require.Equal(t, 10, n3Diff)
}

func TestFindPaths(t *testing.T) {
	lines := strings.Split(testInput, "\n")
	branchesCount, branches := branches(lines)
	require.Equal(t, []int64{1,1, 3, 2, 1, 1, 2, 1, 1, 1, 1, 1}, branchesCount)
	require.Equal(t, [][]int{
		{1},
		{2},
		{3, 4, 5},
		{4, 5},
		{5},
		{6},
		{7, 8},
		{8},
		{9},
		{10},
		{11},
		{12},
	}, branches)

	paths := pathsForward(branches)
	require.Equal(t, 8,paths[0])
}

func TestFindPaths2(t *testing.T) {
	lines := strings.Split(testInput2, "\n")
	_, branches := branches(lines)
	paths := pathsForward(branches)
	require.Equal(t, 19208, paths[0])
}
