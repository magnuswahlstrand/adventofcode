package main

import (
	"fmt"
	"strconv"
	"strings"
)

func intSlice(in string) []int {
	var n []int
	for i, r := range strings.Fields(in) {
		val, err := strconv.Atoi(r)
		if err == nil {
			n[i] = val
		}
	}
	return n
}

const (
	memorySizeSm = 4
	memorySizeLg = 16
)

func main() {

	test2 := "4	10	4	1	8	4	9	14	5	1	14	15	0	15	3	5"

	const memorySize = memorySizeLg
	memory := make(map[[memorySize]int]int)
	var banks [memorySize]int
	for i, r := range strings.Fields(test2) {
		val, err := strconv.Atoi(r)
		if err == nil {
			banks[i] = val
		}
	}

	fmt.Println("Day 6 - 2017")

	var r int
	for ; memory[banks] < 1; r++ {
		memory[banks] = r

		// Find largest
		max := -1
		var maxI int
		for i, b := range banks {
			if b > max {
				max = b
				maxI = i
			}
		}

		// Redistribute its memory
		banks[maxI] = 0
		i := (maxI + 1)
		for p := max; p > 0; p-- {
			i %= len(banks)
			banks[i]++

			// Move forward
			i++
		}
	}
	fmt.Println("Infinite loop found at", r)               // Took 28 minutes
	fmt.Println("Infinite loop found at", r-memory[banks]) // Took 2 minutes
}
