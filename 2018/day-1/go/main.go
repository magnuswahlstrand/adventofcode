package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code - Day 1")

	inputFile := "input.txt"
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	freq := 0
	for scanner.Scan() {
		txt := scanner.Text()

		// Check if number is positive or negative
		factor := 1
		if strings.HasPrefix(txt, "-") {
			factor = -1
		}

		// Get the number part from a row -14 --> 14, or +5 --> 5
		numberPart := txt[1:]
		i, err := strconv.Atoi(numberPart)
		if err != nil {
			log.Fatal("Input not valid", err)
		}
		freq += factor * i
	}

	fmt.Println("Final frequency is:", freq)
}
