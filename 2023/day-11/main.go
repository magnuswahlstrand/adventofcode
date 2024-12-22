package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Point struct {
	x int
	y int
}

func (p Point) distance(p2 Point) int {
	x := p.x - p2.x
	y := p.y - p2.y
	return max(x, -x) + max(y, -y)
}

func part1Tests() {
	input := "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#....."

	sum := sumDistances(input, 2)
	fmt.Println(sum)
}

func part2Tests() {
	input := "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#....."

	sum := sumDistances(input, 100)
	fmt.Println(sum)
}

func sumDistances(input string, multiplier int) int {
	world := strings.Split(input, "\n")
	var emptyRows []int
	var emptyCols []int
	var galaxies []Point
	for i := 0; i < len(world); i++ {

		// Rows
		found := false
		for j := 0; j < len(world[0]); j++ {
			if world[i][j] == '#' {
				found = true

				// Add expanded rows, we already know about
				galaxies = append(galaxies, Point{j, i + len(emptyRows)*(multiplier-1)})
			}
		}
		if !found {
			emptyRows = append(emptyRows, i)
		}

		// Cols
		found = false
		for j := 0; j < len(world[0]); j++ {
			if world[j][i] == '#' {
				found = true
			}
		}
		if !found {
			emptyCols = append(emptyCols, i)
		}
	}

	// Sort galaxies left to right
	slices.SortFunc(galaxies, func(a, b Point) int {
		return a.x - b.x
	})

	var currentOffset int
	for i := range galaxies {
		// First step empty cols forward until it is greater than galaxy.x

		for currentOffset < len(emptyCols) && emptyCols[currentOffset] < galaxies[i].x {
			currentOffset++
		}
		galaxies[i].x += currentOffset * (multiplier - 1)
	}

	var sum int
	for i := range galaxies {
		for j := range galaxies[i+1:] {
			sum += galaxies[i].distance(galaxies[j+i+1])
		}
	}

	return sum
}

func part1(input []byte) {
	sum := sumDistances(string(input), 2)
	fmt.Printf("answer to part 1 is %d\n", sum)
}

func part2(input []byte) {
	sum := sumDistances(string(input), 1000000)
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
