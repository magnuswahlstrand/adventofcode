package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Point struct {
	x, y     int
	infinite bool
}

func getPoints(rows []string) []Point {
	points := []Point{}
	for _, row := range rows {
		p := Point{}
		fmt.Sscanf(row, "%d, %d", &p.x, &p.y)
		points = append(points, p)
	}
	return points
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func grid(points []Point, w, h, limit int, draw bool) (int, int) {
	grid := make([][]rune, h)
	grid2 := make([][]int, h)

	for i := 0; i < h; i++ {
		grid[i] = make([]rune, w)
		grid2[i] = make([]int, w)

		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = ' '
		}
	}

	close := 0
	counter := make(map[int]int)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {

			// Find closest point
			min := 100000
			var owner int
			for i, p := range points {

				dist := abs(p.x-x) + abs(p.y-y)

				// Add distance to grid two
				grid2[y][x] += dist

				// Two share the same distance
				if dist == min {
					owner = -1
					continue
				}

				if dist < min {
					min = dist
					owner = i
				}
			}

			if grid2[y][x] < limit {
				close++
			}

			if owner == -1 {
				grid[y][x] = '.'
				continue
			}

			grid[y][x] = 'a' + rune(owner)
			counter[owner]++

			// Mark potential area as infinite
			if x == 0 || x == w-1 || y == 0 || y == h-1 {
				points[owner].infinite = true
			}
		}
	}

	// Find largest non-infinite area
	max := -1
	for i, p := range points {
		if counter[i] > max && !p.infinite {
			max = counter[i]
		}
	}

	if draw {

		// Draw grid
		for _, row := range grid {
			fmt.Println(string(row))
		}
	}
	return max, close
}

func main() {
	fmt.Println("Advent of Code - Day 4 - Go")

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// contentTest := `1, 1
	// 1, 6
	// 8, 3
	// 3, 4
	// 5, 5
	// 8, 9`
	// rows := strings.Split(contentTest, "\n")
	// points := getPoints(rows)
	// max1, max2 := grid(points, 10, 10, 32, true)
	// fmt.Printf("Test 1 - the largest area is %d\n", max1)
	// fmt.Printf("Test 2 - the largest area is %d\n", max2)

	rows := strings.Split(string(content), "\n")
	points := getPoints(rows)
	max1, max2 := grid(points, 400, 400, 10000, false)
	fmt.Printf("Part 1 - the largest area is %d\n", max1) // 4475
	fmt.Printf("Part 2 - the largest area is %d\n", max2) // 35237
}
