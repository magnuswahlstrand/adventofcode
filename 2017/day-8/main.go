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
	registries := make(map[string]*int)
	for _, cmd := range cmds {
		fields := strings.Fields(cmd)
		target := fields[0]
		cond := fields[4]

		if registries[target] == nil {
			registries[target] = new(int)
		}

		if registries[cond] == nil {
			registries[cond] = new(int)
		}

		// Check condition
		var condition bool
		switch fields[5] {
		case "<":
			condition = *registries[cond] < atoi(fields[6])
		case ">":
			condition = *registries[cond] > atoi(fields[6])
		case ">=":
			condition = *registries[cond] >= atoi(fields[6])
		case "<=":
			condition = *registries[cond] <= atoi(fields[6])
		case "==":
			condition = *registries[cond] == atoi(fields[6])
		case "!=":
			condition = *registries[cond] != atoi(fields[6])
		default:
			log.Fatal(fields[5])
		}

		// Perform command
		if condition {
			switch fields[1] {
			case "inc":
				*registries[target] += atoi(fields[2])
			case "dec":
				*registries[target] -= atoi(fields[2])
			default:
				log.Fatal(fmt.Sprintf("'%s'", fields[1]))
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
