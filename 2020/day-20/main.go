package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return i
}

func main() {
	t := time.Now()
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	part1(input)
	part2(input)

	log.Println("success in", time.Since(t))
}

func part1(input []byte) {
	_, tiles := parseTiles(input)
	product := findCornerProductAndMatches(tiles)
	fmt.Println("answer part 1:", product)
}

func part2(input []byte) {
	//fmt.Println("answer part 2:", sum)
}

type tile struct {
	id    int
	grid  []string
	sides []string
}

func (t *tile) side(i int) string {
	return t.sides[i]
}

func Reverse(in string) string {
	var sb strings.Builder
	runes := []rune(in)
	for i := len(runes) - 1; 0 <= i; i-- {
		sb.WriteRune(runes[i])
	}
	return sb.String()
}

type match struct {
	tile1Side, tile2Side int
	reversed             bool
}

func findCornerProductAndMatches(tiles []*tile) int {
	connectsTo := map[int]map[int]*match{}
	//connectsFrom := map[int]map[int]*tile{}

	cornerProduct := 1
	for i, t1 := range tiles {
		connectsTo[t1.id] = map[int]*match{}
		for t1s := 0; t1s < 4; t1s++ {

		innerloop:
			for j, t2 := range tiles {
				if i == j {
					continue
				}

				for t2s := 0; t2s < 4; t2s++ {
					straight := t1.side(t1s) == t2.side(t2s)
					reversed := t1.side(t1s) == Reverse(t2.side(t2s))
					if straight || reversed {
						connectsTo[t1.id][t2.id] = &match{
							tile1Side: t1s,
							tile2Side: t2s,
							reversed:  reversed,
						}
						break innerloop
					}
				}
			}
		}
		if len(connectsTo[t1.id]) == 2 {
			cornerProduct *= t1.id
		}
	}

	for _, t := range connectsTo {
		fmt.Println(t)
	}
	return cornerProduct
}

func parseTiles(input []byte) (map[int]*tile, []*tile) {
	tilesMap := map[int]*tile{}
	tiles := []*tile{}
	for _, t := range strings.Split(string(input), "\n\n") {
		rows := strings.Split(t, "\n")

		// 1. Parse id
		var tileID int
		_, err := fmt.Sscanf(rows[0], "Tile %d:", &tileID)
		if err != nil {
			log.Fatal("failed to parse title ID", err)
		}

		size := len(rows) - 1
		s0 := rows[1]
		s2 := rows[size]
		var s1, s3 string
		for i := 0; i < size; i++ {
			s1 += string(rows[i+1][size-1])
			s3 += string(rows[i+1][0])
		}

		t := &tile{
			id:   tileID,
			grid: rows[1:],
			sides: []string{
				s0,
				s1,
				s2,
				s3,
			},
		}
		tilesMap[tileID] = t
		tiles = append(tiles, t)
	}
	return tilesMap, tiles
}
