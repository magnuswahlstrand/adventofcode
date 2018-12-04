package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

type elfSleep struct {
	total              int
	s                  [60]int
	highestMinuteIndex int
}

func main() {
	fmt.Println("Advent of Code - Day 4 - Go")

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var ID, minute, hour int
	counter := make(map[int]elfSleep)
	input := strings.Split(string(content), "\n")
	sort.Strings(input)
	for _, row := range input {

		switch row[19:24] {
		case "Guard":
			fmt.Sscanf(row[19:], "Guard #%d begins shift", &ID)
		case "falls":
			fmt.Sscanf(row[12:17], "%d:%d", &hour, &minute)
		case "wakes":
			prevMinute := minute
			fmt.Sscanf(row[12:17], "%d:%d", &hour, &minute)
			elf := counter[ID]
			elf.total += minute - prevMinute
			for i := prevMinute; i < minute; i++ {
				elf.s[i]++
			}
			counter[ID] = elf
		}
	}

	var max, maxID, highestMinute, highestMinuteID int
	max = -1
	for elfID, elf := range counter {
		if elf.total > max {
			max = elf.total
			maxID = elfID
		}

		// Max minute for elf
		var highestIndividualMinute int
		for m, v := range elf.s {
			if v > highestIndividualMinute {
				highestIndividualMinute = v

				// Update max value for this particular elf
				elf.highestMinuteIndex = m
				counter[elfID] = elf

				if v > highestMinute {
					highestMinute = v
					highestMinuteID = elfID
				}
			}
		}
	}
	fmt.Println("Strategy 1", maxID, counter[maxID].highestMinuteIndex, maxID*counter[maxID].highestMinuteIndex)
	fmt.Println("Strategy 2", highestMinuteID, counter[highestMinuteID].highestMinuteIndex, highestMinuteID*counter[highestMinuteID].highestMinuteIndex)

	// [00] Strategy 1 641 41 26281
	// [00] Strategy 2 1973 37 73001

}
