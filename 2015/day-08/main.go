package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func parseLine(input string) (int, int, int) {
	var total, memory int
	total = len(input)

	for i := 1; i < total-1; i++ {

		if input[i] == '\\' {

			if (input[i+1]) == 'x' {
				i += 2
			}
			i++
		}
		memory++
	}

	return total, memory, encode(input)
}

func encode(input string) int {
	return len("\"" + strings.Replace(strings.Replace(input, "\\", "\\\\", -1), "\"", "\\\"", -1) + "\"")
}

func runTestcases() {

	in := "\"\""
	total, memory, encoded := parseLine(in)
	fmt.Printf("input: % 10s --> total =% 3d memory =% 3d encoded=% 3d  ", in, total, memory, encoded)
	fmt.Println(total == 2, memory == 0)

	in = "\"abc\""
	total, memory, encoded = parseLine(in)
	fmt.Printf("input: % 10s --> total =% 3d memory =% 3d encoded=% 3d  ", in, total, memory, encoded)
	fmt.Println(total == 5, memory == 3)

	in = "\"aaa\\\"aaa\""
	total, memory, encoded = parseLine(in)
	fmt.Printf("input: % 10s --> total =% 3d memory =% 3d encoded=% 3d  ", in, total, memory, encoded)
	fmt.Println(total == 10, memory == 7)

	in = "\"\\x27\""
	total, memory, encoded = parseLine(in)
	fmt.Printf("input: % 10s --> total =% 3d memory =% 3d encoded=% 3d  ", in, total, memory, encoded)
	fmt.Println(total == 6, memory == 1)

}

func runReal() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	operations := strings.Split(string(content), "\n")

	var total, memory, encoded int
	for _, row := range operations {
		t, m, e := parseLine(row)
		total += t
		memory += m
		encoded += e
	}
	fmt.Printf("total=%d, memory=%d, encoded=%d, answer=%d", total, memory, encoded, encoded-total)
}

func main() {
	fmt.Println("Day 8 - 2015")
	runTestcases()
	runReal()
}
