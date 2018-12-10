package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code - Day 4 - Go")

	content, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rows := strings.Split(string(content), "\n")

	points := []point{}
	for _, row := range rows {
		p := point{}
		fmt.Sscanf(strings.Replace(row, " ", "", -1), "position=<%d,%d>velocity=<%d,%d>", &p.x, &p.y, &p.vx, &p.vy)
		points = append(points, p)
	}

	// Update
	width := 10000000
	var minX, minY, maxX, maxY int
	for r := 0; r < 10360; r++ {
		minX, minY = 100000, 1000000
		maxX, maxY = -100000, -100000
		for i, p := range points {
			points[i].x, points[i].y = p.x+p.vx, p.y+p.vy

			if p.x > maxX {
				maxX = p.x
			}
			if p.y > maxY {
				maxY = p.y
			}

			if p.x < minX {
				minX = p.x
			}
			if p.y < minY {
				minY = p.y
			}
		}

		if maxX-minX > width {
			//World has started expanding again :-()
			fmt.Println("Smallest area at second", r-1)
			break
		}
		width = maxX - minX
	}

	// Rewind two steps, this was discovered manually :-(
	for x := 0; x < 2; x++ {
		for i, p := range points {
			points[i].x, points[i].y = p.x-p.vx, p.y-p.vy
		}
	}
	drawGrid(points, minX, maxX, minY, maxY)
}

func drawGrid(points []point, minX, maxX, minY, maxY int) {
	grid := make([][]string, maxY-minY+1)
	for i := range grid {
		grid[i] = make([]string, maxX-minX+1)
		for j := range grid[i] {
			grid[i][j] = " "
		}
	}

	for i, p := range points {
		s := strconv.Itoa(i % 10)
		grid[p.y-minY][p.x-minX] = s
	}

	for _, g := range grid {
		fmt.Println(g)
	}
}

type point struct {
	x,
	y,
	vx,
	vy int
}

var test = `position=< 9,  1> velocity=< 0,  2>
position=< 7,  0> velocity=<-1,  0>
position=< 3, -2> velocity=<-1,  1>
position=< 6, 10> velocity=<-2, -1>
position=< 2, -4> velocity=< 2,  2>
position=<-6, 10> velocity=< 2, -2>
position=< 1,  8> velocity=< 1, -1>
position=< 1,  7> velocity=< 1,  0>
position=<-3, 11> velocity=< 1, -2>
position=< 7,  6> velocity=<-1, -1>
position=<-2,  3> velocity=< 1,  0>
position=<-4,  3> velocity=< 2,  0>
position=<10, -3> velocity=<-1,  1>
position=< 5, 11> velocity=< 1, -2>
position=< 4,  7> velocity=< 0, -1>
position=< 8, -2> velocity=< 0,  1>
position=<15,  0> velocity=<-2,  0>
position=< 1,  6> velocity=< 1,  0>
position=< 8,  9> velocity=< 0, -1>
position=< 3,  3> velocity=<-1,  1>
position=< 0,  5> velocity=< 0, -1>
position=<-2,  2> velocity=< 2,  0>
position=< 5, -2> velocity=< 1,  2>
position=< 1,  4> velocity=< 2,  1>
position=<-2,  7> velocity=< 2, -2>
position=< 3,  6> velocity=<-1, -1>
position=< 5,  0> velocity=< 1,  0>
position=<-6,  0> velocity=< 2,  0>
position=< 5,  9> velocity=< 1, -2>
position=<14,  7> velocity=<-2,  0>
position=<-3,  6> velocity=< 2, -1>`
