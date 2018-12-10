package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func calc(cmds []string) (int, int) {
	var maxEver int
	var target, condA, operator, comperator, tmp string
	var change, condB int
	registries := make(map[string]*int)
	for _, cmd := range cmds {
		fmt.Sscanf(cmd, "%s %s %d %s %s %s %d", &target, &operator, &change, &tmp, &condA, &comperator, &condB)

		if registries[target] == nil {
			registries[target] = new(int)
		}

		if registries[condA] == nil {
			registries[condA] = new(int)
		}

		// Check condition
		var condition bool
		switch comperator {
		case "<":
			condition = *registries[condA] < condB
		case ">":
			condition = *registries[condA] > condB
		case ">=":
			condition = *registries[condA] >= condB
		case "<=":
			condition = *registries[condA] <= condB
		case "==":
			condition = *registries[condA] == condB
		case "!=":
			condition = *registries[condA] != condB
		default:
			log.Fatal(comperator)
		}

		// Perform command
		if condition {
			switch operator {
			case "inc":
				*registries[target] += change
			case "dec":
				*registries[target] -= change
			default:
				log.Fatal(fmt.Sprintf("'%s'", operator))
			}

			if *registries[target] > maxEver {
				maxEver = *registries[target]
			}
		}
	}

	max := -1
	for _, r := range registries {
		if *r > max {
			max = *r
		}
	}

	return max, maxEver
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func main() {
	//	instructions := strings.Split(example, "\n")
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	instructions := strings.Split(string(content), "\n")
	a, b := calc(instructions)
	fmt.Printf("The max(ever) value is %d(%d)", a, b)

	// Part 1 = 28 min
	// Part 2 = 2 min
}

var example = `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`
