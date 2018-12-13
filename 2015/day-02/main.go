package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func area(in string) int {
	var area int
	var l, w, h int
	fmt.Sscanf(in, "%dx%dx%d", &l, &w, &h)

	min := 100000
	for _, sideArea := range []int{l * w, w * h, h * l} {
		area += 2 * sideArea
		if sideArea < min {
			min = sideArea
		}
	}

	return area + min
}

func wrap(in string) int {
	var area int
	var l, w, h int
	fmt.Sscanf(in, "%dx%dx%d", &l, &w, &h)

	min := 100000
	wrap := 0
	for _, dim := range [][]int{[]int{l, w}, []int{w, h}, []int{h, l}} {
		sideArea := dim[0] * dim[1]
		area += 2 * sideArea

		if sideArea < min {
			min = sideArea
			wrap = 2*dim[0] + 2*dim[1]
		}
	}

	return wrap + l*w*h
}

func main() {

	fmt.Println("")
	fmt.Println("Advent of Code - 2015 - Day - 2")

	row := "2x3x4"
	fmt.Printf("%s=%d %s\n", row, area(row), strconv.FormatBool(area(row) == 58))
	fmt.Printf("%s=%d %s\n", row, wrap(row), strconv.FormatBool(wrap(row) == 34))

	row = "1x1x10"
	fmt.Printf("%s=%d %s\n", row, area(row), strconv.FormatBool(area(row) == 43))
	fmt.Printf("%s=%d %s\n", row, wrap(row), strconv.FormatBool(wrap(row) == 14))

	input, _ := ioutil.ReadAll(os.Stdin)
	var total, totalWrap int
	for _, row := range strings.Split(string(input), "\n") {
		total += area(row)
		totalWrap += wrap(row)
	}
	fmt.Println(total, totalWrap)
}
