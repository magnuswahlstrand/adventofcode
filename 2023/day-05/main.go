package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Map struct {
	To, From, Range int
}

var testInput = "seeds: 79 14 55 13\n\nseed-to-soil map:\n50 98 2\n52 50 48\n\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\n\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\n\nwater-to-light map:\n88 18 7\n18 25 70\n\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\n\ntemperature-to-humidity map:\n0 69 1\n1 0 69\n\nhumidity-to-location map:\n60 56 37\n56 93 4\n"

func part1Tests() {

	minVal := minLocationVal(testInput)

	fmt.Println("min value", minVal)
}

func part2Tests() {
	minVal := minLocationVal2(testInput)
	fmt.Println("min value", minVal)
}

func minLocationVal(input string) int {
	steps := strings.Split(input, "\n\n")
	source := toSeedNumbers(steps[0])
	for i, step := range steps[1:] {
		fmt.Println("step", i, step)

		// Find all the maps
		maps := parseMaps(step)

		for i, s := range source {
			for _, m := range maps {
				if s >= m.From && s < m.From+m.Range {
					source[i] = s - m.From + m.To
					break
				}
			}
		}
	}
	return slices.Min(source)
}

func parseMaps(step string) []Map {
	var maps []Map
	for _, r := range strings.Split(step, "\n")[1:] {
		m := Map{}
		_, _ = fmt.Sscanf(r, "%d %d %d", &m.To, &m.From, &m.Range)
		maps = append(maps, m)
	}
	return maps
}
func parseMaps2(step string) []segmentMap {
	var maps []segmentMap
	for _, r := range strings.Split(step, "\n")[1:] {
		m := Map{}
		_, _ = fmt.Sscanf(r, "%d %d %d", &m.To, &m.From, &m.Range)

		maps = append(maps, segmentMap{
			segment: segment{
				start: m.From,
				end:   m.Range,
			},
			offset: m.To - m.From,
		})
	}
	return maps
}

func minLocationVal2(input string) int {
	steps := strings.Split(input, "\n\n")
	source := toSeedNumbers2(steps[0])
	for i, step := range steps[1:] {
		fmt.Println("step", i, step)
		maps := parseMaps2(step)

		// Loop over source, and apply the maps, potentially creating new segments
		newSegment := []segment{}
		for len(source) > 0 {
			moved, newSegments := applyMaps(source[0], maps)
			source = append(source[1:], newSegments...)
			newSegment = append(newSegment, moved...)
		}

	}
	return 1
}

func applyMaps(a segment, maps []segmentMap) ([]segment, []segment) {

	for _, b := range maps {
		if a.overlaps(b.segment) {
			switch {
			case a.inside(b.segment):
				// Move the whole thing, no new segments
				return []segment{{a.start + b.offset, a.end}}, []segment{}
			case b.inside(a):
				// Create 3 segments, move middle one
				c1 := a.move(b.offset)
				c2 := segment{start:}.move(b.offset)
				c3 := a.move(b.offset)

				return []segment{
						{a.start, b.start - a.start},
					}, []segment{
						{b.start + b.offset, a.end() - b.end()},
					}

			case a.startsBefore(b.segment):
				// Create two segments. Keep the first, move the second
				return []segment{{a.start, b.start - a.start}},
					[]segment{{b.start + b.offset, a.end() - b.end()}}

			case b.startsBefore(a):
				// Create two segments. Move the first, keep the second
				return []segment{{a.start + b.offset, b.start - a.start}},
					[]segment{{b.start + b.offset, a.end() - b.end()}}
			}
		}
	}
	// Segment is unchanged
	return []segment{a}, []segment{}
}

type segment struct {
	start, end int
}

func (a segment) startsBefore(b segment) bool {
	return a.start < b.start
}

func (a segment) overlaps(b segment) bool {
	return a.start < b.end && b.start < a.end
}

func (a segment) EndsBefore(b segment) bool {
	return a.start+a.end < b.start
}

func (a segment) inside(b segment) bool {
	return a.start >= b.start && a.end <= b.end
}

func (a segment) move(offset int) segment {
	return segment{a.start + offset, a.end}
}

type segmentMap struct {
	segment
	offset int
}

func toSeedNumbers2(s string) []segment {
	var numbers []segment
	seeds := strings.Fields(s)[1:]
	for i := 0; i < len(seeds); i += 2 {
		start, _ := strconv.Atoi(seeds[i])
		width, _ := strconv.Atoi(seeds[i+1])
		numbers = append(numbers, segment{start, width})
	}
	return numbers
}

func toSeedNumbers(s string) []int {
	var numbers []int
	for _, n := range strings.Fields(s)[1:] {
		val, _ := strconv.Atoi(n)
		numbers = append(numbers, val)
	}
	return numbers
}

func part1(input []byte) {
	val := minLocationVal(string(input))
	fmt.Printf("answer to part 1 is %d\n", val)
}

func part2(input []byte) {
	fmt.Println("SEED2", input)
	val := minLocationVal2(string(input))
	fmt.Printf("answer to part 2 is %d\n", val)
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
