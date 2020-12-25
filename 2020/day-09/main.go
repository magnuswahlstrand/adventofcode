package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
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
	val := findInvalid(input, 25)
	seq := findSequence(input, int64(val))
	fmt.Println(seq)
	weakness := encryptionWeakness(seq)
	fmt.Println("answer part 2:", weakness)
}

func part1(input string) {
	n := findInvalid(input, 25)
	fmt.Println("answer part 1:", n)
}

func mustInt(line string) int {
	v, err := strconv.Atoi(line)
	if err != nil {
		log.Fatal("could not convert to int", err)
	}
	return v
}

func mustInt64(s string) int64 {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal("could not convert to int", err)
	}
	return v
}

func findInvalid(input string, preambleLength int) int {
	lines := strings.Split(input, "\n")
	var numbers []int
	var i int

	// Setup initial numbers
	for i < preambleLength {
		numbers = append(numbers, mustInt(lines[i]))
		i++
	}

	// Iterate over all lines
	for i < len(lines) {

		current := mustInt(lines[i])
		// Check if current line is the sum of one of the preceding ones
		var matchFound bool

	outer:
		for i, n := range numbers {
			for _, m := range numbers[i+1:] {
				if current == n+m {
					matchFound = true
					break outer
				}
			}
		}

		if !matchFound {
			return current
		}

		// Match found, put it in the list of numbers
		replaceIndex := i % preambleLength
		numbers[replaceIndex] = current

		i++
	}

	log.Fatal("no invalid found!")
	return -1
}

func encryptionWeakness(seq []int64) int64 {
	var min, max int64 = math.MaxInt64, math.MinInt64
	for _, val := range seq {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return min + max
}

func findSequence(input string, target int64) []int64 {
	lines := strings.Split(input, "\n")
	numbers := make([]int64, len(lines))
	for i, s := range lines {
		numbers[i] = mustInt64(s)
	}

	for i := range numbers {
		var sum int64
		for j, val := range numbers[i:] {
			sum += val
			if sum == target {
				return numbers[i : i+j+1]
			}
			if sum > target {
				break
			}
		}
	}

	log.Fatal("sequence not found :-(")
	return nil
}
