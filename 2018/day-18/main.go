package main

import (
	"bytes"
	"container/ring"
	"fmt"
	"io/ioutil"
	"os"
)

type Grid [][]byte

type CircularBuffer struct {
	*ring.Ring
}

func (cb CircularBuffer) String() string {
	s := ""
	for i := 0; i < 2*cb.Len(); i++ {
		s += fmt.Sprintf("%3d\n", cb.Value) //strconv.Itoa(cb.Value.(int))
		cb.Ring = cb.Next()
	}
	return s
}

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

func newBuffer(n int) CircularBuffer {
	r := ring.New(n)
	for i := 0; i < r.Len(); i++ {
		r.Value = -1
		r = r.Next()
	}
	return CircularBuffer{
		r,
	}
}

func main() {
	fmt.Println("Day 18 - 2018\n")

	content, _ := ioutil.ReadFile("input.txt")
	grid := startingGrid(content)

	buf := newBuffer(40)
	found := make(map[int]bool)
	wanted := 1000000000
	matchingRound := -1
	for round := 1; round <= 500; round++ {

		grid = grid.updated()
		_, trees, lumberyards := grid.count()
		score := trees * lumberyards

		if _, scoreSeen := found[score]; scoreSeen {

			// Search for periodicity
			if isPeriodic, period := findPeriod(buf, score); isPeriodic {

				matchingRound = (wanted-round)%period + round
				fmt.Printf("Found that round=%d should have the same score as round=%d\n", wanted, matchingRound)
			}

		}

		if matchingRound > 0 && matchingRound == round {
			fmt.Printf("Round %d has score: %d\n", round, score)
			os.Exit(0)
		}

		// Save score and mark as used
		buf.Value = score
		buf.Ring = buf.Next()
		found[score] = true
	}
	// Found value for 10 at 10: 652344
	// Found value for 1000000000 at 496: 202272
}

const requiredLength = 5

type match struct {
	*ring.Ring
	index int
}

func findPeriod(cb CircularBuffer, score int) (bool, int) {
	search := cb.Ring

	// Search for position in list
	matches := []match{}
	for i := 0; i < cb.Len()-requiredLength; i++ {
		if search.Value == score {
			matches = append(matches, match{search, i})
		}
		search = search.Prev()
	}

	// fmt.Println("s:", search.Value)
	if len(matches) == 0 {
		return false, -1
	}

Outer:
	for _, match := range matches {
		cmp := cb.Ring
		for i := 0; i < requiredLength; i++ {
			cmp, match.Ring = cmp.Prev(), match.Prev()
			if cmp.Value != match.Value {
				continue Outer
			}
		}
		return true, match.index
	}
	return false, -1
}

func matchingRound(n int) int {
	period := 28
	p := 486
	return (n-p)%period + p
}
