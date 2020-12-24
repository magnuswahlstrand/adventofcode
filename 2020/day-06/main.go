package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func countGroupsAndYes(input string) (int, int) {
	var totalYes int
	groups := strings.Split(input, "\n\n")
	for _, group := range groups {
		counter := make(map[rune]int)

		chars := strings.Join(strings.Fields(group), "")
		for _, c := range chars {
			counter[c]++
		}

		totalYes += len(counter)
	}
	return len(groups), totalYes
}

func countGroupsAndYes2(input string) (int, int) {
	var totalYes int
	groups := strings.Split(input, "\n\n")
	for _, group := range groups {
		counter := make(map[rune]int)

		// Count occurrences of characters
		groupAnswers := strings.Split(group, "\n")
		for _, member := range groupAnswers {
			for _, c := range member {
				counter[c]++
			}
		}

		// Count number of characters where everyone in the group
		nGroupMembers := len(groupAnswers)
		for _, count := range counter {
			if count == nGroupMembers {
				totalYes++
			}
		}
	}
	return len(groups), totalYes
}

func part1(input string) {
	_, totalYes := countGroupsAndYes(input)
	fmt.Println("total # yes for part 1:", totalYes)
}
func part2(input string) {
	_, totalYes := countGroupsAndYes2(input)
	fmt.Println("total # yes for part 2:", totalYes)
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1(string(input))
	part2(string(input))

	log.Println("success")
}
