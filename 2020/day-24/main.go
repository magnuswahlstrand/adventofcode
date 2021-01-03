package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
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
	world := createWorld(input)
	blackCount := count(world)
	fmt.Println("answer part 1:", blackCount)
}

// 3672 too low
func part2(input []byte) {
	world := createWorld(input)
	world2, processList := prepareForProcessing(world)
	count := runNTimes(100, processList, world2)
	fmt.Println("answer part 2:", count)
}

func wrapPos(x int, y int) [2]int {
	return [2]int{x, y}
}

func count(world map[[2]int]int) int {
	var countBlack int
	for _, val := range world {
		if isActive(val) {
			countBlack++
		}
	}
	return countBlack
}

func isActive(val int) bool {
	return val%2 == 1
}

func createWorld(input []byte) map[[2]int]int {
	world := map[[2]int]int{}
	for _, line := range strings.Split(string(input), "\n") {
		var x, y int
		for _, dir := range parseDirections(line) {
			x, y = nextPos(x, y, dir)
		}
		world[[2]int{x, y}]++
	}
	return world
}

func nextPos(x, y int, dir string) (int, int) {

	isOddRow := isOdd(y)
	switch dir {
	case "w", "nw", "sw":
		x--
	case "e", "ne", "se":
		x++
	}

	switch dir {
	case "nw":
		y--
		if isOddRow {
			x++
		}
	case "ne":
		y--
		if !isOddRow {
			x--
		}
	case "sw":
		y++
		if isOddRow {
			x++
		}
	case "se":
		y++
		if !isOddRow {
			x--
		}
	}

	return x, y
}

func isOdd(v int) bool {
	if v < 0 {
		return (-v)%2 == 1
	}
	return v%2 == 1
}

func parseDirections(s string) []string {
	re, err := regexp.Compile("((se|sw|nw|ne|e|w))")
	if err != nil {
		log.Fatal("failed to compile regex", err)
	}
	return re.FindAllString(s, -1)
}

// Part 1: 252 too low

type tile struct {
	//coord           [2]int
	currentlyActive bool
	nextActive      bool
}

var (
	directions = []string{"sw", "se", "e", "ne", "nw", "w"}
)

func runNTimes(n int, processList map[[2]int]*tile, world2 map[[2]int]*tile) int {
	var count int
	for i := 0; i < n; i++ {
		prepareTiles(processList, world2)
		count, processList = updateAndCount(world2)
	}
	return count
}

func prepareTiles(processList map[[2]int]*tile, world2 map[[2]int]*tile) {
	handled := map[[2]int]bool{}
	for coord := range processList {
		var activeNeighbors int

		for _, dir := range directions {
			neighbor := wrapPos(nextPos(coord[0], coord[1], dir))

			if c, ok := world2[neighbor]; ok && c.currentlyActive {
				activeNeighbors++
			}

			// Add to list if not already handled or planned to be processed
			_, toBeProcessed := processList[neighbor]
			if !toBeProcessed && !handled[neighbor] {
				// TODO: exclude already in planned to be handled
				updateNeighbor(neighbor, world2, handled)
			}
		}

		// Rule 1: Any black tile with zero or more than 2 black tiles immediately adjacent to it is flipped to white.
		switch activeNeighbors {
		case 1, 2:
			world2[coord].nextActive = true
		default:
			world2[coord].nextActive = false
		}

		// Rule 2: Any white tile with exactly 2 black tiles immediately adjacent to it is flipped to black.
		// - Ignore - All handled tiles are black (active)

		// Mark as handled
		handled[coord] = true
	}
}

func updateAndCount(world2 map[[2]int]*tile) (int, map[[2]int]*tile) {
	processList := map[[2]int]*tile{}
	var count int
	for coord, tile := range world2 {
		tile.currentlyActive, tile.nextActive = tile.nextActive, false
		if tile.currentlyActive {
			count++
			processList[coord] = tile
		}
	}
	return count, processList
}

func prepareForProcessing(world map[[2]int]int) (map[[2]int]*tile, map[[2]int]*tile) {
	world2 := map[[2]int]*tile{}
	processList := map[[2]int]*tile{}
	for coord, val := range world {
		if !isActive(val) {
			continue
		}
		t := &tile{
			//coord:           coord,
			currentlyActive: true,
			nextActive:      false,
		}
		world2[coord] = t
		processList[coord] = t
	}
	return world2, processList
}

func updateNeighbor(neighbor [2]int, world2 map[[2]int]*tile, handled map[[2]int]bool) {
	var count int
	for _, dir2 := range directions {
		neighbor2 := wrapPos(nextPos(neighbor[0], neighbor[1], dir2))

		if c, ok := world2[neighbor2]; ok && c.currentlyActive {
			count++
		}
	}

	// Rule 1: Any black tile with zero or more than 2 black tiles immediately adjacent to it is flipped to white.
	// - Ignore - "Neighbors" are only white

	// Rule 2: Any white tile with exactly 2 black tiles immediately adjacent to it is flipped to black.
	if count == 2 {
		world2[neighbor] = &tile{
			//coord:           neighbor,
			currentlyActive: false,
			nextActive:      true,
		}
	}

	// Mark as handled
	handled[neighbor] = true
}
