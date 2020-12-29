package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

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
	min, minBus := calcAnswer(string(input))
	fmt.Println("answer part 1:", min*minBus)
}

func part2(input []byte) {
	val := findTimestamp(string(input))
	fmt.Println("answer part 2:", val)
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return i
}

func mustInt64(s string) int64 {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return int64(i)
}

func calcAnswer(input string) (int, int) {
	lines := strings.SplitN(input, "\n", 2)
	minutes := mustInt(lines[0])

	min := math.MaxInt32
	var minBus int
	for _, v := range strings.Split(lines[1], ",") {
		if v == "x" {
			continue
		}
		bus := mustInt(v)
		busMinutes := bus - minutes%bus
		if busMinutes < min {
			min = busMinutes
			minBus = bus
		}
	}
	return min, minBus
}

func findTimestamp(input string) int64 {
	lines := strings.SplitN(input, "\n", 2)

	// Find buses and their offset
	var buses [][]int64
	for i, v := range strings.Split(lines[1], ",") {
		if v == "x" {
			continue
		}
		buses = append(buses, []int64{mustInt64(v), int64(i)})
	}

	var factor int64 = 1
	var timestamp int64
	for _, bus := range buses {
		for (timestamp+bus[1])%bus[0] != 0 {
			timestamp += factor
		}
		factor *= bus[0]
	}

	fmt.Println(timestamp)
	return timestamp
}

// Too high:
// 1556590287177056
// 3200966845281517
// 640856202464541
