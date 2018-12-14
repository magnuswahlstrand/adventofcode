package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func isNice(in string) bool {
	var hasDoubles, hasSpecial bool
	var vowels int

	for i, r := range in {

		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		}

		if i == len(in)-1 {
			continue
		}

		if in[i] == in[i+1] {
			hasDoubles = true
		}

		check := in[i : i+2]
		switch check {
		case "ab", "cd", "pq", "xy":
			hasSpecial = true
			break

		}
	}

	return (vowels >= 3 && hasDoubles && !hasSpecial)
}

func isNew(used []int, a, b int) bool {
	new := true
	for _, v := range used {
		if v == a || v == b {
			// Value already used
			new = false
			break
		}
	}

	return new
}

func isNiceR(in string) bool {
	var hasDoubles bool
	counter := make(map[string][]int)

	for i := 0; i < len(in)-1; i++ {
		part := in[i : i+2]

		if isNew(counter[part], i, i+1) {
			counter[part] = append(counter[part], []int{i, i + 1}...)
		}

		if i < len(in)-2 && in[i] == in[i+2] {
			hasDoubles = true
		}
	}

	var hasDoublePairs bool
	for _, v := range counter {
		if len(v)/2 > 1 {
			hasDoublePairs = true
			break
		}
	}

	return (hasDoublePairs && hasDoubles)
}

func main() {
	fmt.Println("Day 5 - 2015")
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// runTests()

	var nice int
	for _, row := range strings.Split(string(content), "\n") {
		if isNice(row) {
			nice++
		}
	}
	fmt.Println("Part 1: Number of nice words:", nice)

	nice = 0
	for _, row := range strings.Split(string(content), "\n") {
		if isNiceR(row) {
			nice++
		}
	}
	fmt.Println("Part 2: Number of nice words:", nice)

}

func runTests() {
	fmt.Println("Q: ugknbfddgicrmopn is nice? A:", isNice("ugknbfddgicrmopn") == true)
	fmt.Println("Q: aaa is nice? A:", isNice("aaa") == true)
	fmt.Println("Q: jchzalrnumimnmhp is naughty? A:", isNice("jchzalrnumimnmhp") == false)
	fmt.Println("Q: haegwjzuvuyypxyu is naughty? A:", isNice("haegwjzuvuyypxyu") == false)
	fmt.Println("Q: dvszwmarrgswjxmb is naughty? A:", isNice("dvszwmarrgswjxmb") == false)

	fmt.Println("Q: qjhvhtzxzqqjkmpb is nicer? A:", isNiceR("qjhvhtzxzqqjkmpb") == true)
	fmt.Println("Q: xxyxx is nicer? A:", isNiceR("xxyxx") == true)
	fmt.Println("Q: aaa is naughty? A:", isNiceR("aaa") == false)
	fmt.Println("Q: uurcxstgmygtbstg is naughty? A:", isNiceR("uurcxstgmygtbstg") == false)
	fmt.Println("Q: ieodomkazucvgmuy is naughty? A:", isNiceR("ieodomkazucvgmuy") == false)
}
