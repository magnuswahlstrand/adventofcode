package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func maps(input []string) (map[rune]int, map[rune]string) {
	remaining := make(map[rune]int)
	unlocks := make(map[rune]string)
	var before, after rune
	for _, row := range input {
		fmt.Sscanf(row, "Step %c must be finished before step %c can begin.", &before, &after)
		remaining[after]++
		unlocks[before] += string(after)

		// Create before key witha no-op
		remaining[before] += 0
	}
	return remaining, unlocks
}

func steps(input []string) string {
	var instruction string
	remaining, unlocks := maps(input)

	for {

		next := getNext(remaining)
		// No new characters found
		if next == 'Z'+1 {
			return instruction
		}

		instruction += string(next)
		remaining[next] = -1

		for _, r := range unlocks[next] {
			remaining[r]--
		}
	}
}

func sortedKeys(remaining map[rune]int) []rune {
	var keys []rune
	for k := range remaining {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func stepsParallel(input []string, addTime, nWorkers int) int {
	var totalTime int

	remainingTime := make(map[rune]int)
	remaining, unlocks := maps(input)
	sortedKeys := sortedKeys(remaining)

	for {

		// Find workers
		for _, r := range sortedKeys {
			if remaining[r] == 0 && remainingTime[r] == 0 {
				remainingTime[r] = addTime + int(r-'A')

				nWorkers--
				if nWorkers == 0 {
					break
				}
			}
		}

		min := 10000
		for _, val := range remainingTime {
			if val < min {
				min = val
			}
		}
		if min == 10000 {
			return totalTime
		}

		for r, val := range remainingTime {

			remainingTime[r] = val - min

			if remainingTime[r] == 0 {
				remaining[r] = -1
				delete(remainingTime, r)
				nWorkers++

				//Unlock dependencies
				for _, r2 := range unlocks[r] {
					remaining[r2]--
				}
			}
		}
		totalTime += min
		// fmt.Println(remainingTime, totalTime)
	}
}

func getNext(remaining map[rune]int) rune {
	// Find the lowest character that h
	next := 'Z' + 1
	for r, val := range remaining {
		if val == 0 {
			if r < next {
				next = r
			}
		}
	}
	return next
}

func main() {
	fmt.Println("Advent of Code - Day 4 - Go")

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rows := strings.Split(string(content), "\n")

	test := `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`
	rowsTest := strings.Split(test, "\n")

	fmt.Printf("Part 2 - the time taken to assemble the sleigh is %d\n", stepsParallel(rowsTest, 1, 2))
	fmt.Printf("Part 2 - the time taken to assemble the sleigh is %d\n", stepsParallel(rows, 61, 5))

	// 1003 Too low
	// fmt.Printf("Part 1 - the steps to assemble the sleigh are %s\n", steps(rows))
}
