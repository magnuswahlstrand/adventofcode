package main

import (
	"fmt"

	"github.com/pkg/profile"
)

type node struct {
	prev, next *node
	val        int //For simplicity
}

func linkedMarbles(nPlayers, maxScore int) int {
	current := &node{val: 0}
	current.prev = current
	current.next = current

	score := make(map[int]int)
	p, m := 0, 1
	for ; m <= maxScore; p, m = (p+1)%nPlayers, m+1 {

		if m%23 == 0 {
			// fmt.Print("Current is", current)
			score[p] += m

			for i := 0; i < 7; i++ {
				current = current.prev
			}
			score[p] += current.val

			// Link over the old node
			current.prev.next = current.next
			current.next.prev = current.prev

			current = current.next
		} else {
			// Find the two nodes to insert between
			prev := current.next
			next := current.next.next

			n := &node{
				prev: prev,
				next: next,
				val:  m,
			}

			// Insert between
			prev.next = n
			next.prev = n
			current = n
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

	fmt.Println("Test 1 -", linkedMarbles(9, 25))
	fmt.Println("Test 1 -", linkedMarbles(10, 1618), linkedMarbles(10, 1618) == 8317)
	fmt.Println("Test 1 -", linkedMarbles(17, 1104), linkedMarbles(17, 1104) == 2764)

	//fmt.Println("Part 1 -", marble(478, 71240*3))
	fmt.Println("Part 1 -", linkedMarbles(478, 71240))
	fmt.Println("Part 2 -", linkedMarbles(478, 71240*100))
}
