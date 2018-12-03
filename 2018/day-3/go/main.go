package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

func combineElfInput(input string, fabricSize int) [][]int {
	fabric := make([][]int, fabricSize)
	for i := range fabric {
		fabric[i] = make([]int, fabricSize)
	}

	// Add elf claims to the fabric
	re := regexp.MustCompile(`#(\d*) @ (\d*),(\d*): (\d*)x(\d*)`)
	parsed := re.FindAllStringSubmatch(input, -1)

	for _, elfInput := range parsed {
		var e [5]int

		// Skip the full match and the id
		for i, p := range elfInput[2:] {
			// Expect good formatted input
			e[i], _ = strconv.Atoi(p)
		}

		xStart, yStart := e[0], e[1]
		xSteps, ySteps := e[2], e[3]
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

func main() {
	fmt.Println("Advent of Code - Day 3 - Go")

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fabric := combineElfInput(string(content), 1000)
	fmt.Println("The number of square of fabric overlapping is:", countOverlap(fabric))
}
