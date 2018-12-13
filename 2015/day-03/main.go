package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Santa struct {
	x, y int
}

func countPresents(input string, nSantas int) int {

	var santas []*Santa
	for n := 0; n < nSantas; n++ {
		santas = append(santas, &Santa{})
	}

	locations := make(map[[2]int]int)
	locations[[2]int{0, 0}] += nSantas

	var active *Santa
	for i, r := range input {

		active = santas[i%nSantas]

		switch r {
		case '^':
			active.y--
		case 'v':
			active.y++
		case '<':
			active.x--
		case '>':
			active.x++
		}
		locations[[2]int{active.x, active.y}]++
	}

	var presents int
	for range locations {
		presents++
	}
	return presents
}

func main() {

	fmt.Println(countPresents(">", 1) == 2)
	fmt.Println(countPresents("^>v<", 1) == 4)
	fmt.Println(countPresents("^v^v^v^v^v", 1) == 2)

	fmt.Println(countPresents("^v", 2) == 3)
	fmt.Println(countPresents("^>v<", 2) == 3)
	fmt.Println(countPresents("^v^v^v^v^v", 2) == 11)

	input, _ := ioutil.ReadAll(os.Stdin)
	fmt.Println("Part 1:", countPresents(string(input), 1))
	fmt.Println("Part 2:", countPresents(string(input), 2))
}
