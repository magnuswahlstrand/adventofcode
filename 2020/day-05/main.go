package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func parseRow(s string) int {
	binary := strings.Map(func(r rune) rune {
		switch r {
		case 'F':
			return '0'
		case 'B':
			return '1'
		default:
			log.Fatal("unexpected character", string(r))
			return ' '
		}
	}, s[:7])

	row, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(row)
}

func parseCols(s string) int {
	binary := strings.Map(func(r rune) rune {
		switch r {
		case 'R':
			return '1'
		case 'L':
			return '0'
		default:
			log.Fatal("unexpected character", string(r))
			return ' '
		}
	}, s[7:])

	row, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(row)
}

func parseSeatID(s string) int {
	binary := strings.Map(func(r rune) rune {
		switch r {
		case 'F':
			return '0'
		case 'B':
			return '1'
		case 'R':
			return '1'
		case 'L':
			return '0'
		default:
			log.Fatal("unexpected character", string(r))
			return ' '
		}
	}, s)
	row, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(row)
}

func part1(input string) {
	maxSeatID := math.MinInt64
	for _, line := range strings.Split(input, "\n") {
		seatID := parseSeatID(line)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	fmt.Println("highest seat ID for part 1:", maxSeatID)
}

func part2(input string) {
	lines := strings.Split(input, "\n")

	rows := map[int]int{}
	cols := map[int]int{}

	// Find the first and last row
	var firstRow, lastRow int
	for _, line := range lines {
		row := parseRow(line)
		if row < firstRow {
			firstRow = row
		}
		if row > lastRow {
			lastRow = row
		}
	}

	// Count seats in row and col for all rows except the first
	for _, line := range lines {
		row := parseRow(line)
		if row == firstRow || row == lastRow {
			continue
		}
		col := parseCols(line)
		rows[row]++
		cols[col]++
	}

	minN := math.MaxInt64
	var minRow int
	for row, n := range rows {
		if n < minN {
			minN = n
			minRow = row
		}
	}

	minN = math.MaxInt64
	var minCol int
	for col, n := range cols {
		if n < minN {
			minN = n
			minCol = col
		}
	}
	fmt.Printf("seat ID for part 2: %d * 8 + %d = %d\n", minRow, minCol, minRow*8+minCol)
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
