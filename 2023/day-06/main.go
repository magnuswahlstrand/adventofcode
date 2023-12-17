package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func toInts(s string) []int {
	var ints []int
	for _, s2 := range strings.Fields(s) {
		n, _ := strconv.Atoi(s2)
		ints = append(ints, n)
	}
	return ints
}

func part1Tests() {
	input := "Time:      7  15   30\nDistance:  9  40  200"
	total := totalMoE(input)
	fmt.Println("total", total)
}

func part2Tests() {
	input := "Time:      7  15   30\nDistance:  9  40  200"
	total := totalMoE2(input)
	fmt.Println("total", total)
}

func totalMoE(input string) int {
	ts, ds, _ := strings.Cut(input, "\n")
	times := toInts(ts[5:])
	distances := toInts(ds[9:])

	total := 1
	for i := range times {
		moe := marginOfError(times[i], distances[i])
		total *= moe
	}
	return total
}

func totalMoE2(input string) int {
	ts, ds, _ := strings.Cut(input, "\n")
	t, _ := strconv.Atoi(strings.ReplaceAll(ts[5:], " ", ""))
	d, _ := strconv.Atoi(strings.ReplaceAll(ds[9:], " ", ""))
	return marginOfError(t, d)
}

func marginOfError(tTot int, cMax int) int {
	pt1 := float64(tTot)
	pt2 := math.Sqrt(math.Pow(float64(tTot), 2) - 4*float64(cMax))

	b := int(math.Ceil((pt1+pt2)/2.0 - 1))
	a := int(math.Floor((pt1-pt2)/2.0 + 1))
	fmt.Println("tTot", tTot, "cMax", cMax)
	fmt.Println("a", a, "b", b, "p1", pt1, "p2", pt2)
	return b - a + 1
}

func part1(input []byte) {
	total := totalMoE(string(input))
	fmt.Printf("answer to part 1 is %d\n", total)
}
func part2(input []byte) {
	total := totalMoE2(string(input))
	fmt.Printf("answer to part 2 is %d\n", total)
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
