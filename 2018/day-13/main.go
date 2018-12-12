package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

type cart struct {
	dir     string
	x, y    int
	turn    Turn
	crashed bool
}

// Turn is an enum of possible ways to turn
type Turn int

const (
	left     Turn = 0
	straight Turn = 1
	right    Turn = 2
)

var mapping = map[string]string{
	">S": ">",
	"<S": "<",
	"vS": "v",
	"^S": "^",

	">L": "^",
	"<L": "v",
	"vL": ">",
	"^L": "<",

	">R": "v",
	"<R": "^",
	"vR": "<",
	"^R": ">",

	">-": ">",
	"<-": "<",
	"v|": "v",
	"^|": "^",

	">\\": "v",
	"<\\": "^",
	"v\\": ">",
	"^\\": "<",

	">/": "^",
	"</": "v",
	"v/": "<",
	"^/": ">",
}

func (c *cart) move(x, y int, target string) {

	// Turn
	if target == "+" {
		target = c.nextTurn()
	}

	if _, ok := mapping[c.dir+target]; !ok {
		log.Fatal("Should not happend! '", target, "'")
	}

	newDir := mapping[c.dir+target]
	// fmt.Printf("Target is %s, moving from %s to %s (c.x=%d,x=%d,c.y=%d,y=%d)\n", target, c.dir, newDir, c.x, x, c.y, y)

	c.dir = newDir
	c.x = x
	c.y = y
}

func (c *cart) nextTurn() string {
	t := Turn(c.turn)
	c.turn = (c.turn + 1) % 3

	var s string
	switch t {
	case left:
		s = "L"
	case straight:
		s = "S"
	case right:
		s = "R"
	}
	return s
}

func (c cart) next() (int, int) {

	switch c.dir {
	case ">":
		return c.x + 1, c.y
	case "<":
		return c.x - 1, c.y
	case "^":
		return c.x, c.y - 1
	case "v":
		return c.x, c.y + 1
	}

	return -1, -1
}

// Board represents a 2d grid
type Board []string

// Get gets the value in a 2d grid at x,y coordinates
func (b Board) Get(x, y int) string {
	return string(b[y][x])
}

// Set sets the value in a 2d grid at x,y coordinates
func (b *Board) Set(x, y int, s string) {
	updated := (*b)[y][:x] + s + (*b)[y][x+1:]
	(*b)[y] = updated
}

func (b Board) String() string {
	s := ""
	for _, row := range b {
		s += fmt.Sprintln(row)
	}
	return s
}

func parseBoard(in string) (Board, []cart) {
	var carts []cart
	var board Board
	board = strings.Split(in, "\n")

	// Find carts
	for y, row := range board {
		for x, r := range row {

			switch r {
			case 'v', '^':
				board.Set(x, y, "|")
				carts = append(carts, cart{
					string(r),
					x,
					y,
					left,
					false,
				})
			case '<', '>':
				board.Set(x, y, "-")
				carts = append(carts, cart{
					string(r),
					x,
					y,
					left,
					false,
				})
			default:

			}
		}
	}

	return board, carts
}

func initialBoard(input string) (Board, Board, []cart) {

	emptyBoard, carts := parseBoard(input)

	var board Board
	board = make([]string, len(emptyBoard))
	copy(board, emptyBoard)

	for _, cart := range carts {
		board.Set(cart.x, cart.y, cart.dir)
	}

	return emptyBoard, board, carts
}

func filterCrashed(carts []cart) []cart {
	var filtered []cart
	for _, c := range carts {
		if !c.crashed {

			filtered = append(filtered, c)
		}
	}
	return filtered
}

func main() {

	testMode := false
	filename := "input.txt"
	if testMode {
		// filename = "example.txt"
		filename = "example2.txt"
	}

	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Setup initial board
	emptyBoard, board, carts := initialBoard(string(input))

	var round int
	for round = 0; round < 20000 && len(carts) > 1; round++ {

		// Sort cart based on location
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].y == carts[j].y {
				return carts[i].x > carts[j].x
			}
			return carts[i].y < carts[j].y
		})

		for i := range carts {

			if carts[i].crashed {
				continue
			}

			// Find next step
			nX, nY := carts[i].next()

			// Check next step
			target := board.Get(nX, nY)

			// Patch previous step
			board.Set(carts[i].x, carts[i].y, emptyBoard.Get(carts[i].x, carts[i].y))

			// Move
			switch target {
			case ">", "<", "^", "v":

				// Remove this cart
				carts[i].crashed = true

				// Remove other cart
				for j := range carts {
					if carts[j].x == nX && carts[j].y == nY {
						carts[j].crashed = true
						break
					}
				}

				// Remove from board
				board.Set(nX, nY, emptyBoard.Get(nX, nY))

				continue

			default:
				carts[i].move(nX, nY, target)
			}

			board.Set(carts[i].x, carts[i].y, carts[i].dir)
		}
		if testMode {
			fmt.Println(board)
		}

		carts = filterCrashed(carts)
	}

	fmt.Printf("Only 1 remaining after round %d at x=%d,y=%d", round, carts[0].x, carts[0].y)

	// Guessed 137,54, it was wrong, had forgotten to sort inside loop
	// Final answer: round 12011 at x=50,y=100

}
