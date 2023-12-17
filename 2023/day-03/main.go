package main

import (
	"bytes"
	"fmt"
	"golang.org/x/exp/maps"
	"log"
	"os"
	"strconv"
	"text/scanner"
)

type SetColor struct {
	Count int
	Color string
}

type Number struct {
	x    int
	maxX int
	y    int
	val  int
}

func part1Tests() {
	input := []byte("467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..")

	sum := sumOfNumbers(input)
	fmt.Print(sum)
}

func sumOfNumbers(input []byte) int {
	symbolCoordinates := map[[2]int]bool{}
	numbers := []Number{}

	//for lineNo, line := range bytes.Split(input, []byte{'\n'}) {
	var scnner scanner.Scanner
	scnner.Init(bytes.NewReader(bytes.Replace(input, []byte{'.'}, []byte{' '}, -1)))
	for tok := scnner.Scan(); tok != scanner.EOF; tok = scnner.Scan() {
		switch tok {
		case scanner.Int:
			val, _ := strconv.Atoi(scnner.TokenText())
			numbers = append(numbers, Number{
				scnner.Position.Column,
				scnner.Position.Column + len(scnner.TokenText()),
				scnner.Position.Line,
				val,
			})
			//for i := 0; i < len(s); i++ {
			//	fmt.Println(scnner.Position.Line, scnner.Position.Column+i)
			//}
		default:
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					symbolCoordinates[[2]int{scnner.Position.Column + dx, scnner.Position.Line + dy}] = true
				}
			}
		}
	}

	var sum int
	for _, number := range numbers {
		for x := number.x; x < number.maxX; x++ {
			if symbolCoordinates[[2]int{x, number.y}] {
				sum += number.val
				//fmt.Println("Found", number.val, "at", number.y, x)
				break
			}
		}

		//fmt.Println(string(line))

		//}
	}
	return sum
}

func sumOfNumbers2(input []byte) int {
	numberCoordinates := map[[2]int]int{}
	numbers := []Number{}
	gears := [][2]int{}

	//for lineNo, line := range bytes.Split(input, []byte{'\n'}) {
	var scnner scanner.Scanner
	scnner.Init(bytes.NewReader(bytes.Replace(input, []byte{'.'}, []byte{' '}, -1)))

	var numberIndex int
	for tok := scnner.Scan(); tok != scanner.EOF; tok = scnner.Scan() {
		switch tok {
		case scanner.Int:
			val, _ := strconv.Atoi(scnner.TokenText())
			numberIndex++

			//numbers = append(numbers, Number{
			//	scnner.Position.Column,
			//	scnner.Position.Column + len(scnner.TokenText()),
			//	scnner.Position.Line,
			//	val,
			//})

			for i := 0; i < len(scnner.TokenText()); i++ {
				numberCoordinates[[2]int{scnner.Position.Column + i, scnner.Position.Line}] = len(numbers)
			}
			numbers = append(numbers, Number{
				val: val,
			})

		default:
			if scnner.TokenText() == "*" {
				gears = append(gears, [2]int{scnner.Position.Column, scnner.Position.Line})
			}
		}
	}

	uniqueNumbers := map[int]bool{}
	var sum int
	for _, gear := range gears {
		fmt.Println(gear)
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				num, found := numberCoordinates[[2]int{gear[0] + dx, gear[1] + dy}]
				if found {
					uniqueNumbers[num] = true
				}
			}
		}
		if len(uniqueNumbers) == 2 {
			fmt.Println("YEAH")
			factor := 1
			for k := range uniqueNumbers {
				factor *= numbers[k].val
				fmt.Println(numbers[k].val, factor)
			}
			sum += factor
		}
		maps.Clear(uniqueNumbers)
	}

	return sum
}

func part2Tests() {
	input := []byte("467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..")

	sum := sumOfNumbers2(input)
	fmt.Print(sum)
}

func part1(input []byte) {
	sum := sumOfNumbers(input)
	fmt.Printf("answer to part 1 is %d\n", sum)
}
func part2(input []byte) {
	sum := sumOfNumbers2(input)
	fmt.Printf("answer to part 2 is %d\n", sum)
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1Tests()
	part1(input)
	part2Tests()
	part2(input)

	log.Println("success")
}
