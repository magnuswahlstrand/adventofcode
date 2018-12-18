package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type Grid [][]byte

func (g Grid) count() (int, int, int) {
	var empty, trees, lumberyards int
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[0]); x++ {
			switch g[y][x] {
			case '.':
				empty++
			case '|':
				trees++
			case '#':
				lumberyards++
			}
		}
	}
	return empty, trees, lumberyards
}

func (g Grid) updated() Grid {
	var newGrid Grid = make([][]byte, len(g))

	for y := 0; y < len(g); y++ {
		b := make([]byte, len(g[0]))
		for x := 0; x < len(g[0]); x++ {

			// Check surrounding
			b[x] = g[y][x]

			var trees, lumberyards, empty int
			for dy := max(y-1, 0); dy <= min(y+1, len(g)-1); dy++ {
				for dx := max(x-1, 0); dx <= min(x+1, len(g[0])-1); dx++ {
					if x == dx && y == dy {
						continue
					}

					switch g[dy][dx] {
					case '.':
						empty++
					case '|':
						trees++
					case '#':
						lumberyards++
					}
				}
			}
			// An open acre will become filled with trees if three or more adjacent acres contained trees. Otherwise, nothing happens.
			// An acre filled with trees will become a lumberyard if three or more adjacent acres were lumberyards. Otherwise, nothing happens.
			// An acre containing a lumberyard will remain a lumberyard if it was adjacent to at least one other lumberyard and at least one acre containing trees. Otherwise, it becomes open.
			switch {
			case g[y][x] == '.' && trees >= 3:
				b[x] = '|'
			case g[y][x] == '|' && lumberyards >= 3:
				b[x] = '#'
			case g[y][x] == '#' && (lumberyards < 1 || trees < 1):
				b[x] = '.'
			default:
				b[x] = g[y][x]
			}
		}
		newGrid[y] = b
	}
	return newGrid
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func startingGrid(input []byte) Grid {
	rows := bytes.Split(input, []byte{'\n'})
	return rows
}

func (g Grid) String() string {
	s := ""
	for y := 0; y < len(g); y++ {
		s += fmt.Sprintf("%s\n", g[y])
	}
	return s
}

func main() {
	fmt.Println("Day 18 - 2018\n")

	content, _ := ioutil.ReadFile("input.txt")
	grid := startingGrid(content)
	wanted := 1000000000
	for round := 1; round <= 10110; round++ {

		if round == 10 {
			_, trees, lumberyards := grid.count()
			fmt.Printf("Found value for %d at %d: %d\n", round, round, trees*lumberyards)
		}

		grid = grid.updated()
		if matchingRound(wanted) == round {
			_, trees, lumberyards := grid.count()
			fmt.Printf("Found value for %d at %d: %d\n", wanted, round, trees*lumberyards)
			break
		}
	}
	fmt.Println("")
}

func matchingRound(n int) int {
	period := 28
	p := 486
	return (n-p)%period + p
}
