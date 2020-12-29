package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

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
	sum := calculateSumPart1(input)
	fmt.Println("answer part 1:", sum)
}

func part2(input []byte) {
	sum := calculateSumPart2(input)
	fmt.Println("answer part 2:", sum)
}

func clearBit(n uint64, pos int) uint64 {
	mask := uint64(^(1 << pos))
	n &= mask
	return n
}

func setBit(n uint64, pos int) uint64 {
	n |= 1 << pos
	return n
}

func calculateSumPart1(input []byte) uint64 {
	var mask string
	memory := map[uint64]uint64{}
	for _, line := range strings.Split(string(input), "\n") {
		if strings.HasPrefix(line, "mask") {
			mask = strings.Split(line, " = ")[1]
			continue
		}

		var address, value uint64
		fmt.Sscanf(line, "mem[%d] = %d", &address, &value)

		// Apply mask
		value = applyMask(mask, value)
		memory[address] = value
	}

	var sum uint64
	for _, val := range memory {
		sum += val
	}
	return sum
}

func applyMask(mask string, value uint64) uint64 {
	for i, c := range mask {
		switch c {
		case '0':
			value = clearBit(value, 36-i-1)
		case '1':
			value = setBit(value, 36-i-1)
		case 'X':
			// Do nothing
		default:
			log.Fatal("unexpected mask value", mask)
		}
		if c == 'X' {
			continue
		}
	}
	return value
}

func calculateSumPart2(input []byte) uint64 {
	var mask string
	memory := map[uint64]uint64{}
	for _, line := range strings.Split(string(input), "\n") {
		if strings.HasPrefix(line, "mask") {
			mask = strings.Split(line, " = ")[1]
			continue
		}

		var originalAddress, value uint64
		fmt.Sscanf(line, "mem[%d] = %d", &originalAddress, &value)

		for _, address := range getAddresses(originalAddress, mask) {
			memory[address] = value
		}
	}

	var sum uint64
	for _, val := range memory {
		sum += val
	}
	return sum
}

func getAddresses(address uint64, mask string) []uint64 {
	return getAddressWithIndex(address, mask, 0)
}

func getAddressWithIndex(address uint64, mask string, offset int) []uint64 {
	if offset == len(mask) {
		return []uint64{address}
	}

	newOffset := offset + 1

	switch mask[offset] {
	case '0':
		// Do nothing
		return getAddressWithIndex(address, mask, newOffset)
	case '1':
		// Replace with 1
		return getAddressWithIndex(setBit(address, 36-offset-1), mask, newOffset)
	case 'X':

		// Branch where bit is 1
		oneAddresses := getAddressWithIndex(setBit(address, 36-offset-1), mask, newOffset)

		// Add branch where bit is 0
		return append(oneAddresses, getAddressWithIndex(clearBit(address, 36-offset-1), mask, newOffset)...)
	default:
		log.Fatal("unexpected mask value", mask)
		return []uint64{}
	}
}
