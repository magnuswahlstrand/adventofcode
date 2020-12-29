package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
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
	const size = 21
	active := solvePart1(size, input)
	fmt.Println("answer part 1:", active)
}

func part2(input []byte) {
	const size = 21
	active := solvePart2(size, input)
	fmt.Println("answer part 2:", active)
}

var testInput = []byte(`.#.
..#
###`)

type space struct {
	current uint8
	next    uint8
}

func solvePart1(size int, input []byte) int {
	mid := size / 2
	inputXY := bytes.Split(input, []byte("\n"))
	offset := len(inputXY) / 2
	xyzWorld := initWorld(size, inputXY, mid, offset)

	var active int
	cycles := 6
	for r := 0; r < cycles; r++ {
		prepareNext(size, xyzWorld)
		active = applyNext(size, xyzWorld)
	}
	return active
}

func applyNext(size int, xyzWorld [][][]*space) int {
	var activeCount int
	for z := 1; z < size-1; z++ {
		for y := 1; y < size-1; y++ {
			for x := 1; x < size-1; x++ {
				if xyzWorld[z][y][x].next == '#' {
					activeCount++
				}

				xyzWorld[z][y][x].current = xyzWorld[z][y][x].next
			}
		}
	}
	return activeCount
}

func prepareNext(size int, xyzWorld [][][]*space) {
	for z := 1; z < size-1; z++ {
		for y := 1; y < size-1; y++ {
			for x := 1; x < size-1; x++ {

				// Count all occupied in cube
				var count int
				for dz := z - 1; dz <= z+1; dz++ {
					for dy := y - 1; dy <= y+1; dy++ {
						for dx := x - 1; dx <= x+1; dx++ {
							// Skip self
							if dz == z && dy == y && dx == x {
								continue
							}

							if xyzWorld[dz][dy][dx].current == '#' {
								count++
							}
						}
					}
				}

				// Check rules
				isActive := xyzWorld[z][y][x].current == '#'
				switch {
				case isActive && (count == 2 || count == 3):
					// Remain active
					xyzWorld[z][y][x].next = '#'
				case !isActive && count == 3:
					// Become active
					xyzWorld[z][y][x].next = '#'
				default:
					// Become/remain inactive
					xyzWorld[z][y][x].next = 0
				}
			}
		}
	}
}

func initWorld(size int, inputXY [][]byte, mid int, offset int) [][][]*space {
	xyzWorld := make([][][]*space, size)
	for z := 0; z < size; z++ {
		xyzWorld[z] = make([][]*space, size)
		for y := 0; y < size; y++ {
			xyzWorld[z][y] = make([]*space, size)
			for x := 0; x < size; x++ {
				xyzWorld[z][y][x] = &space{current: 0}
			}
		}
	}

	// Add initial active cubes
	for iy := 0; iy < len(inputXY); iy++ {
		for ix := 0; ix < len(inputXY[0]); ix++ {
			if inputXY[iy][ix] == '#' {
				xyzWorld[mid][iy+mid-offset][ix+mid-offset].current = '#'
			}
		}
	}
	return xyzWorld
}

func printWorld(z, mid int, offset int, xyzWorld [][][]*space) {
	for y := mid - offset; y <= mid+offset; y++ {
		for x := mid - offset; x <= mid+offset; x++ {
			switch xyzWorld[z][y][x].current {
			case '#':
				fmt.Print("#")
			default:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
