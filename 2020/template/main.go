package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return i
}

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
	//fmt.Println("answer part 1:", sum)
}

func part2(input []byte) {
	//fmt.Println("answer part 2:", sum)
}
