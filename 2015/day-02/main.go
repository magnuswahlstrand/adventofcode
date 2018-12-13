package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func paper(in string) int {
	s := make([]int, 3)
	fmt.Sscanf(in, "%dx%dx%d", &s[0], &s[1], &s[2])
	sort.Ints(s)
	return 2*(s[0]*s[1]+s[1]*s[2]+s[0]*s[2]) + s[0]*s[1]
}

func ribbon(in string) int {
	s := make([]int, 3)
	fmt.Sscanf(in, "%dx%dx%d", &s[0], &s[1], &s[2])
	sort.Ints(s)
	return 2*(s[0]+s[1]) + s[0]*s[1]*s[2]
}

func main() {

	fmt.Println("")
	fmt.Println("Advent of Code - 2015 - Day - 2")

	row := "2x3x4"
	fmt.Printf("%s=%d %s\n", row, paper(row), strconv.FormatBool(paper(row) == 58))
	fmt.Printf("%s=%d %s\n", row, ribbon(row), strconv.FormatBool(ribbon(row) == 34))

	row = "1x1x10"
	fmt.Printf("%s=%d %s\n", row, paper(row), strconv.FormatBool(paper(row) == 43))
	fmt.Printf("%s=%d %s\n", row, ribbon(row), strconv.FormatBool(ribbon(row) == 14))

	input, _ := ioutil.ReadAll(os.Stdin)
	var total, totalRibbon int
	for _, row := range strings.Split(string(input), "\n") {
		total += paper(row)
		totalRibbon += ribbon(row)
	}
	fmt.Println(total, totalRibbon)
}
