package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const testWorld = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func mustTrue(t bool, txt string) {
	if !t {
		log.Fatal("condition failed: ", txt)
	}
}

type TreeWorld []string

func (w TreeWorld) at(x, y int) uint8 {
	return w[y][x%len(w[y])]
}

//func (w TreeWorld) countTrees() int {
//	var x, y, trees int
//	for range w {
//		if w.at(x, y) == '#' {
//			trees++
//		}
//
//		x += 3
//		y++
//	}
//	return trees
//}

func (w TreeWorld) countTreesWithSlope(r, d int) int {
	var x, y, trees int
	for y < len(w) {
		if w.at(x, y) == '#' {
			trees++
		}
		x += r
		y += d
	}
	return trees
}

func part1Tests() {
	tWorld := TreeWorld(strings.Split(testWorld, "\n"))
	mustTrue(tWorld.at(2, 0) == '#', "2,0 == #")
	mustTrue(tWorld.at(3, 0) == '#', "3,0 == #")
	mustTrue(tWorld.at(4, 0) != '#', "4,0 != #")
	mustTrue(tWorld.at(0, 1) == '#', "0,1 == #")
	mustTrue(tWorld.at(11, 1) == tWorld.at(0, 1), "")
	mustTrue(tWorld.at(12, 1) == tWorld.at(1, 1), "")

	trees := tWorld.countTreesWithSlope(3, 1)
	mustTrue(trees == 7, "# trees == 7")
}

func part2Tests() {
	tWorld := TreeWorld(strings.Split(testWorld, "\n"))
	mustTrue(tWorld.countTreesWithSlope(1, 1) == 2, "# trees == 2")
	mustTrue(tWorld.countTreesWithSlope(3, 1) == 7, "# trees == 7")
	mustTrue(tWorld.countTreesWithSlope(5, 1) == 3, "# trees == 3")
	mustTrue(tWorld.countTreesWithSlope(7, 1) == 4, "# trees == 4")
	mustTrue(tWorld.countTreesWithSlope(1, 2) == 2, "# trees == 2")
}

func part1(w TreeWorld) {
	trees := w.countTreesWithSlope(3, 1)
	fmt.Println("valid passwords for part 1:", trees)
}

func part2(w TreeWorld) {
	answer := w.countTreesWithSlope(1, 1) *
		w.countTreesWithSlope(3, 1) *
		w.countTreesWithSlope(5, 1) *
		w.countTreesWithSlope(7, 1) *
		w.countTreesWithSlope(1, 2)
	fmt.Println("valid passwords for part 2:", answer)
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	world := TreeWorld(strings.Split(string(input), "\n"))

	part1Tests()
	part1(world)

	part2Tests()
	part2(world)

	log.Println("success")
}
