package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return i
}

func main() {
	t := time.Now()
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	part1(input)
	part2(input)

	log.Println("success in", time.Since(t))
}

func part1(input []byte) {
	fmt.Println("answer part 1:", findNumberPart1(input, 2020))

}

func part2(input []byte) {
	fmt.Println("answer part 1:", findNumberPart1(input, 30000000))
}

func findNumberPart1(input []byte, endAfter int) int {
	var previousWord int
	buffer1 := map[int]int{}

	round := 1
	for _, s := range strings.Split(string(input), ",") {
		i := mustInt(s)

		buffer1[i] = round
		//fmt.Println(round, i)
		previousWord = i
		round++
	}

	for round <= endAfter {
		var wordToSpeak int
		previousRound := round - 1

		// Check if was first time for previous word
		_, wasOld := buffer1[previousWord]
		if !wasOld {
			// Print 0
			wordToSpeak = 0
			//fmt.Println(previousWord, "was first time, speak 0")
		} else {
			lastTimeSpoken := buffer1[previousWord]
			wordToSpeak = previousRound - lastTimeSpoken
			//fmt.Println(previousWord, "was NOT first time, speak diff", wordToSpeak)
		}

		// Update previous word only now
		buffer1[previousWord] = previousRound

		// Has word been spoken before?

		// Speak
		//fmt.Printf("%d: speak %d\n", round, wordToSpeak)

		// Prepare next round
		previousWord = wordToSpeak
		round++
	}
	return previousWord
}
