package main

import (
	"fmt"

	"github.com/pkg/profile"
)

func marble(nPlayers, maxScore int) int {
	marbles := make([]int, maxScore+10)
	current := 0

	score := make(map[int]int)
	p, m := 0, 1
	lenM := 1
	for ; m <= maxScore; p, m = (p+1)%nPlayers, m+1 {

		if m%23 == 0 {
			// fmt.Print("Current is", current)
			score[p] += m
			current = (current - 7 + lenM) % lenM

			//Pick it up, and keep that position
			score[p] += marbles[current]
			// fmt.Println(": Picked up ", marbles[current], "Len is", lenM)
			marbles = append(marbles[:current], marbles[current+1:]...)
			lenM--

		} else {
			current = (current + 2) % lenM
			//marbles = append(marbles, 0)
			lenM++ //Same as append
			copy(marbles[current+1:], marbles[current:])
			marbles[current] = m
		}

	}

	max := -1
	for _, v := range score {
		if v > max {
			max = v
		}
	}

	return max
}

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	fmt.Println("Advent of Code - Day 9 - Go")

	// fmt.Println("Test 1 -", marble(9, 25))            // 45 min
	// fmt.Println("Test 1 -", marble(10, 1618) == 8317) // 45 min
	// fmt.Println("Test 1 -", marble(17, 1104) == 2764) // 45 min

	fmt.Println("Part 1 -", marble(478, 71240)) // 45 min'
	// 30ms --> 200ms --> 1.13s
	// 960ms --> 6.63s --> 25.48s
}
