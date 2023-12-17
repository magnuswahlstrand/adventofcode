package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

type SetColor struct {
	Count int
	Color string
}

func part1Tests() {
	input := []byte("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green")
	sum := sumPossible(input)
	fmt.Println(sum)
}

func part2Tests() {
	input := []byte("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green")
	sum := sumMinimum(input)
	fmt.Println(sum)
}

func sumPossible(input []byte) int {
	var n int
	var clr string
	var sum int
	for gameNumber, game := range bytes.Split(input, []byte{'\n'}) {
		valid := true
		start := bytes.Index(game, []byte(":"))
		for _, set := range bytes.Split(game[start+1:], []byte{';'}) {
			var r, g, b int
			for _, color := range bytes.Split(set, []byte{','}) {
				fmt.Sscanf(string(color), " %d %s", &n, &clr)
				switch clr {
				case "red":
					r += n
				case "green":
					g += n
				case "blue":
					b += n
				default:
					panic("unknown color")
				}
			}
			if r > 12 || g > 13 || b > 14 {
				valid = false
				break
			}
		}

		if valid {
			sum += gameNumber + 1
		}

	}
	return sum
}

func sumMinimum(input []byte) int {
	var n int
	var clr string
	var sum int
	for _, game := range bytes.Split(input, []byte{'\n'}) {
		start := bytes.Index(game, []byte(":"))
		r, g, b := 0, 0, 0
		for _, set := range bytes.Split(game[start+1:], []byte{';'}) {
			for _, color := range bytes.Split(set, []byte{','}) {
				fmt.Sscanf(string(color), " %d %s", &n, &clr)
				switch clr {
				case "red":
					r = max(r, n)
				case "green":
					g = max(g, n)
				case "blue":
					b = max(b, n)
				default:
					panic("unknown color")
				}
			}
		}
		sum += r * g * b
	}
	return sum
}

func part1(input []byte) {
	sum := sumPossible(input)
	fmt.Printf("answer to part 1 is %d\n", sum)
}
func part2(input []byte) {
	sum := sumMinimum(input)
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
