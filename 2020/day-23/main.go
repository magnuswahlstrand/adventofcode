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
	circle := setupCupCircle(input)
	for i := 0; i < 100; i++ {
		circle.step()
	}
	fmt.Println("answer part 1:", circle.StringFrom1())
}

func part2(input []byte) {
	circle := setupCupCirclePart2(input)
	// Run 10 m
	//for i := 0; i < 1; i++ {
	for i := 0; i < 10*1000*1000; i++ {
		if i%1000000 == 0 {
			fmt.Println(i)
		}
		circle.step()
	}
	fmt.Println("answer part 2:", circle.Part2Cup1()*circle.Part2Cup2())
}

type cup struct {
	label int
	next  *cup
}

type cupCircle struct {
	current     *cup
	cupsByLabel []*cup
	count       int

	destination *cup
}

func (c *cupCircle) step() {
	// Select 3 cups
	c.destination = c.cupsByLabel[c.getNextDestinationLabel()-1]
	threeCupEnd := c.current.next.next.next

	// Move all relevant "next cup" at once
	//fmt.Println(c.current, destination, threeCupEnd, destLabel)
	c.current.next, c.destination.next, threeCupEnd.next = threeCupEnd.next, c.current.next, c.destination.next

	// New current cup
	c.current = c.current.next
}

func (c *cupCircle) getNextDestinationLabel() int {
	destLabel := c.current.label
	for {
		destLabel--
		if destLabel < 1 {
			destLabel += c.count
		}

		switch destLabel {
		case c.current.next.label,
			c.current.next.next.label,
			c.current.next.next.next.label:
			continue
		}
		break
	}
	return destLabel
}

func setupCupCircle(input []byte) cupCircle {
	startingCups := string(input)

	cupsByLabel := make([]*cup, 9)

	label := mustInt(startingCups[0:1])
	current := &cup{label: label}
	cupsByLabel[current.label-1] = current

	// Create linked list
	previous := current
	for _, s := range startingCups[1:] {
		label := mustInt(string(s))
		c := &cup{label: label}

		// Link from previous
		previous.next = c

		// Store in lookup table
		cupsByLabel[c.label-1] = c

		previous = c
	}
	previous.next = current
	circle := cupCircle{
		current:     current,
		cupsByLabel: cupsByLabel,
		count:       len(cupsByLabel),
	}
	return circle
}

const million = 1000000

func setupCupCirclePart2(input []byte) cupCircle {
	startingCups := string(input)

	cupsByLabel := make([]*cup, million)

	label := mustInt(startingCups[0:1])
	current := &cup{label: label}
	cupsByLabel[current.label-1] = current

	// Create linked list
	previous := current
	for _, s := range startingCups[1:] {
		label := mustInt(string(s))

		c := &cup{label: label}

		// Link from previous
		previous.next = c

		// Store in lookup table
		cupsByLabel[c.label-1] = c

		previous = c
	}

	// Add another million
	for i := 9 + 1; i <= million; i++ {
		c := &cup{label: i}
		previous.next = c
		cupsByLabel[c.label-1] = c
		previous = c
	}
	previous.next = current

	circle := cupCircle{
		current:     current,
		cupsByLabel: cupsByLabel,
		count:       million,
	}
	return circle
}

func (c *cupCircle) String() string {
	var b strings.Builder

	// Add first cup
	cup := c.current
	for {
		b.WriteString(strconv.Itoa(cup.label))

		cup = cup.next
		if cup == c.current {
			break
		}
	}
	return b.String()
}

func (c *cupCircle) StringFrom1() string {
	var b strings.Builder

	// Start from cup after 1
	cup := c.cupsByLabel[0].next
	for {
		b.WriteString(strconv.Itoa(cup.label))

		cup = cup.next
		if cup == c.cupsByLabel[0] {
			break
		}
	}
	return b.String()
}

func (c *cupCircle) Part2Cup1() int {
	return c.cupsByLabel[0].next.label
}

func (c *cupCircle) Part2Cup2() int {
	return c.cupsByLabel[0].next.next.label
}

// Part 2: 345996672 too low
