package main

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

var testInput = []byte(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)

var testStep2 = []byte(`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`)

var testStep3 = []byte(`#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##`)

var testStep3Part2 = []byte(`#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#`)

var testStep3Part3 = []byte(`#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#`)

var testPart2Final = []byte(`#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`)

func TestUpdateWorld(t *testing.T) {
	world := parseInput(testInput)

	world.step()
	require.Equal(t, string(testStep2), world.String())

	// Check pre-conditions
	require.Equal(t, 2, world.adjacentCount(0, 0))
	require.Equal(t, 6, world.adjacentCount(1, 1))
	world.step()

	require.Equal(t, string(testStep3), world.String())
}

func TestUpdateWorld2(t *testing.T) {
	world := parseInput2(testInput)

	world.stepTwo()
	require.Equal(t, string(testStep2), world.String())

	// Check pre-conditions
	require.Equal(t, 3, world.adjacentCount2(0, 0))
	require.Equal(t, 7, world.adjacentCount2(1, 1))
	world.stepTwo()

	require.Equal(t, string(testStep3Part2), world.String())

	world.stepTwo()
	require.Equal(t, string(testStep3Part3), world.String())

	for world.stepTwo() {
	}
	count := strings.Count(world.String(), string(SeatOccupied))
	require.Equal(t, 26, count)
}

func TestUpdateWorld2Stable(t *testing.T) {
	world := parseInput2(testInput)

	for world.stepTwo() {
	}
	count := strings.Count(world.String(), string(SeatOccupied))
	require.Equal(t, 26, count)

	require.Equal(t, string(testPart2Final), world.String())
}

func TestPart2TC2(t *testing.T) {
	world := parseInput2([]byte(`.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.`))
	require.Equal(t, 0, world.adjacentCount2(3, 3))
}

func TestPart2TC3(t *testing.T) {
	world := parseInput2([]byte(`.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`))
	require.Equal(t, 8, world.adjacentCount2(4, 3))
}

func TestPart2TC4(t *testing.T) {
	world := parseInput2([]byte(`.............
.L.L.#.#.#.#.
.............`))
	require.Equal(t, 0, world.adjacentCount2(1, 1))
	require.Equal(t, 1, world.adjacentCount2(1, 3))
	require.Equal(t, 2, world.adjacentCount2(1, 7))
}
