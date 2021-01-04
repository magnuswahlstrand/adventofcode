package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return i
}

func main() {
	t := time.Now()
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	part1(input)
	part2(input)

	log.Println("success in", time.Since(t))
}

func part1(input []byte) {
	_, tiles := parseTiles(input)
	product, _ := findCornerProductAndMatches(tiles)
	fmt.Println("answer part 1:", product)
}

// Too low: 505
// Too high: 2461
func part2(input []byte) {
	tilesMap, tiles := parseTiles(input)
	_, connectsTo := findCornerProductAndMatches(tiles)

	// Select a starting corner
	ww := layoutTiles(connectsTo, tiles, tilesMap)
	fmt.Println(len(ww.world), len(ww.world[0]))
	ww.print()

	fmt.Println("answer part 2:", countOsInWorld(ww.cut()))
}

type tile struct {
	id    int
	grid  []string
	sides []string
}

func (t *tile) side(i int) string {
	return t.sides[i]
}

func (t *tile) print() {
	fmt.Println(strings.Join(t.grid, "\n"))
}

func (t *tile) rotate() {
	newGrid := make([]string, len(t.grid))
	for y := 0; y < len(t.grid); y++ {
		var b strings.Builder
		for x := 0; x < len(t.grid); x++ {
			b.WriteByte(t.grid[len(t.grid)-1-x][y])
		}
		newGrid[y] = b.String()
	}
	t.grid = newGrid
	t.updateSides()
}

func (t *tile) flip() {
	newGrid := make([]string, len(t.grid))
	for y := 0; y < len(t.grid); y++ {
		var b strings.Builder
		for x := 0; x < len(t.grid); x++ {
			b.WriteByte(t.grid[x][y])
		}
		newGrid[y] = b.String()
	}
	t.grid = newGrid
	t.updateSides()
}

type worldHolder struct {
	world [][]byte
}

func (w *worldHolder) flip() {
	newGrid := make([][]byte, len(w.world))
	for y := 0; y < len(w.world); y++ {
		newGrid[y] = make([]byte, len(w.world[0]))

		for x := 0; x < len(w.world[0]); x++ {
			newGrid[y][x] = w.world[x][y]
		}
	}

	w.world = newGrid
}

func (w *worldHolder) rotate() {
	newGrid := make([][]byte, len(w.world))
	for y := 0; y < len(w.world); y++ {
		newGrid[y] = make([]byte, len(w.world[0]))

		for x := 0; x < len(w.world[0]); x++ {
			newGrid[y][x] = w.world[len(w.world)-1-x][y]
		}
	}

	w.world = newGrid
}

func (t *tile) updateSides() {
	s0, s1, s2, s3 := calcSides(t.grid)
	t.sides = []string{s0, s1, s2, s3}
}

func Reverse(in string) string {
	var sb strings.Builder
	runes := []rune(in)
	for i := len(runes) - 1; 0 <= i; i-- {
		sb.WriteRune(runes[i])
	}
	return sb.String()
}

type match struct {
	tile1Side, tile2Side int
	reversed             bool
}

func findCornerProductAndMatches(tiles []*tile) (int, map[int]map[int]*match) {
	connectsTo := map[int]map[int]*match{}

	cornerProduct := 1
	for i, t1 := range tiles {
		connectsTo[t1.id] = map[int]*match{}
		for t1s := 0; t1s < 4; t1s++ {

		innerloop:
			for j, t2 := range tiles {
				if i == j {
					continue
				}

				for t2s := 0; t2s < 4; t2s++ {
					straight := t1.side(t1s) == t2.side(t2s)
					reversed := t1.side(t1s) == Reverse(t2.side(t2s))
					if straight || reversed {
						connectsTo[t1.id][t2.id] = &match{
							tile1Side: t1s,
							tile2Side: t2s,
							reversed:  reversed,
						}
						break innerloop
					}
				}
			}
		}
		if len(connectsTo[t1.id]) == 2 {
			cornerProduct *= t1.id
		}
	}
	return cornerProduct, connectsTo
}

