package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func parseRules(input string) map[string]map[string]bool {
	rules := map[string]map[string]bool{}
	for _, line := range strings.Split(input, "\n") {
		outerColor, innerColors := parseColors(line)
		for _, innerColor := range innerColors {
			colorMap, ok := rules[innerColor]
			if !ok {
				// Init map if it doesn't exist
				colorMap = map[string]bool{}
			}
			colorMap[outerColor] = true

			// Save mapping
			rules[innerColor] = colorMap
		}
	}
	return rules
}

func parseRules2(input string) map[string]map[string]int {
	rules := map[string]map[string]int{}
	for _, line := range strings.Split(input, "\n") {
		outerColor, innerColors := parseColors2(line)
		colorMap := map[string]int{}
		for _, innerColor := range innerColors {
			colorMap[innerColor.colorname] = innerColor.count
		}
		rules[outerColor] = colorMap
	}
	return rules
}

type bags struct {
	colorname string
	count     int
}

func parseColors2(line string) (string, []bags) {
	split := strings.SplitN(line[:len(line)-1], " contain ", 2)

	// Get color from outer bag "6 dotted black bags"
	outerColor := strings.Join(strings.Fields(split[0])[0:2], " ")
	if split[1] == "no other bags" {
		return outerColor, []bags{}
	}

	var innerColors []bags
	for _, inner := range strings.Split(split[1], ", ") {
		// Get color from string "6 dotted black bags"
		fields := strings.Fields(inner)
		color := strings.Join(fields[1:3], " ")
		count, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal(err)
		}
		innerColors = append(innerColors, bags{
			colorname: color,
			count:     count,
		})
	}
	return outerColor, innerColors
}

func parseColors(line string) (string, []string) {
	split := strings.SplitN(line[:len(line)-1], " contain ", 2)

	// Get color from outer bag "6 dotted black bags"
	outerColor := strings.Join(strings.Fields(split[0])[0:2], " ")
	if split[1] == "no other bags" {
		return outerColor, []string{}
	}

	var innerColors []string
	for _, inner := range strings.Split(split[1], ", ") {
		// Get color from string "6 dotted black bags"
		color := strings.Join(strings.Fields(inner)[1:3], " ")
		innerColors = append(innerColors, color)
	}
	return outerColor, innerColors
}

func findColorsThatLeadToColor(color string, colors map[string]bool, rules map[string]map[string]bool) {
	for outsideColor := range rules[color] {
		colors[outsideColor] = true
		findColorsThatLeadToColor(outsideColor, colors, rules)
	}
}

func part1(input string) {
	rules := parseRules(input)
	colors := map[string]bool{}
	findColorsThatLeadToColor("shiny gold", colors, rules)

	fmt.Println("total # colors that lead to shiny gold for part 1:", len(colors))
}

func countBags(color string, rules map[string]map[string]int) int {
	total := 1
	for outsideColor, val := range rules[color] {
		total += val * countBags(outsideColor, rules)
	}
	return total
}

func part2(input string) {
	rules := parseRules2(input)
	n := countBags("shiny gold", rules) - 1
	fmt.Println("total # bags inside shiny gold bag for part 2:", n)
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1(string(input))
	part2(string(input))

	log.Println("success")
}
