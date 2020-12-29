package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var testInput = []byte(`F10
N3
F7
R90
F11`)

func TestSomething(t *testing.T) {
	input := string(testInput)

	_, pos := finalPos(input)
	d := dist(pos)

	require.Equal(t, 25, d)
	require.Equal(t, 17, pos.x)
	require.Equal(t, 8, pos.y)
}

var testInput2 = []byte(`F10
N3
F7
R90
F11
L180
F11
N5
S1`)

func TestSomething2(t *testing.T) {
	input := string(testInput2)

	_, pos := finalPos(input)
	d := dist(pos)

	require.Equal(t, 17, pos.x)
	require.Equal(t, -7, pos.y)
	require.Equal(t, 20, d)
}

func TestPart2(t *testing.T) {
	input := string(testInput)

	_, pos := finalPos2(input)

	require.Equal(t, 214, pos.x)
	require.Equal(t, 72, pos.y)

	d := dist(pos)
	require.Equal(t, 286, d)
}

func dist(pos vec) int {
	var sum int
	if pos.x < 0 {
		sum -= pos.x
	} else {
		sum += pos.x
	}

	if pos.y < 0 {
		sum -= pos.y
	} else {
		sum += pos.y
	}
	return sum
}

func TestChangeDir(t *testing.T) {
	require.Equal(t, "N", changeDir("L", "E", 90))
	require.Equal(t, "W", changeDir("L", "E", 180))
	require.Equal(t, "S", changeDir("L", "E", 270))
	require.Equal(t, "E", changeDir("L", "E", 360))

	require.Equal(t, "S", changeDir("R", "E", 90))
	require.Equal(t, "W", changeDir("R", "E", 180))
	require.Equal(t, "N", changeDir("R", "E", 270))
	require.Equal(t, "E", changeDir("R", "E", 360))
}

func TestRotateWaypoint(t *testing.T) {
	wp := vec{x: 10, y: -1}
	require.Equal(t, vec{x: -1, y: -10}, rotateWaypoint("L", wp, 90))
	require.Equal(t, vec{x: -10, y: 1}, rotateWaypoint("L", wp, 180))
	require.Equal(t, vec{x: 1, y: 10}, rotateWaypoint("L", wp, 270))
	require.Equal(t, wp, rotateWaypoint("L", wp, 360))

	require.Equal(t, vec{x: 1, y: 10}, rotateWaypoint("R", wp, 90))
	require.Equal(t, vec{x: -10, y: 1}, rotateWaypoint("R", wp, 180))
	require.Equal(t, vec{x: -1, y: -10}, rotateWaypoint("R", wp, 270))
	require.Equal(t, wp, rotateWaypoint("R", wp, 360))
}
