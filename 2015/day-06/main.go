package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gonum.org/v1/gonum/mat"
)

type Operation int

const (
	toggle Operation = 0
	off    Operation = 1
	on     Operation = 2
)

func operation(input string) (Operation, int, int, int, int) {
	var op Operation
	var start int
	var x1, x2, y1, y2 int
	switch input[0:7] {
	case "toggle ":
		op = toggle
		start = 7
	case "turn of":
		op = off
		start = 9
	case "turn on":
		op = on
		start = 8
	}

	fmt.Sscanf(input[start:], "%d,%d through %d,%d", &x1, &y1, &x2, &y2)

	return op, x1, x2, y1, y2
}

func lightSum(width, height int, instructions string) int {

	grid := mat.NewDense(width, height, nil)
	var x1, x2, y1, y2 int
	var op Operation
	for _, row := range strings.Split(instructions, "\n") {

		op, x1, x2, y1, y2 = operation(row)

		for y := y1; y <= y2; y++ {
			for x := x1; x <= x2; x++ {

				switch op {
				case toggle:
					grid.Set(x, y, 1-grid.At(x, y))
				case off:
					grid.Set(x, y, 0)
				case on:
					grid.Set(x, y, 1)

				}
			}
		}

		// fmt.Println(mat.Formatted(grid, mat.Prefix(""), mat.Squeeze()))

	}
	return int(mat.Sum(grid))
}

func lightSumV2(width, height int, instructions string) int {

	grid := mat.NewDense(width, height, nil)
	var x1, x2, y1, y2 int
	var op Operation
	for _, row := range strings.Split(instructions, "\n") {

		op, x1, x2, y1, y2 = operation(row)

		for y := y1; y <= y2; y++ {
			for x := x1; x <= x2; x++ {

				switch op {
				case toggle:
					grid.Set(x, y, grid.At(x, y)+2)
				case off:
					min := grid.At(x, y) - 1
					if min < 0 {
						min = 0
					}
					grid.Set(x, y, min)
				case on:
					grid.Set(x, y, grid.At(x, y)+1)

				}
			}
		}

		// fmt.Println(mat.Formatted(grid, mat.Prefix(""), mat.Squeeze()))

	}
	return int(mat.Sum(grid))
}

func runTestcases() {
	instructions := `turn on 0,0 through 2,2
	turn off 0,0 through 0,2
	toggle 2,1 through 2,1
	toggle 0,1 through 0,1`

	sumTest := lightSum(3, 3, instructions)
	fmt.Println("Solution test 1:", sumTest)

	sumTest2 := lightSumV2(3, 3, instructions)
	fmt.Println("Solution test 2:", sumTest2)

}

func main() {

	fmt.Println("Day 6 - 2015")

	// runTestcases()

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	instructions := string(content)

	sum := lightSum(1000, 1000, instructions)
	fmt.Println("Solution part 1:", sum)

	sum2 := lightSumV2(1000, 1000, instructions)
	fmt.Println("Solution part 2:", sum2)

	// Answered 17325717, it was wrong, forgot cap brightness at 0
	// Answered 17836115, it was correct
}
