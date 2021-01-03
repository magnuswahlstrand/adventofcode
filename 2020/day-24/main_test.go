package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = []byte(`sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`)

func TestPart1(t *testing.T) {
	input := testInput
	world := createWorld(input)
	require.Equal(t, 10, count(world))
}

func TestPart2(t *testing.T) {
	input := testInput
	world := createWorld(input)

	world2, processList := prepareForProcessing(world)
	require.Len(t, world2, 10)

	prepareTiles(processList, world2)
	count, processList := updateAndCount(world2)
	require.Equal(t, 15, count)

	prepareTiles(processList, world2)
	count, processList = updateAndCount(world2)
	require.Equal(t, 12, count)

	count = runNTimes(100-2, processList, world2)
	require.Equal(t, 2208, count)
}

func TestPart1FlipsSelf(t *testing.T) {
	world := createWorld([]byte("nwwswee"))
	count := count(world)
	require.Equal(t, 1, count)
	require.Equal(t, 1, world[[2]int{0, 0}])
}

func TestPart1FlipsAdjacent(t *testing.T) {
	world := createWorld([]byte("esew"))
	count := count(world)
	require.Equal(t, 1, count)
	require.Equal(t, 1, world[[2]int{0, 1}])
}

func TestPart1Small3(t *testing.T) {
	input := []byte(`sesese
sesese
sesese`)

	world := createWorld(input)
	require.Len(t, world, 1)
	blackCount := count(world)
	require.Equal(t, 1, blackCount)
	require.Equal(t, 3, world[[2]int{1, 3}])
}

func TestPart1Small(t *testing.T) {
	input := []byte(`sw
se
e
ne
nw
w`)

	world := createWorld(input)
	require.Len(t, world, 6)
	blackCount := count(world)
	fmt.Println("count", blackCount)
	require.Equal(t, 6, blackCount)
}

func TestPart1Small2(t *testing.T) {
	input := []byte(`swswsw
swswsw
swswsw`)

	world := createWorld(input)
	require.Len(t, world, 1)
	blackCount := count(world)
	require.Equal(t, 1, blackCount)
}

func TestNextPos(t *testing.T) {
	require.Equal(t, [2]int{1, 2}, wrapPos(nextPos(1, 1, "sw")))
	require.Equal(t, [2]int{2, 2}, wrapPos(nextPos(1, 1, "se")))
	require.Equal(t, [2]int{2, 1}, wrapPos(nextPos(1, 1, "e")))
	require.Equal(t, [2]int{2, 0}, wrapPos(nextPos(1, 1, "ne")))
	require.Equal(t, [2]int{1, 0}, wrapPos(nextPos(1, 1, "nw")))
	require.Equal(t, [2]int{0, 1}, wrapPos(nextPos(1, 1, "w")))

	require.Equal(t, [2]int{1, 3}, wrapPos(nextPos(2, 2, "sw")))
	require.Equal(t, [2]int{2, 3}, wrapPos(nextPos(2, 2, "se")))
	require.Equal(t, [2]int{3, 2}, wrapPos(nextPos(2, 2, "e")))
	require.Equal(t, [2]int{2, 1}, wrapPos(nextPos(2, 2, "ne")))
	require.Equal(t, [2]int{1, 1}, wrapPos(nextPos(2, 2, "nw")))
	require.Equal(t, [2]int{1, 2}, wrapPos(nextPos(2, 2, "w")))

	require.Equal(t, [2]int{1, 2}, wrapPos(nextPos(0, 1, "se")))
}

func TestParse(t *testing.T) {
	expected := []string{"se", "se", "nw", "ne", "ne", "ne", "w", "se", "e", "sw", "w", "sw", "sw", "w", "ne", "ne", "w", "se", "w", "sw"}
	require.Equal(t, expected, parseDirections("sesenwnenenewseeswwswswwnenewsewsw"))

	require.Equal(t, []string{"se", "e", "e", "e", "w"}, parseDirections("seeeew"))
}
