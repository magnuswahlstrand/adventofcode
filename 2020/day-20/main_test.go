package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = []byte(`Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...`)

func TestPart1Product(t *testing.T) {
	input := testInput
	tilesMap, tiles := parseTiles(input)
	require.Equal(t, "..##.#..#.", tilesMap[2311].side(0))
	require.Equal(t, "...#.##..#", tilesMap[2311].side(1))
	require.Equal(t, "..###..###", tilesMap[2311].side(2))
	require.Equal(t, ".#####..#.", tilesMap[2311].side(3))

	product, _ := findCornerProductAndMatches(tiles)

	require.Equal(t, product, 20899048083289)
}

func TestPart2Rotate(t *testing.T) {
	ti := &tile{
		id:   123,
		grid: []string{"123", "456", "789"},
	}

	ti.rotate()
	require.Equal(t, []string{"741", "852", "963"}, ti.grid)
	ti.rotate()
	//ti.print()
	require.Equal(t, []string{"987", "654", "321"}, ti.grid)
}

func TestPart2Rotate2(t *testing.T) {
	ti := &tile{
		id:   123,
		grid: []string{"0123", "4567", "89ab", "cdef"},
	}

	ti.rotate()
	//ti.print()
	require.Equal(t, []string{"c840", "d951", "ea62", "fb73"}, ti.grid)
}

func TestPart2FindMonster(t *testing.T) {
	world2 := [][]byte{
		[]byte("                  # "),
		[]byte("#    ##    ##    ###"),
		[]byte(" #  #  #  #  #  #   "),
		[]byte("                    "),
		[]byte("                    "),
		[]byte("                    "),
		[]byte("                    "),
		[]byte("                    "),
		[]byte("                    "),
		[]byte("                    "),
		[]byte("       ###          "),
		[]byte("                    "),
		[]byte("                    "),
		[]byte("                    "),
		[]byte("                    "),
		[]byte("                    "),
		[]byte("                    "),
		[]byte("                  # "),
		[]byte("#    ##    ##    ###"),
		[]byte(" #  #  #  #  #  #   "),
	}

	h := worldHolder{
		world: world2,
	}

	wormCount := h.countWorms()

	countOs := h.countOs()

	require.Equal(t, 2, wormCount)
	require.Equal(t, 3, countOs)
}

var parsedTestWorldString = `.#.#..#.##...#.##..#####
###....#.#....#..#......
##.##.###.#.#..######...
###.#####...#.#####.#..#
##.#....#.##.####...#.##
...########.#....#####.#
....#..#...##..#.#.###..
.####...#..#.....#......
#..#.##..#..###.#.##....
#.####..#.####.#.#.###..
###.#.#...#.######.#..##
#.####....##..########.#
##..##.#...#...#.#.#.#..
...#..#..#.#.##..###.###
.#.#....#.##.#...###.##.
###.#...#..#.##.######..
.#.#.###.##.##.#..#.##..
.####.###.#...###.#..#.#
..#.#..#..#.#.#.####.###
#..####...#.#.#.###.###.
#####..#####...###....##
#.##..#..#...#..####...#
.#.###..##..##..####.##.
...###...##...#...#..###`

var parsedTestWorld = bytes.Split([]byte(parsedTestWorldString), []byte("\n"))

func TestPart2Count(t *testing.T) {
	os := countOsInWorld(parsedTestWorld)
	require.Equal(t, 273, os)
}

func TestPart2Complete(t *testing.T) {
	tilesMap, tiles := parseTiles(testInput)
	_, connectsTo := findCornerProductAndMatches(tiles)

	// Select a starting corner
	ww := layoutTiles(connectsTo, tiles, tilesMap)

	fmt.Println("from web")
	fmt.Println(parsedTestWorldString)

	fmt.Println("")
	printWorld(ww.cut())

	os := countOsInWorld(ww.cut())
	require.Equal(t, 273, os)
}

func TestPart2Flip(t *testing.T) {
	ti := &tile{
		id:   123,
		grid: []string{"123", "456", "789"},
	}

	ti.flip()
	//ti.print()
	require.Equal(t, []string{"147", "258", "369"}, ti.grid)
}

func TestPart2World(t *testing.T) {
	input := testInput
	tilesMap, tiles := parseTiles(input)
	product, connectsTo := findCornerProductAndMatches(tiles)
	require.Equal(t, product, 20899048083289)

	// Select a starting corner
	ww := layoutTiles(connectsTo, tiles, tilesMap)

	cutWorld := ww.cut()
	printWorld(cutWorld)
	require.Len(t, cutWorld, 3*10-6)
	require.Len(t, cutWorld, 3*10-6)
}
