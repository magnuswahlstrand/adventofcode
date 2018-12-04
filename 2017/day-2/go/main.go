package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func diff(input string) int {
	max := -1
	min := 10000

	for _, n := range getNumbers(input) {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	return max - min
}

func divisableDiff(input string) int {
	max := -1
	min := 10000
	numbers := getNumbers(input)
	for i := 0; i < len(numbers); i++ {
		n := numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			k := numbers[j]

			if n%k == 0 {
				return (n / k)
			}

			if k%n == 0 {
				return (k / n)
			}
		}
	}
	return max - min
}

func getNumbers(input string) []int {
	var out []int
	for _, s := range strings.Fields(input) {
		n, err := strconv.Atoi(s)
		if err != nil {
			//Should not happend
			log.Fatal(err)
		}

		out = append(out, n)
	}
	return out
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rows := strings.Split(string(input), "\n")

	fmt.Println("Advent of Code - 2017 - Day 2 - Go")

	// fmt.Println("5 1 9 5", diff("5 1 9 5") == 8)
	// fmt.Println("7 5 3", diff("7 5 3") == 4)
	// fmt.Println("2 4 6 8", diff("2 4 6 8") == 6)

	sum := 0
	for _, row := range rows {
		sum += diff(row)
	}
	// Time taken: 12 min
	fmt.Println("Result:", sum)

	// fmt.Println("5 9 2 8", divisableDiff("5 9 2 8") == 4)
	// fmt.Println("9 4 7 3", divisableDiff("9 4 7 3") == 3)
	// fmt.Println("3 8 6 5", divisableDiff("3 8 6 5") == 2)

	sum = 0
	for _, row := range rows {
		sum += divisableDiff(row)
	}
	// Time taken: 14 min
	fmt.Println("Result:", sum)

}
