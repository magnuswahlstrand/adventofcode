package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func reactionResults(polys string) string {
	pLen := len(polys)
	reactingTypes := 'a' - 'A'
	var sum rune
	for i := 0; i < pLen-1; i++ {
		if polys[i] > polys[i+1] {
			sum = rune(polys[i] - polys[i+1])
		} else {
			sum = rune(polys[i+1] - polys[i])
		}

		if sum == reactingTypes {
			// Remove reacting types and step back (unless at )
			polys = polys[:i] + polys[i+2:]
			i -= 2
			if i < 0 {
				i = -1
			}
			pLen -= 2
		}
	}

	return polys
}

func removeIngredient(poly string, r rune) string {
	rLower := string(r)
	rUpper := string(r + 'A' - 'a')
	return strings.Replace(strings.Replace(poly, rUpper, "", -1), rLower, "", -1)
}

func main() {
	fmt.Println("Advent of Code - Day 4 - Go")

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Part 1 - took 29 min - 9202
	result := reactionResults(string(content))
	fmt.Println("Part 1 - final polymer has length ", len(result))

	// Part 2 - took 12 min - 6394
	min := 1000000
	var minUnit rune
	for _, r := range "abcdefghijklmnopqrstuvwxyz" {

		res := reactionResults(removeIngredient(string(content), r))

		if len(res) < min {
			min = len(res)
			minUnit = r
		}
	}
	fmt.Println("Part 2 - removing", string(minUnit), "gives the a shortest polymer of length", min)
}
