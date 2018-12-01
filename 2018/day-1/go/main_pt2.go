package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func findDuplicateFrequency(filename string) (int, error) {

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	s := string(b)
	lines := strings.Split(s, "\n")

	visitedFrequencies := make(map[int]struct{})

	freq := 0
	for {
		for _, line := range lines {
			i, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal("Input not valid", err)
			}
			freq += i

			// Check if frequency has been visited before
			if _, ok := visitedFrequencies[freq]; ok {
				return freq, nil
			}

			// If not, add it to the list
			visitedFrequencies[freq] = struct{}{}
		}
	}
}

func main() {
	fmt.Println("---------------------------")
	fmt.Println("Advent of Code - Day 2 - Go")

	duplicatedFrequency, err := findDuplicateFrequency("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("First duplicate frequency is:", duplicatedFrequency)
}