func parseTiles(input []byte) (map[int]*tile, []*tile) {
	tilesMap := map[int]*tile{}
	tiles := []*tile{}
	for _, t := range strings.Split(string(input), "\n\n") {
		rows := strings.Split(t, "\n")

		// 1. Parse id
		var tileID int
		_, err := fmt.Sscanf(rows[0], "Tile %d:", &tileID)
		if err != nil {
			log.Fatal("failed to parse title ID", err)
		}

		grid := rows[1:]
		s0, s2, s1, s3 := calcSides(grid)
		t := &tile{
			id:   tileID,
			grid: grid,
			sides: []string{
				s0,
				s1,
				s2,
				s3,
			},
		}
		tilesMap[tileID] = t
		tiles = append(tiles, t)
	}
	return tilesMap, tiles
}

func calcSides(grid []string) (string, string, string, string) {
	size := len(grid)
	s0 := grid[0]
	s2 := grid[size-1]
	var s1, s3 string
	for i := 0; i < size; i++ {
		s1 += string(grid[i][size-1])
		s3 += string(grid[i][0])
	}
	return s0, s2, s1, s3
}

var worm = bytes.Split(
	[]byte(`                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `), []byte("\n"))

func (h *worldHolder) countWorms() int {
	var wormCount int
	fmt.Println(len(h.world) - len(worm) + 1)
	fmt.Println(len(h.world[0]) - len(worm[0]) + 1)
	for y := 0; y < len(h.world)-len(worm)+1; y++ {

	outer:
		for x := 0; x < len(h.world[0])-len(worm[0])+1; x++ {

			for wy := 0; wy < len(worm); wy++ {
				for wx := 0; wx < len(worm[0]); wx++ {
					if worm[wy][wx] != '#' {
						// Ignore
						continue
					}

					if h.world[y+wy][x+wx] != '#' && h.world[y+wy][x+wx] != 'O' {
						continue outer
					}
				}
			}
			fmt.Println("search successful, fill in with Os")
			wormCount++
			for wy := 0; wy < len(worm); wy++ {
				for wx := 0; wx < len(worm[0]); wx++ {
					if worm[wy][wx] != '#' {
						// Ignore
						continue
					}
					h.world[y+wy][x+wx] = 'O'
				}
			}
		}
	}
	return wormCount
}

func (h *worldHolder) countOs() int {
	var countOs int
	for y := 0; y < len(h.world); y++ {
		for x := 0; x < len(h.world[0]); x++ {
			if h.world[y][x] == '#' {
				countOs++
			}
		}
	}
	return countOs
}

func countOsInWorld(world [][]byte) int {
	h := worldHolder{world}

	for i := 0; i < 4; i++ {
		h.countWorms()
		h.rotate()
	}
	h.flip()
	for i := 0; i < 4; i++ {
		h.countWorms()
		h.rotate()
	}
	os := h.countOs()
	printWorld(h.world)
	fmt.Println(len(h.world), len(h.world[0]))
	return os
}

type worldWrapper struct {
	world        [][]byte
	tileSize     int
	totalSize    int
	used         map[int]bool
	tilesPerSide int
}

