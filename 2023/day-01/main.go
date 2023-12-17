package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
)

func part1Tests() {
	input := []byte("1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet")

	sum := calculateSum(input)

	if sum != 142 {
		log.Fatal("not the same!")
	}
}

func calculateSum(input []byte) int {
	var first int
	var last int
	var sum int
	for _, line := range bytes.Split(input, []byte{'\n'}) {
		first, last = -1, -1
		for _, c := range line {
			if c >= '0' && c <= '9' {
				if first == -1 {
					first = int(c - '0')
				}
				last = int(c - '0')
			}
		}

		if first == -1 || last == -1 {
			panic("first or last is -1")
		}
		sum += first*10 + last
	}
	return sum
}

var matcher = regexp.MustCompile(`(zero|one|two|three|four|five|six|seven|eight|nine|[0-9])`)

var matchToValue = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func calculateSum2(input []byte) int {
	var first, last, sum int
	for _, line := range bytes.Split(input, []byte{'\n'}) {
		first, last = -1, -1
		for i := range line {
			if number := getNumber(i, line); number != -1 {
				if first == -1 {
					first = number
				}
				last = number
				fmt.Println("found", number)
			}
		}

		//fmt.Println(first, last)
		sum += first*10 + last
	}
	return sum
}

func getNumber(i int, line []byte) int {
	if line[i] >= '0' && line[i] <= '9' {
		return int(line[i] - '0')
	}

	//fmt.Println("foo", i, string(line[i:]))
	switch {
	case bytes.HasPrefix(line[i:], []byte("one")):
		return 1
	case bytes.HasPrefix(line[i:], []byte("two")):
		return 2
	case bytes.HasPrefix(line[i:], []byte("three")):
		return 3
	case bytes.HasPrefix(line[i:], []byte("four")):
		return 4
	case bytes.HasPrefix(line[i:], []byte("five")):
		return 5
	case bytes.HasPrefix(line[i:], []byte("six")):
		return 6
	case bytes.HasPrefix(line[i:], []byte("seven")):
		return 7
	case bytes.HasPrefix(line[i:], []byte("eight")):
		return 8
	case bytes.HasPrefix(line[i:], []byte("nine")):
		return 9
	default:
		return -1
	}
}

func part2Tests() {
	input := []byte("two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen")

	sum := calculateSum2(input)

	if sum != 281 {
		log.Fatal("not the same!")
	}
}

func part1(input []byte) {
	sum := calculateSum(input)
	fmt.Printf("answer to part 1 is %d\n", sum)
}
func part2(input []byte) {
	sum := calculateSum2(input)
	fmt.Printf("answer to part 2 is %d\n", sum)
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1Tests()
	part1(input)

	part2Tests()
	part2(input)

	log.Println("success")
}
