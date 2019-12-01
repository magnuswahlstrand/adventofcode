package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func fuelRequired(mass int64) int64 {
	return int64((mass / 3) - 2)
}

func fuelRequired2(mass int64) int64 {
	f := int64((mass / 3) - 2)

	if f < 0 {
		return 0
	}

	// Add fuel needed for the fuel itself
	f += fuelRequired2(f)

	return f
}

func sumFuelRequired(masses []int64) int64 {
	var sum int64
	for _, m := range masses {
		sum += fuelRequired(m)
	}
	return sum
}

func part1Tests() {
	if fuelRequired(12) != 2 {
		log.Fatal("12 != 2")
	}

	if fuelRequired(100756) != 33583 {
		log.Fatal("100756 != 33583")
	}

	if sumFuelRequired([]int64{12, 14, 1969, 100756}) != (2 + 2 + 654 + 33583) {
		log.Fatal("not the same!")
	}
}

func part2Tests() {
	if fuelRequired2(14) != 2 {
		log.Fatal("14 != 2")
	}

	if fuelRequired2(1969) != 966 {
		log.Fatal("1969 != 966")
	}

	if fuelRequired2(100756) != 50346 {
		log.Fatal("100756 != 50346")
	}
}

func parseMass(input []byte) []int64 {
	var mass []int64

	rows := bytes.Split(input, []byte{'\n'})
	for _, r := range rows {
		m, err := strconv.ParseInt(string(r), 10, 64)
		if err != nil {
			log.Fatal("failed to convert", err)
		}

		mass = append(mass, m)
	}
	return mass
}

func part1(mass []int64) {
	var sum int64
	for _, m := range mass {
		sum += fuelRequired(m)
	}
	fmt.Println("total sum for part 1 is:", sum)
}

func part2(mass []int64) {
	var sum int64
	for _, m := range mass {
		sum += fuelRequired2(m)
	}
	fmt.Println("total sum for part 2 is:", sum)
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	mass := parseMass(input)

	part1Tests()
	part1(mass)

	part2Tests()
	part2(mass)

	log.Println("success")
}
