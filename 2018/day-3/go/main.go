package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

func combineElfInput(elfInput [][5]int, fabricSize int) [][]int {
	fabric := make([][]int, fabricSize)
	for i := range fabric {
		fabric[i] = make([]int, fabricSize)
	}

	for _, e := range elfInput {
		xStart, yStart := e[1], e[2]
		xSteps, ySteps := e[3], e[4]
		for dy := 0; dy < ySteps; dy++ {
			for dx := 0; dx < xSteps; dx++ {
				fabric[yStart+dy][xStart+dx]++
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

func parseElfInput(input string) [][5]int {
	parsedInput := [][5]int{}

	// Add elf claims to the fabric
	re := regexp.MustCompile(`#(\d*) @ (\d*),(\d*): (\d*)x(\d*)`)
	parsed := re.FindAllStringSubmatch(input, -1)

	for _, elfInput := range parsed {
		var e [5]int

		// Skip the full match
		for i, p := range elfInput[1:] {
			// Expect good formatted input
			e[i], _ = strconv.Atoi(p)
		}

		parsedInput = append(parsedInput, e)
	}

	return parsedInput
}

func nonOverlapping(fabric [][]int, elfInput [][5]int) int {

	for _, e := range elfInput {
		ID := e[0]
		xStart, yStart := e[1], e[2]
		xSteps, ySteps := e[3], e[4]
		overlapFound := false
		for dy := 0; dy < ySteps; dy++ {
			for dx := 0; dx < xSteps; dx++ {

				if fabric[yStart+dy][xStart+dx] > 1 {
					overlapFound = true
				}
			}
		}

		if overlapFound == false {
			return ID
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

	parsed := parseElfInput(string(content))
	fabric := combineElfInput(parsed, 1000)
	fmt.Println("The number of square of fabric overlapping is:", countOverlap(fabric))
	fmt.Println("The only square not overlapping is:", nonOverlapping(fabric, parsed))
}
