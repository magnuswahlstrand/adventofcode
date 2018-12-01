package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Advent of Code - Day 1 - Go")

	inputFile := "../input.txt"
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	freq := 0
	for scanner.Scan() {
		txt := scanner.Text()

		i, err := strconv.Atoi(txt)
		if err != nil {
			log.Fatal("Input not valid", err)
		}
		freq += i
	}

	fmt.Println("Final frequency is:", freq)
}
