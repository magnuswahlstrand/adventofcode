package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

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
	w := parseInput(input)
	for w.step() {
	}
	count := strings.Count(w.String(), string(SeatOccupied))
	fmt.Println("answer part 1:", count)
}

// Broken
func part2(input []byte) {
	w := parseInput2(input)
	for w.stepTwo() {
	}

	count := strings.Count(w.String(), string(SeatOccupied))
	fmt.Println("answer part 2:", count)
}

type coord []int

type World struct {
	current   [][]byte
	buffer    [][]byte
	neighbors map[int][]coord
	tolerance int
}

func (w *World) String() string {
	return string(bytes.Join(w.current, []byte("\n")))
}

var (
	SeatEmpty    = byte('L')
	SeatOccupied = byte('#')
	Floor        = byte('.')
)

func (w *World) step() bool {
	var updatedCount int
	for i := 0; i < len(w.current); i++ {
		for j := 0; j < len(w.current[0]); j++ {
			c := w.current[i][j]
			switch {
			case c == SeatEmpty && !w.adjacentOccupied(i, j):
				w.buffer[i][j] = SeatOccupied
				updatedCount++
			case c == SeatOccupied && w.adjacentCount(i, j) >= w.tolerance:
				w.buffer[i][j] = SeatEmpty
				updatedCount++
			default:
				w.buffer[i][j] = c
			}
		}
	}

	w.current, w.buffer = w.buffer, w.current

	return updatedCount > 0
}
func (w *World) stepTwo() bool {
	var updatedCount int
	for i := 0; i < len(w.current); i++ {
		for j := 0; j < len(w.current[0]); j++ {
			v := w.current[i][j]
			switch {
			case v == SeatEmpty && !w.adjacentOccupied2(i, j):
				w.buffer[i][j] = SeatOccupied
				updatedCount++
			case v == SeatOccupied && w.adjacentCount2(i, j) >= w.tolerance:
				w.buffer[i][j] = SeatEmpty
				updatedCount++
			default:
				w.buffer[i][j] = v
			}
		}
	}

	w.current, w.buffer = w.buffer, w.current
	return updatedCount > 0
}

func (w *World) adjacentOccupied2(i int, j int) bool {
	for _, coord := range w.neighbors[100*i+j] {
		if w.current[coord[0]][coord[1]] == SeatOccupied {
			return true
		}
	}
	return false
}

func (w *World) adjacentCount2(i int, j int) int {
	var count int
	for _, coord := range w.neighbors[100*i+j] {
		if w.current[coord[0]][coord[1]] == SeatOccupied {
			count++
		}
	}
	return count
}

func (w *World) adjacentOccupied(i int, j int) bool {
	return w.isOcc(i-1, j-1) || w.isOcc(i-1, j) || w.isOcc(i-1, j+1) ||
		w.isOcc(i, j-1) || w.isOcc(i, j+1) ||
		w.isOcc(i+1, j-1) || w.isOcc(i+1, j) || w.isOcc(i+1, j+1)
}

func (w *World) adjacentCount(i int, j int) int {
	var count int
	for _, occupied := range []bool{w.isOcc(i-1, j-1), w.isOcc(i-1, j), w.isOcc(i-1, j+1), w.isOcc(i, j-1), w.isOcc(i, j+1), w.isOcc(i+1, j-1), w.isOcc(i+1, j), w.isOcc(i+1, j+1)} {
		if !occupied {
			continue
		}

		count++
	}

	return count
}

func (w *World) isSeat(i, j int) bool {
	if i < 0 || i >= len(w.current) {
		return false
	}

	if j < 0 || j >= len(w.current[0]) {
		return false
	}

	return w.current[i][j] == SeatOccupied || w.current[i][j] == SeatEmpty
}

func (w *World) isOutOfBounds(i, j int) bool {
	if i < 0 || i >= len(w.current) {
		return true
	}

	if j < 0 || j >= len(w.current[0]) {
		return true
	}

	return false
}

func (w *World) isOcc(i, j int) bool {
	if i < 0 || i >= len(w.current) {
		return false
	}

	if j < 0 || j >= len(w.current[0]) {
		return false
	}

	return w.current[i][j] == SeatOccupied
}

func parseInput(input []byte) World {
	tmp := make([]byte, len(input))
	copy(tmp, input)
	current := bytes.Split(input, []byte("\n"))
	buffer := bytes.Split(tmp, []byte("\n"))

	w := World{
		current:   current,
		buffer:    buffer,
		tolerance: 4,
	}

	var neighbors = make(map[int][]coord, len(current)*len(current[0]))
	for i := 0; i < len(current); i++ {
		for j := 0; j < len(current[0]); j++ {
			var nn []coord
			for _, coord := range []coord{
				{i - 1, j - 1},
				{i - 1, j},
				{i - 1, j + 1},

				{i, j - 1},
				{i, j + 1},

				{i + 1, j - 1},
				{i + 1, j},
				{i + 1, j + 1},
			} {
				if w.isSeat(coord[0], coord[1]) {
					nn = append(nn, coord)
				}

				neighbors[100*i+j] = nn
			}
		}
	}

	w.neighbors = neighbors
	return w
}

// 2315
func parseInput2(input []byte) World {
	tmp := make([]byte, len(input))
	copy(tmp, input)
	current := bytes.Split(input, []byte("\n"))
	buffer := bytes.Split(tmp, []byte("\n"))

	w := World{
		current:   current,
		buffer:    buffer,
		tolerance: 5,
	}

	var neighbors = make(map[int][]coord, len(current)*len(current[0]))
	for i := 0; i < len(current); i++ {
		for j := 0; j < len(current[0]); j++ {
			var nn []coord
			for _, coord := range []coord{
				{- 1, - 1},
				{- 1, 0},
				{- 1, + 1},

				{0, - 1},
				{0, + 1},

				{1, - 1},
				{1, 0},
				{1, + 1},
			} {
				ii := i
				jj := j
				for {
					ii += coord[0]
					jj += coord[1]
					if i == 3 && j == 3 {
						//fmt.Println("xx", ii, jj)
					}
					if w.isOutOfBounds(ii, jj) {
						if i == 3 && j == 3 {
							//fmt.Println("oob", ii, jj)
						}
						break
					}

					if w.isSeat(ii, jj) {
						if i == 3 && j == 3 {
							//fmt.Println("seat", ii, jj)
						}
						nn = append(nn, []int{ii, jj})
						break
					}
				}

				neighbors[100*i+j] = nn
			}
		}
	}

	w.neighbors = neighbors
	return w
}
