package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func hasTwoThreeIdentical(s string) (bool, bool) {
	charMap := make(map[rune]int)
	for _, c := range s {
		charMap[c]++
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

func printCheckSum(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

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
	}

	checkSum := countTwo * countThree
	fmt.Printf("- The checksum for boxes with twos (%d) and threes (%d) is: %d\n", countTwo, countThree, checkSum)
	return nil
}

func findAlmostEquals(filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	ids := strings.Split(string(bytes), "\n")

	// Compare all IDs to all other IDs
	for i := 0; i < len(ids); i++ {
		id1 := ids[i]

		for j := i + 1; j < len(ids); j++ {
			id2 := ids[j]

			//Compare ID 1 to ID 2
			differencesFound := 0
			for k := 0; k < len(id1); k++ {

				if id1[k] != id2[k] {
					differencesFound++
				}
			}

			if differencesFound == 1 {
				fmt.Print("- First almost right id: ")
				printIdenticalParts(id1, id2)
				return nil
			}
		}
	}
	return nil
}

func printIdenticalParts(id1, id2 string) {
	for k := 0; k < len(id1); k++ {
		if id1[k] == id2[k] {
			fmt.Print(string(id1[k]))
		}
	}
}

func main() {
	fmt.Println("Advent of Code - Day 2 - Go")

	// Print checksum for part 1
	fmt.Println("Part 1")
	err := printCheckSum("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 2")
	err = findAlmostEquals("input.txt")
	if err != nil {
		log.Fatal(err)
	}

}
