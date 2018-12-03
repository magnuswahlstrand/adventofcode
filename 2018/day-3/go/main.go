package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func markFabric(elfClaims []claim, fabricSize int) [][]int {
	fabric := make([][]int, fabricSize)
	for i := range fabric {
		fabric[i] = make([]int, fabricSize)
	}

	for _, c := range elfClaims {
		for dy := 0; dy < c.H; dy++ {
			for dx := 0; dx < c.W; dx++ {
				fabric[c.Y+dy][c.X+dx]++
			}
		}
	}

	return fabric
}

func countOverlap(fabric [][]int) int {

	// Calculate overlap
	overlap := 0
	for _, row := range fabric {
		for _, square := range row {
			if square > 1 {
				overlap++
			}
		}
	}
	return overlap
}

type claim struct {
	ID, X, Y, W, H int
}

func parsedClaims(input string) []claim {
	parsedInput := []claim{}

	// Add elf claims to the fabric
	for _, row := range strings.Split(input, "\n") {
		var c claim
		fmt.Sscanf(row, "#%d @ %d,%d: %dx%d", &c.ID, &c.X, &c.Y, &c.W, &c.H)
		parsedInput = append(parsedInput, c)
	}

	return parsedInput
}

func nonOverlappingClaims(fabric [][]int, elfClaims []claim) int {

	for _, c := range elfClaims {
		overlapFound := false

	breaker:
		for dy := 0; dy < c.H; dy++ {
			for dx := 0; dx < c.W; dx++ {
				if fabric[c.Y+dy][c.X+dx] > 1 {
					overlapFound = true
					continue breaker
				}
			}
		}

		if overlapFound == false {
			return c.ID
		}
	}
	return -1
}

func main() {
	fmt.Println("Advent of Code - Day 3 - Go")

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	parsed := parsedClaims(string(content))
	fabric := markFabric(parsed, 1000)
	fmt.Println("The number of square of fabric overlapping is:", countOverlap(fabric))
	fmt.Println("The only square not overlapping is:", nonOverlappingClaims(fabric, parsed))
}
