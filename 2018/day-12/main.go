package main

import (
	"fmt"
	"strings"

	"github.com/pkg/profile"
)

const (
	labelLen       = len("initial state: ")
	initialPadding = "......."
)

func parseInput(input string) (map[string]byte, string) {
	var state string
	rows := strings.Split(input, "\n")

	// Parse initial state + padding
	state = initialPadding + rows[0][labelLen:] + initialPadding

	// Parse rules
	rules := make(map[string]byte)
	for _, row := range rows[2:] {
		rules[row[0:5]] = row[9]
	}

	return rules, state

}

const padding = ".."

func next(state string, rules map[string]byte) string {
	nextState := padding

	for i := 0; i < len(state)-4; i++ {
		current := state[i : i+5]
		switch rules[current] {
		case '#':
			nextState += "#"
		default:
			nextState += "."

		}
		// fmt.Println(current, string(rules[current]))
	}

	if strings.ContainsAny(nextState[len(nextState)-5:], "#") {
		nextState += padding + "."
	}
	return nextState
}

func countSeeds(state string) int {
	var c int
	for i, pot := range state {
		if pot == '#' {
			c += i - len(initialPadding)
		}
	}
	return c
}

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	rules, state := parseInput(realInput)

	var gen int
	for ; gen < 130; gen++ {

		fmt.Printf("% 3d: %.100s % 3d\n", gen, state, countSeeds(state))
		state = next(state, rules)

	}
	fmt.Printf("Total number of seeds after generation %d is %d\n", gen, countSeeds(state))

	gen = 50000000000
	fmt.Printf("Total number of seeds after generation %d is %d\n", gen, expected(gen))

	// Guessed 6299999999457, was too high
	// Guessed 6299999999331, was too high as well
	// Guessed 5450000001275, it was incorrect
	// Guessed 5450000001166, it was correct!
}

func expected(gen int) int {
	return 28416 + ((gen - 250) * 109)
}

var input = `initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #`

var realInput = `initial state: #...#...##..####..##.####.#...#...#.#.#.#......##....#....######.####.##..#..#..##.##..##....#######

.#### => .
...#. => .
.##.. => #
#.##. => .
#..## => .
##### => #
####. => #
.##.# => #
#.### => .
...## => #
.#.## => #
#..#. => #
#.#.. => #
.###. => #
##.## => #
##..# => .
.#... => #
###.# => .
..##. => .
..... => .
###.. => #
..#.# => .
.#..# => #
##... => #
#.... => .
##.#. => .
..#.. => #
....# => .
#...# => .
#.#.# => #
..### => .
.#.#. => #`
