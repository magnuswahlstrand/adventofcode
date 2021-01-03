package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
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
	fmt.Println("answer part 1:", calculatePart1EncryptionKey(input))
}

func part2(input []byte) {
	//fmt.Println("answer part 2:", sum)
}

type keyHolder struct {
	public, private int
	loopSize        int
}

func calculatePart1EncryptionKey(input []byte) int {
	split := strings.Split(string(input), "\n")

	cardKeys := keyHolder{
		public: mustInt(split[0]),
	}

	doorKeys := keyHolder{
		public: mustInt(split[1]),
	}

	cardKeys.calculateLoopSize()
	encryptionKey := cardKeys.calculateEncryptionKey(doorKeys.public)
	return encryptionKey
}

func (keys *keyHolder) calculateLoopSize() int {
	v := 1
	subjectNumber := 7
	keys.loopSize = 0
	for v != keys.public {
		keys.loopSize++

		// 1. Set the value to itself multiplied by the subject number.
		v *= subjectNumber

		// Set the value to the remainder after dividing the value by 20201227.
		v = v % 20201227
	}
	return keys.loopSize
}

func (keys *keyHolder) calculateEncryptionKey(subjectNumber int) int {
	v := 1
	for i := 0; i < keys.loopSize; i++ {
		// 1. Set the value to itself multiplied by the subject number.
		v *= subjectNumber

		// Set the value to the remainder after dividing the value by 20201227.
		v = v % 20201227
	}
	keys.private = v
	return keys.private
}
