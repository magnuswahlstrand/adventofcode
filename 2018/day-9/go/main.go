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

type node struct {
	prev, next int
	val        int //For simplicity
}

func linkedMarbles(nPlayers, maxScore int) int {
	marbles := make([]node, maxScore+10, maxScore+10)
	current := &node{
		0,
		0,
		0,
	}
	marbles = append(marbles, *current)

	score := make(map[int]int)
	p, m := 0, 1
	for ; m <= maxScore; p, m = (p+1)%nPlayers, m+1 {

		if m%23 == 0 {
			// fmt.Print("Current is", current)
			score[p] += m

			for i := 0; i < 7; i++ {
				current = &marbles[current.prev]
			}
			// current = (current - 7 + len(marbles)) % len(marbles)
			score[p] += current.val

			// Link past node
			marbles[current.prev].next = current.next
			marbles[current.next].prev = current.prev

			// To make bugs easier to find :-)
			current = &marbles[current.next]
		} else {
			prev := &marbles[current.next]
			next := &marbles[prev.next]

			// Insert between
			prev.next = m
			next.prev = m
			marbles[m] = node{
				prev: prev.val,
				next: next.val,
				val:  m,
			}

			current = &marbles[m]
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

	fmt.Println("Test 1 -", linkedMarbles(9, 25))                                     // 45 min
	fmt.Println("Test 1 -", linkedMarbles(10, 1618), linkedMarbles(10, 1618) == 8317) // 45 min
	fmt.Println("Test 1 -", linkedMarbles(17, 1104), linkedMarbles(17, 1104) == 2764) // 45 min

	// fmt.Println("Part 1 -", marble(478, 71240*3)) // 45 min'
	fmt.Println("Part 1 -", linkedMarbles(478, 71240*100)) // 45 min'
}
