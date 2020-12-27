package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1(string(input))
	part2(string(input))

	log.Println("success")
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	_, branches := branches(lines)
	paths := pathsForward(branches)
	fmt.Println("answer part 2:", paths[0])
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	n1Diff, n3Diff := diffs(lines)
	fmt.Println("answer part 1:", n1Diff*n3Diff)
}

func mustInt64(s string) int64 {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal("could not convert to int", err)
	}
	return v
}

func diffs(lines []string) (int, int) {
	// Parse ints
	adapters := make([]int64, len(lines)+2)
	for i, line := range lines {
		adapters[i] = mustInt64(line)
	}

	// Sort
	sort.Slice(adapters, func(i, j int) bool {
		return adapters[i] < adapters[j]
	})

	// Add own device
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	// Find differences
	var previous int64 = 0
	var n1Diff, n3Diff int
	for _, adapter := range adapters {
		diff := adapter - previous
		switch diff {
		case 1:
			n1Diff++
		case 3:
			n3Diff++
		}

		previous = adapter
	}
	return n1Diff, n3Diff
}

func branches(lines []string) ([]int64, [][]int) {
	// Parse ints
	adapters := make([]int64, len(lines)+1)
	for i, line := range lines {
		adapters[i] = mustInt64(line)
	}

	// Sort
	sort.Slice(adapters, func(i, j int) bool {
		return adapters[i] < adapters[j]
	})

	// Paths
	paths := make([]int64, len(adapters))
	paths2 := make([][]int, len(adapters))

	// Add own device
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	// Find possible branches
	for i, a1 := range adapters {
		for j, a2 := range adapters[i+1:] {
			if a2 > a1+3 {
				break
			}
			paths2[i] = append(paths2[i], i+j+1)
			paths[i]++
		}
	}
	return paths, paths2
}

func pathsForward(branches [][]int) []int {
	pathsForward := make([]int, len(branches))
	pathsForward[len(pathsForward)-1] = 1
	for i := len(branches) - 2; i >= 0; i-- {
		var sum int
		for _, b := range branches[i] {
			// Get
			sum += pathsForward[b]
		}
		pathsForward[i] = sum
	}
	return pathsForward
}

