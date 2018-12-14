package main

import (
	"fmt"
)

type node struct {
	prev, next *node
	val        int
}

func (n node) String() string {
	return fmt.Sprintf("%d", n.val)
}

func newRecipes(sum int) []int {
	if sum > 9 {
		return []int{sum / 10, sum % 10}
	}
	return []int{sum}
}

func printList(n, e1, e2 *node) {
	p := n
	first := true
	for p != n || first {
		first = false

		switch p {
		case e1:
			fmt.Printf("(%s)", e1)
		case e2:
			fmt.Printf("[%s]", e2)

		default:
			fmt.Printf(" %s ", p)
		}
		p = p.next
	}
	fmt.Println()
}

func move(n *node) *node {
	steps := n.val + 1
	for s := 0; s < steps; s++ {
		n = n.next
	}
	return n
}

func initialElfs() (*node, *node) {
	elf1 := &node{val: 3}
	elf2 := &node{val: 7}

	elf1.prev = elf2
	elf1.next = elf2

	elf2.prev = elf1
	elf2.next = elf1
	return elf1, elf2
}

func recipeScore(after int) string {
	var nRecipes = 2
	elf1, elf2 := initialElfs()
	first := elf1
	end := elf2

	for nRecipes <= after+10 {

		// Create new recipes
		sum := elf1.val + elf2.val

		// Add to list
		for _, score := range newRecipes(sum) {
			n := &node{
				prev: end,
				next: first,
				val:  score,
			}

			// Insert node at the end
			end.next = n
			first.prev = n
			end = n

			nRecipes++

		}
		// printList(first, elf1, elf2)

		// Move forward
		elf1 = move(elf1)
		elf2 = move(elf2)
	}

	// Correct for one extra recipe
	if nRecipes-after == 11 {
		end = end.prev
	}

	s := ""
	for i := 0; i < 10; i++ {
		s = fmt.Sprint(end) + s
		end = end.prev
	}
	return s
}

func check(n *node, wanted []int) bool {
	for i := len(wanted) - 1; i >= 0; i-- {

		if wanted[i] != n.val {
			return false
		}
		n = n.prev
	}
	return true
}

func toInt(digits string) []int {

	var out []int
	for _, d := range digits {
		out = append(out, int(d-'0'))
	}
	return out
}

func recipeScoreV2(digits string) int {
	wanted := toInt(digits)
	var nRecipes = 2
	elf1, elf2 := initialElfs()
	first := elf1
	end := elf2

	for i := 0; i < 100000000; i++ {

		// Create new recipes
		sum := elf1.val + elf2.val

		// Add to list
		for _, score := range newRecipes(sum) {
			n := &node{
				prev: end,
				next: first,
				val:  score,
			}
			end.next = n
			first.prev = n
			end = n
			nRecipes++

			if check(end, wanted) {
				return nRecipes - len(wanted)
			}
		}

		// Move forward
		elf1 = move(elf1)
		elf2 = move(elf2)
	}

	return -1
}

func main() {

	fmt.Println("Day 13 - 2018")

	// Testcases
	// fmt.Println("The score after 5 iterations is", recipeScore(5) == "0124515891")
	// fmt.Println("The score after 9 iterations is", recipeScore(9) == "5158916779")
	// fmt.Println("The score after 18 iterations is", recipeScore(18) == "9251071085")
	// fmt.Println("The score after 2018 iterations is", recipeScore(2018) == "5941429882")
	// res := recipeScoreV2("51589")
	// fmt.Println("51589 was found after", res, "iterations", res == 9)
	// res = recipeScoreV2("01245")
	// fmt.Println("01245 was found after", res, "iterations", res == 5)
	// res = recipeScoreV2("92510")
	// fmt.Println("92510 was found after", res, "iterations", res == 18)
	// res = recipeScoreV2("59414")
	// fmt.Println("59414 was found after", res, "iterations", res == 2018)

	fmt.Println("The score after 760221 iterations is", recipeScore(760221))
	res := recipeScoreV2("760221")
	fmt.Println("760221 was found after", res, "iterations")
}
