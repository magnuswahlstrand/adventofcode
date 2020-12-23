package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func findTwoEntriesThatSumsTo(entries []int64, targetSum int64) (int64, int64) {
	for i, m := range entries {
		for _, n := range entries[i+1:] {
			if m+n == targetSum {
				return m, n
			}
		}
	}
	log.Fatal("no match found")
	return 0, 0
}

func findThreeEntriesThatSumsTo(entries []int64, targetSum int64) (int64, int64, int64) {
	for i, m := range entries {
		for j, n := range entries {
			if i == j {
				continue
			}

			for k, o := range entries[j+1:] {
				if k == i || k == j {
					continue
				}

				if m+n+o == targetSum {
					return m, n, o
				}
			}

		}
	}
	log.Fatal("no match found")
	return 0, 0, 0
}

func part1Tests() {
	m, n := findTwoEntriesThatSumsTo([]int64{1721, 979, 366, 299, 675, 1456}, 2020)
	if m*n != 514579 {
		log.Fatal("not the same!")
	}
}

func part2Tests() {
	m, n, o := findThreeEntriesThatSumsTo([]int64{1721, 979, 366, 299, 675, 1456}, 2020)
	if m*n*o != 241861950 {
		log.Fatal("not the same!")
	}
}

func parseEntries(input []byte) []int64 {
	var entries []int64

	rows := bytes.Split(input, []byte{'\n'})
	for _, r := range rows {
		m, err := strconv.ParseInt(string(r), 10, 64)
		if err != nil {
			log.Fatal("failed to convert", err)
		}

		entries = append(entries, m)
	}
	return entries
}

func part1(entries []int64) {
	m, n := findTwoEntriesThatSumsTo(entries, 2020)
	fmt.Printf("product for part 1 is %d x %d = %d\n", m, n, m*n)
}

func part2(entries []int64) {
	m, n, o := findThreeEntriesThatSumsTo(entries, 2020)
	fmt.Printf("product for part 2 is %d x %d x %d = %d\n", m, n, o, m*n*o)
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	entries := parseEntries(input)

	part1Tests()
	part1(entries)

	part2Tests()
	part2(entries)

	log.Println("success")
}