func layoutTiles(connectsTo map[int]map[int]*match, tiles []*tile, tilesMap map[int]*tile) worldWrapper {
	startID := selectTopLeftCorner(connectsTo)
	tilesPerSide := int(math.Sqrt(float64(len(connectsTo))))
	tileSize := len(tiles[0].grid)

	totalSize := tileSize * tilesPerSide
	ww := worldWrapper{
		world:        make([][]byte, totalSize),
		tileSize:     tileSize,
		tilesPerSide: tilesPerSide,
		totalSize:    totalSize,
		used:         map[int]bool{},
	}

	for y := 0; y < ww.totalSize; y++ {
		ww.world[y] = bytes.Repeat([]byte("-"), ww.totalSize)
	}

	startTile := tilesMap[startID]
	ww.insertAt(0, 0, startTile)

	previousTile := startTile
	firstFromPreviousRow := startTile
	for dy := 0; dy < tilesPerSide; dy++ {
		for dx := 0; dx < tilesPerSide; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}

			var matchingTile *tile
			var col string
			var matchDirection int
			var connectingTiles map[int]*match

			// Hideous code :-) Quick fix to proceed with day 25
			if dx == 0 {
				matchDirection = 0

				// Get bottom row
				var b strings.Builder
				for x := 0; x < ww.tileSize; x++ {
					b.WriteByte(ww.world[dy*ww.tileSize-1][x+dx*ww.tileSize])
				}
				col = b.String()

				//ww.print()
				connectingTiles = connectsTo[firstFromPreviousRow.id]

			} else {
				// Get rightmost column
				var b strings.Builder
				for y := 0; y < ww.tileSize; y++ {
					b.WriteByte(ww.world[y+dy*ww.tileSize][dx*ww.tileSize-1])
				}
				col = b.String()

				matchDirection = 3
				connectingTiles = connectsTo[previousTile.id]
			}

		findNeighbor:
			for tID := range connectingTiles {
				if ww.used[tID] {
					continue
				}

				// Check if right side matches left side
				t := tilesMap[tID]

				// Rotate 4 times
				for i := 0; i < 4; i++ {
					if t.side(matchDirection) == col {
						matchingTile = t
						break findNeighbor
					}
					t.rotate()
				}

				// Flip
				t.flip()

				// Rotate 4 times
				for i := 0; i < 4; i++ {
					if t.side(matchDirection) == col {
						matchingTile = t
						break findNeighbor
					}
					t.rotate()
				}
			}

			if matchingTile == nil {
				break
				log.Fatal("no matching tile found")
			}

			ww.insertAt(dx, dy, matchingTile)
			previousTile = matchingTile

			if dx == 0 {
				firstFromPreviousRow = matchingTile
			}
		}
	}
	return ww
}

func (ww *worldWrapper) print() string {
	var b strings.Builder
	for y := 0; y < ww.totalSize; y++ {
		b.Write(ww.world[y])
		b.WriteString("\n")
	}
	fmt.Println(b.String())
	return b.String()
}

func (ww *worldWrapper) insertAt(wx, wy int, t *tile) {
	ww.used[t.id] = true
	for y := 0; y < ww.tileSize; y++ {
		for x := 0; x < ww.tileSize; x++ {
			ww.world[y+wy*ww.tileSize][x+wx*ww.tileSize] = t.grid[y][x]
		}
	}
}

func (ww *worldWrapper) cut() [][]byte {
	size := len(ww.world) - 2*ww.tilesPerSide
	cutGrid := make([][]byte, size)
	for y := 0; y < size; y++ {
		cutGrid[y] = bytes.Repeat([]byte("X"), size)
	}
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			wx := 2*(x/(ww.tileSize-2)) + 1 + x
			wy := 2*(y/(ww.tileSize-2)) + 1 + y
			//fmt.Println(x, wx)
			cutGrid[y][x] = ww.world[wy][wx]
		}
	}

	return cutGrid
}

func printWorld(world [][]byte) {
	var b strings.Builder
	for y := 0; y < len(world); y++ {
		b.Write(world[y])
		b.WriteString("\n")
	}
	fmt.Println(b.String())
}

// Quick fix to select top left corner
func selectTopLeftCorner(connectsTo map[int]map[int]*match) int {
	for id, t2 := range connectsTo {
		if len(t2) == 2 {

			var matchingSide []int
			for _, match := range t2 {
				matchingSide = append(matchingSide, match.tile1Side)
			}
			if (matchingSide[0] == 1 && matchingSide[1] == 2) || (matchingSide[0] == 2 && matchingSide[1] == 1) {
				return id
			}
		}
	}
	log.Fatal("did not find top left corner")
	return -1
}
