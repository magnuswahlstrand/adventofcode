package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func checksum(ids []string) int {
	var countTwo, countThree int
	for _, id := range ids {
		charMap := make(map[rune]int)
		for _, c := range id {
			charMap[c]++
		}
		var hasTwo, hasThree bool
		for _, v := range charMap {
			switch v {
			case 2:
				hasTwo = true
			case 3:
				hasThree = true
			default:
			}
		}

		if hasTwo {
			countTwo++
		}
		if hasThree {
			countThree++
		}
	}

	checkSum := countTwo * countThree
	return checkSum
}

func commonLettersOfCorrectBoxes(ids []string) string {

	// Compare all IDs to all other IDs
	for i := 0; i < len(ids); i++ {
		id1 := ids[i]

		for j := i + 1; j < len(ids); j++ {
			id2 := ids[j]

			//Compare ID 1 to ID 2
			commonChars := ""
			for k := 0; k < len(id1); k++ {
				if id1[k] == id2[k] {
					commonChars += string(id1[k])
				}
			}

			// IDs differ by only 1 character
			if len(commonChars) == len(id1)-1 {
				return commonChars
			}
		}
	}
	return ""
}

func printIdenticalParts(id1, id2 string) string {
	s := ""
	for k := 0; k < len(id1); k++ {
		if id1[k] == id2[k] {
			s += string(id1[k])
		}
	}
	return s
}

func main() {
	fmt.Println("Advent of Code - Day 2 - Go")

	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	ids := strings.Split(string(bytes), "\n")

	// Print checksum for part 1
	fmt.Println("Part 1")
	fmt.Println("- Checksum is:", checksum(ids))

	fmt.Println("Part 2")
	fmt.Println("- Common characters of ID is:", commonLettersOfCorrectBoxes(ids))
}
