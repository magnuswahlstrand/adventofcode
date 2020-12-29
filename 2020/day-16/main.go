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
	sum := part1ErrorRate(input)
	fmt.Println("answer part 1:", sum)
}

func part2(input []byte) {
	// Rule --> index mapping from printout :)
	// Rule:   0   1   2   3   4  5   6  7  8  9   10  11  12  13 14  15  16  17 18 19
	//                     x          x  x     x                  x       x
	// Ticket: 127,109,139,113,67,137,71,97,53,103,163,167,131,83,157,101,107,79,73,89
	fmt.Println("answer part 2:", 113*71*97*103*157*107)
}

type rangePair struct {
	name          string
	first, second [2]int
}

func (p *rangePair) valInRange(i int) bool {
	return (i >= p.first[0] && i <= p.first[1]) || (i >= p.second[0] && i <= p.second[1])
}

func parseRange(line string) rangePair {
	pair := rangePair{
		first:  [2]int{},
		second: [2]int{},
	}
	// Quick fix
	vals := strings.Split(line, ": ")
	pair.name = vals[0]
	fields := strings.Split(vals[1], " or ")
	_, err := fmt.Sscanf(fields[0], "%d-%d", &pair.first[0], &pair.first[1])
	if err != nil {
		log.Fatal("failed first scan", err)
	}
	_, err = fmt.Sscanf(fields[1], "%d-%d", &pair.second[0], &pair.second[1])
	if err != nil {
		log.Fatal("failed first scan", err)
	}
	return pair
}

func part1ErrorRate(input []byte) int {
	var rules []rangePair
	lines := strings.Split(string(input), "\n")
	var firstSectionEnd int
	for i, line := range lines {
		if line == "" {
			firstSectionEnd = i
			break
		}
		rules = append(rules, parseRange(line))
	}

	var validTickets [][]int

	var errorRate int
	for _, line := range lines[firstSectionEnd:] {
		if line == "" || strings.HasPrefix(line, "your") || strings.HasPrefix(line, "nearby") {
			continue
		}

		var ticket []int
		for _, s := range strings.Split(line, ",") {
			// Parse number
			number := mustInt(s)
			ticket = append(ticket, number)
		}

		ticketValid := true
		for _, number := range ticket {
			// Check number against rules
			var matchFound bool
			for _, rule := range rules {
				if rule.valInRange(number) {
					// Number is in a range!
					matchFound = true
					break
				}
			}

			if !matchFound {
				ticketValid = false
				errorRate += number
			}
		}

		if ticketValid {
			validTickets = append(validTickets, ticket)
		}
	}

	// Find indexes that could match rules
	possbileRuleIndexes := make([]map[int]bool, len(rules))
	for i, rule := range rules {
		rMap := map[int]bool{}

		// Iterate over tickets, column by column
		for fi := 0; fi < len(validTickets[0]); fi++ {
			isValid := true

			// Iterate over all tickets index by index
			for ti := 0; ti < len(validTickets); ti++ {
				// If a single number in this column breaks the rule,
				// we know this rule cannot be applied to this field
				if !rule.valInRange(validTickets[ti][fi]) {
					isValid = false
					break
				}
			}

			if isValid {
				rMap[fi] = true
			}
		}

		possbileRuleIndexes[i] = rMap
	}

	ruleIndexes := map[int]int{}

	// Run until we have found all rules
	for len(ruleIndexes) != len(rules) {
		fmt.Println(len(ruleIndexes))
		// Rules which only have one possible index
		for ri, _ := range rules {
			if len(possbileRuleIndexes[ri]) != 1 {
				continue
			}

			// Get the i from map
			var index int
			for i := range possbileRuleIndexes[ri] {
				index = i
				break
			}

			ruleIndexes[ri] = index
			for _, mp := range possbileRuleIndexes {
				delete(mp, index)
				//possbileRuleIndexes[i] = mp
			}

		}
	}

	// One final quick fix :D
	for i, rule := range rules {
		if strings.HasPrefix(rule.name, "departure") {
			fmt.Println(rule.name, ruleIndexes[i])
		}
	}
	return errorRate
}
