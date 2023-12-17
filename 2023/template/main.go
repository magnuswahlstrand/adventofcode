package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func part1Tests() {
	input := []byte("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green")

	for _, line := range bytes.Split(input, []byte{'\n'}) {

	}
}

func part1(input []byte) {
	sum := calculateSum(input)
	fmt.Printf("answer to part 1 is %d\n", sum)
}

func main() {
	_, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1Tests()
	//part1(input)

	//part2Tests()
	//part2(input)

	log.Println("success")
}
