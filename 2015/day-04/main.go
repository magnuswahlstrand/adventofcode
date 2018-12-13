package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"strconv"
)

const veryHighNumber = 100000000

func findHash(initial []byte) int {
	for i := 0; i < veryHighNumber; i++ {
		test := append(initial, []byte(strconv.Itoa(i))...)
		res := md5.Sum(test)

		if res[0] == 0 && res[1] == 0 && (res[2]&0xF0) == 0 {
			return i
		}

	}
	return -1
}

func findHash6(initial []byte) int {
	zeroes := []byte{0, 0, 0}
	for i := 0; i < veryHighNumber; i++ {
		test := append(initial, []byte(strconv.Itoa(i))...)
		res := md5.Sum(test)

		if bytes.Equal(res[:3], zeroes) {
			return i
		}

	}
	return -1
}

func main() {
	fmt.Println("")
	fmt.Println("Day 4 - 2015")
	fmt.Println(findHash([]byte("abcdef")) == 609043)
	fmt.Println(findHash([]byte("pqrstuv")) == 1048970)

	fmt.Println("Part 1:", findHash([]byte("yzbqklnj")))
	fmt.Println("Part 1:", findHash6([]byte("yzbqklnj")))
}
