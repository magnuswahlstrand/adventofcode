package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func toIntSlice(in string) []int {
	var s []int
	for _, v := range strings.Split(in, " ") {
		i, _ := strconv.Atoi(v)
		s = append(s, i)
	}
	return s
}

const (
	CHILDREN = 0
	METADATA = 1
)

func traverseTree(numbers []int) int {
	var total int
	var stack [][]int

	// Add initial item in stack
	stack = append(stack, []int{1, 0})

	for i := 0; i < len(numbers); i++ {

		current := stack[len(stack)-1]
		if current[CHILDREN] == 0 {
			// All children has been parsed, parse any remaining metadata
			nodeValue := 0
			for ; current[METADATA] > 0; current[METADATA]-- {
				nodeValue += numbers[i]
				i++
			}
			total += nodeValue
			i-- // Correct position

			// Remove current from stack
			stack = stack[:len(stack)-1]

			// Decrease previous with 1
			stack[len(stack)-1][CHILDREN]--

			// Check if parent had this node in it

		} else {
			// Parse a new child and add to stack
			header := numbers[i : i+2]
			i++
			stack = append(stack, header)

		}
	}

	return total
}

type Node struct {
	nChildren      int
	nMeta          int
	childrenValues []int
	value          int
	edge           bool
}

func traverseTreeV2(numbers []int) int {
	var stack []Node

	// Add initial item in stack
	stack = append(stack, Node{nChildren: 1, nMeta: 0})

	for i := 0; i < len(numbers); i++ {

		current := &stack[len(stack)-1]

		if current.nChildren == 0 {
			// All children has been parsed, parse any remaining metadata
			for ; current.nMeta > 0; current.nMeta-- {

				childNumber := numbers[i]
				if childNumber <= len(current.childrenValues) {
					current.value += current.childrenValues[childNumber-1]
				}
				i++
			}
			i-- // Correct position

			// Remove current from stack
			stack = stack[:len(stack)-1]

			// Decrease previous with 1
			stack[len(stack)-1].nChildren--
			stack[len(stack)-1].childrenValues = append(stack[len(stack)-1].childrenValues, current.value)

		} else {
			// Parse a new child and add to stack
			n := Node{
				nChildren: numbers[i],
				nMeta:     numbers[i+1],
			}
			i++ // Increment for meta-field

			if n.nChildren == 0 {
				// Edge-node, calculate value directly
				for ; n.nMeta > 0; n.nMeta-- {
					i++ //
					n.value += numbers[i]
				}

				n.edge = true
				current.childrenValues = append(current.childrenValues, n.value)

				stack[len(stack)-1].nChildren--

			} else {
				stack = append(stack, n)
			}
		}
	}

	return stack[0].childrenValues[0]
}

func main() {
	fmt.Println("Advent of Code - Day 8 - Go")

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// test := toIntSlice("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2")
	//fmt.Println("Test 1 -", traverseTree(test))
	// fmt.Println("Test 2 -", traverseTreeV2(test))
	fmt.Println("Part 1 -", traverseTree(toIntSlice(string(content))))   // 45 min
	fmt.Println("Part 1 -", traverseTreeV2(toIntSlice(string(content)))) // 45 min

}
