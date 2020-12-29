package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

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
	_, pos := finalPos(string(input))
	fmt.Println("answer part 1:", pos.x+pos.y)
}

func part2(input []byte) {
	_, pos := finalPos2(string(input))
	fmt.Println("answer part 2:", pos.x+pos.y)
}

type vec struct {
	x int
	y int
}

func finalPos(input string) (string, vec) {
	dir := "E"
	pos := vec{0, 0}
	for _, line := range strings.Split(input, "\n") {
		cmd := line[0:1]

		i, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(line, err)
		}

		switch cmd {
		case "L":
			//Action L means to turn left the given number of degrees.
			dir = changeDir("L", dir, i)
		case "R":
			//Action R means to turn right the given number of degrees.
			dir = changeDir("R", dir, i)
		case "F":
			//Action F means to move forward by the given value in the direction the ship is currently facing.
			dx, dy := moveForward(dir)
			pos.x += i * dx
			pos.y += i * dy
		case "E", "S", "W", "N":
			//Action N means to move north by the given value.
			//Action S means to move south by the given value.
			//Action E means to move east by the given value.
			//Action W means to move west by the given value.
			dx, dy := moveForward(cmd)
			pos.x += i * dx
			pos.y += i * dy
		}
	}
	return dir, pos
}
func finalPos2(input string) (string, vec) {
	dir := "E"
	pos := vec{0, 0}
	wpPos := vec{10, -1}
	for _, line := range strings.Split(input, "\n") {
		cmd := line[0:1]

		i, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(line, err)
		}

		switch cmd {
		case "L":
			wpPos = rotateWaypoint("L", wpPos, i)
		case "R":
			wpPos = rotateWaypoint("R", wpPos, i)
		case "F":
			pos.x += i * wpPos.x
			pos.y += i * wpPos.y
		case "E", "S", "W", "N":
			dx, dy := moveForward(cmd)
			wpPos.x += i * dx
			wpPos.y += i * dy
		}
	}
	return dir, pos
}

func rotateWaypoint(action string, pos vec, i int) vec {
	var steps int
	if action == "L" {
		steps = (360 - i) / 90
	} else {
		steps = i / 90
	}

	x, y := pos.x, pos.y
	for i := 0; i < steps; i++ {
		x, y = -y, x
	}
	return vec{x, y}
}

func moveForward(dir string) (int, int) {
	switch dir {
	case "N":
		return 0, -1
	case "S":
		return 0, 1
	case "E":
		return 1, 0
	case "W":
		return -1, 0
	default:
		log.Fatal("invalid direction")
		return -1, -1
	}
}

func changeDir(action string, dir string, i int) string {
	var steps int
	if action == "L" {
		steps = (360 - i) / 90
	} else {
		steps = i / 90
	}

	for i := 0; i < steps; i++ {
		switch dir {
		case "E":
			dir = "S"
		case "S":
			dir = "W"
		case "W":
			dir = "N"
		case "N":
			dir = "E"
		}
	}
	return dir
}
