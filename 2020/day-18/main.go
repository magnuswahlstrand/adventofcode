package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"text/scanner"
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
	var sum int
	for _, s := range strings.Split(string(input), "\n") {
		sum += evaluate(s)
	}

	fmt.Println("answer part 1:", sum)
}

func part2(input []byte) {
	var sum int
	for _, s := range strings.Split(string(input), "\n") {
		v := evaluate2(s)
		fmt.Println(v)
		sum += v
	}

	fmt.Println("answer part 2:", sum)
}

type stackItem struct {
	num      int
	operator string
	isFirst  bool
}

type stackItem2 struct {
	num        int
	operator   string
	multiStack []int
	isFirst    bool
}

func evaluate(src string) int {

	var s scanner.Scanner
	s.Init(strings.NewReader(src))

	var stack []stackItem

	var num int
	var operator string

	isFirstNumber := true
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		txt := s.TokenText()

		switch txt {
		case "*", "+":
			operator = txt
		case "(":
			stack = append(stack, stackItem{
				num:      num,
				operator: operator,
				isFirst:  isFirstNumber,
			})
			isFirstNumber = true
			num = 0
		case ")":
			// Pop from stack
			n := len(stack) - 1 // Top element
			prev := stack[n]
			stack = stack[:n]

			// Complete previous operation
			switch {
			case prev.isFirst:
				// keep num
				// num = num
				isFirstNumber = false
			default:
				//num = calculate(num, val, operator)
				num = calculate(prev.num, num, prev.operator)
			}

		default:
			val, err := strconv.Atoi(txt)
			if err != nil {
				log.Fatal("failed to parse int", err)
			}

			switch {
			case isFirstNumber:
				num = val
				isFirstNumber = false
			default:
				num = calculate(num, val, operator)
			}
		}
		//fmt.Printf("%s - %d \n", s.TokenText(), num)
	}
	return num
}

func calculate(v1 int, v2 int, operator string) int {
	switch operator {
	case "*":
		return v1 * v2
	case "+":
		return v1 + v2
	default:
		log.Fatal("unexpected operator", operator)
		return -1
	}
}

func evaluate2(src string) int {
	var s scanner.Scanner
	s.Init(strings.NewReader(src))

	var stack []stackItem2
	var multiStack []int

	var num int
	var operator string
	isFirstNumber := true
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		txt := s.TokenText()

		switch txt {
		case "*", "+":
			operator = txt

		case "(":
			fmt.Println("down", num, operator, multiStack)
			stack = append(stack, stackItem2{
				num:        num,
				operator:   operator,
				multiStack: multiStack,
				isFirst:    isFirstNumber,
			})
			isFirstNumber = true
			num = 0
			multiStack = []int{}
		case ")":
			// Evaluate multiplication stack
			num = runMultiplication(multiStack, num)
			// Pop from stack
			n := len(stack) - 1 // Top element
			prev := stack[n]
			stack = stack[:n]

			// current num
			multiStack = prev.multiStack
			fmt.Println("up", num, multiStack)

			if prev.isFirst {

			} else {
				switch prev.operator {
				case "*":
					multiStack = append(multiStack, prev.num)
				case "+":
					num = prev.num + num
				}
			}
			isFirstNumber = false

		default:
			val, err := strconv.Atoi(txt)
			if err != nil {
				log.Fatal("failed to parse int", err)
			}

			switch {
			case isFirstNumber:
				num = val
				isFirstNumber = false
			default:
				switch operator {
				case "*":
					multiStack = append(multiStack, num)
					num = val
				case "+":
					num = num + val
				default:
					log.Fatalf("unexpected operator %q\n", operator)
					return -1
				}
			}
		}
	}

	// Evaluate final multistack
	num = runMultiplication(multiStack, num)

	return num
}

func runMultiplication(multiStack []int, num int) int {
	for _, factor := range multiStack {
		num *= factor
	}
	return num
}

// Part 2 - too low : 109359708097321
