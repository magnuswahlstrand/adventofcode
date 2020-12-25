package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func part1(input string) {
	n := findLoopLine(input)
	fmt.Println("answer part 1:", n)
}

func part2(input string) {
	n := findAndFixLoop(input)
	fmt.Println("answer part 2:", n)
}

func parseInstruction(instr string) (string, int) {
	val, err := strconv.Atoi(strings.ReplaceAll(instr[4:], "+", ""))
	if err != nil {
		log.Fatal("failed to parse", instr, ":", err)
	}
	cmd := instr[:3]
	return cmd, val
}

func applyInstruction(cmd string, val int, line int, acc int) (int, int) {
	switch cmd {
	case "nop":
		line++
	case "jmp":
		line += val
	case "acc":
		acc += val
		line++
	default:
		log.Fatal("invalid command")
	}
	return line, acc
}

func findLoopLine(input string) int {
	instructions := strings.Split(input, "\n")
	n, _ := findLoopLineInner(instructions)
	return n
}

func findAndFixLoop(input string) int {
	instructions := strings.Split(input, "\n")
	for i, instruction := range instructions {
		cmd, _ := parseInstruction(instruction)
		if cmd == "acc" {
			continue
		}

		oldInstruction := instructions[i]
		switch cmd {
		case "nop":
			instructions[i] = strings.ReplaceAll(instructions[i], "nop", "jmp")
		case "jmp":
			instructions[i] = strings.ReplaceAll(instructions[i], "jmp", "nop")
		}

		acc, endReached := findLoopLineInner(instructions)
		if endReached {
			// Success!
			return acc
		} else {
			// Failure, return instruction to old format
			instructions[i] = oldInstruction
		}
	}
	return -1
}

func findLoopLineInner(instructions []string) (int, bool) {
	visited := map[int]int{}
	var acc int
	line := 0
	for {
		visited[line]++
		cmd, val := parseInstruction(instructions[line])
		line, acc = applyInstruction(cmd, val, line, acc)

		// Found loop
		if visited[line] > 0 {
			return acc, false
		}

		if line >= len(instructions) {
			return acc, true
		}
	}
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1(string(input))
	part2(string(input))

	log.Println("success")
}
