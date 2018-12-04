package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func circularSum(input string) int {

	var numbers []int
	for _, r := range input {
		numbers = append(numbers, int(r-'0'))
	}

	// Turn it circular
	numbers = append(numbers, numbers[0])

	sum := 0
	for i := 0; i < len(input); i++ {
		if numbers[i] == numbers[i+1] {
			sum += numbers[i]
		}
	}
	return sum
}

func halfSum(input string) int {

	var numbers []int
	for _, r := range input {
		numbers = append(numbers, int(r-'0'))
	}

	sum := 0
	halfLen := len(input) / 2
	for i := 0; i < len(input); i++ {
		if numbers[i] == numbers[(i+halfLen)%len(input)] {
			sum += numbers[i]
		}
	}
	return sum
}

func main() {

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Advent of Code - 2017 - Day 1 - Go")

	// fmt.Println("1122 -> ", circularSum("1122") == 3)
	// fmt.Println("1111 -> ", circularSum("1111") == 4)
	// fmt.Println("1234 -> ", circularSum("1234") == 0)
	// fmt.Println("91212129 -> ", circularSum("91212129") == 9)
	// Time taken: 11 min
	fmt.Println("input -> ", circularSum(string(input)))

	// fmt.Println("1212", halfSum("1212") == 6)
	// fmt.Println("1221", halfSum("1221") == 0)
	// fmt.Println("123425", halfSum("123425") == 4)
	// fmt.Println("123123", halfSum("123123") == 12)
	// fmt.Println("12131415", halfSum("12131415") == 4)

	// Time taken: 5 min
	fmt.Println("input -> ", halfSum(string(input)))

}
