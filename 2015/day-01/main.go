package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func floor(input []byte) int {
	var floor int
	for _, r := range input {
		switch r {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return floor
}

func firstBasement(input []byte) int {
	var floor int
	for step, r := range input {
		switch r {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor < 0 {
			return step + 1
		}
	}
	return -1
}

type testcase struct {
	in       []byte
	expected int
}

func main() {

	fmt.Println("")
	fmt.Println("Advent of Code - 2015 - Day - 1")

	tcs := []testcase{
		testcase{[]byte("(())"), 0},
		testcase{[]byte("()()"), 0},
		testcase{[]byte("((("), 3},
		testcase{[]byte("(()(()("), 3},
		testcase{[]byte("))((((("), 3},
		testcase{[]byte("())"), -1},
		testcase{[]byte("))("), -1},
		testcase{[]byte(")))"), -3},
		testcase{[]byte(")())())"), -3},
	}

	runTestcases := false
	if runTestcases {

		for _, tc := range tcs {
			fmt.Println(tc.in, "-->", tc.in, floor(tc.in) == tc.expected)
		}

		tcs = []testcase{
			testcase{[]byte(")"), 1},
			testcase{[]byte("()())"), 5},
		}
		for _, tc := range tcs {
			fmt.Println(tc.in, "-->", tc.in, firstBasement(tc.in) == tc.expected)
		}
	}

	input, _ := ioutil.ReadAll(os.Stdin)
	fmt.Println("Part 1 is", floor(input))         // 9 minutes
	fmt.Println("Part 2 is", firstBasement(input)) // 4 minutes
}
