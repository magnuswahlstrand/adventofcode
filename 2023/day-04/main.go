package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func part1Tests() {
	input := []byte("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11")

	sum := calcCardValue(input)
	fmt.Println(sum)
}

func calcCardValue(input []byte) int {
	var sum int
	for _, card := range bytes.Split(input, []byte{'\n'}) {
		stringCard := string(card)
		_, rest, _ := strings.Cut(stringCard, ":")
		winners, actual, _ := strings.Cut(rest, "|")

		isWinner := map[string]bool{}
		for _, w := range strings.Fields(winners) {
			isWinner[w] = true
		}

		var cardValue int
		for _, a := range strings.Fields(actual) {
			if isWinner[a] {
				if cardValue == 0 {
					cardValue = 1
				} else {
					cardValue *= 2
				}
			}
		}

		//number, _ := strconv.ParseInt(strings.Repeat("1", count), 2, 64)
		sum += cardValue
	}
	return sum
}

func calcCardValue2(input []byte) int {
	numberOfScratchCards := map[int]int{}
	cards := bytes.Split(input, []byte{'\n'})

	var sum int
	for i, card := range cards {
		numberOfScratchCards[i] += 1

		stringCard := string(card)
		_, rest, _ := strings.Cut(stringCard, ":")
		winners, actual, _ := strings.Cut(rest, "|")

		isWinner := map[string]int{}
		for _, w := range strings.Fields(winners) {
			isWinner[w] = 1
		}

		var count int
		for _, a := range strings.Fields(actual) {
			count += isWinner[a]
		}

		for j := i + 1; j < i+1+count && j < len(cards); j++ {
			numberOfScratchCards[j] += numberOfScratchCards[i]
		}

		sum += numberOfScratchCards[i]
	}
	return sum
}

func part2Tests() {
	input := []byte("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11")

	sum := calcCardValue2(input)
	fmt.Println(sum)
}
func part1(input []byte) {
	sum := calcCardValue(input)
	fmt.Printf("answer to part 1 is %d\n", sum)
}

func part2(input []byte) {
	sum := calcCardValue2(input)
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
