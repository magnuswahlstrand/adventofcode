package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var str string = "rvefnvyxzbodgpnpkumawhijsc"

func hasTwoThreeIdentical(s string) (bool, bool) {
	charMap := make(map[rune]int)
	for _, c := range s {
		charMap[c] += 1
	}
	hasTwo := false
	hasThree := false
	for _, v := range charMap {
		switch v {
		case 2:
			hasTwo = true
		case 3:
			hasThree = true
		default:
		}
	}
	return hasTwo, hasThree
}

func main() {

	fmt.Println("Advent of Code - Day 2 - Go")
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var countTwo, countThree int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		hasTwo, hasThree := hasTwoThreeIdentical(scanner.Text())

		if hasTwo {
			countTwo++
		}
		if hasThree {
			countThree++
		}
		fmt.Println(hasTwo, hasThree)
	}

	checkSum := countTwo * countThree
	fmt.Printf("The checksum for boxes with twos (%d) and threes (%d) is: %d", countTwo, countThree, checkSum)
}
