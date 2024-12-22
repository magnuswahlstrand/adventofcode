package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	id    string
	left  string
	right string
}

func part1Tests() {
	input := "LLR\n\nAAA = (BBB, BBB)\nBBB = (AAA, ZZZ)\nZZZ = (ZZZ, ZZZ)"
	step := calculateSteps(input)
	fmt.Println("step", step)
}

func part2Tests() {
	input := "LR\n\n11A = (11B, XXX)\n11B = (XXX, 11Z)\n11Z = (11B, XXX)\n22A = (22B, XXX)\n22B = (22C, 22C)\n22C = (22Z, 22Z)\n22Z = (22B, 22B)\nXXX = (XXX, XXX)"
	step := calculateSteps2(input)
	fmt.Println("step", step)
}

func calculateSteps(input string) int {
	pattern, rest, _ := strings.Cut(input, "\n\n")
	nodes := map[string]Node{}
	for _, line := range strings.Split(rest, "\n") {
		n := Node{
			id:    line[:3],
			left:  line[7:10],
			right: line[12:15],
		}
		nodes[n.id] = n
	}

	var step int
	current := "AAA"
	for ; current != "ZZZ"; step++ {
		if pattern[step%len(pattern)] == 'L' {
			current = nodes[current].left
		} else {
			current = nodes[current].right
		}
	}
	return step
}

var primeNumbers = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541, 547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911, 919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997}

func calculateSteps2(input string) int {
	pattern, rest, _ := strings.Cut(input, "\n\n")
	nodes := map[string]Node{}
	currents := []string{}

	for _, line := range strings.Split(rest, "\n") {
		n := Node{
			id:    line[:3],
			left:  line[7:10],
			right: line[12:15],
		}
		nodes[n.id] = n
		if n.id[2] == 'A' {
			currents = append(currents, n.id)
		}
	}

	var step int

	var cycleTimes []int
	for i := range currents {
		cycleTime, firstCycleAfter := -1, -1
		for {
			//fmt.Println("search", currents[p])
			if pattern[step%len(pattern)] == 'L' {
				currents[i] = nodes[currents[i]].left
			} else {
				currents[i] = nodes[currents[i]].right
			}
			step = step + 1
			hit := currents[i][2] == 'Z'
			if hit && cycleTime == -1 && firstCycleAfter != -1 {
				cycleTime = step - firstCycleAfter
				cycleTimes = append(cycleTimes, cycleTime)
				break
			}
			//
			if hit && firstCycleAfter == -1 {
				firstCycleAfter = step
			}
		}
		fmt.Println(cycleTime, firstCycleAfter, currents[i])
	}

	found := map[int]bool{}
	for i := range cycleTimes {
		for _, p := range primeNumbers {
			// TODO: Should really check if it is a factor multiple times
			if cycleTimes[i]%p == 0 {
				fmt.Println("found", p, cycleTimes[i])
				cycleTimes[i] = cycleTimes[i] / p
				found[p] = true
			}

			if cycleTimes[i] == 1 {
				break
			}
		}
	}

	LCM := 1
	for k := range found {
		fmt.Print(k, " ")
		LCM = LCM * k
	}

	fmt.Println("LCM", LCM)
	return LCM
}

func part1(input []byte) {
	sum := calculateSteps(string(input))
	fmt.Printf("answer to part 1 is %d\n", sum)
}

func part2(input []byte) {
	sum := calculateSteps2(string(input))
	fmt.Printf("answer to part 2 is %d\n", sum)
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1Tests()
	//part1(input)

	part2Tests()
	part2(input)

	log.Println("success")
}
