package main

import (
	"log"
	"strconv"
)

func matchesPasswordCritiera(val int) bool {
	s := strconv.Itoa(val)
	if !(s[0] == s[1] ||
		s[1] == s[2] ||
		s[2] == s[3] ||
		s[3] == s[4] ||
		s[4] == s[5]) {
		return false
	}

	if s[0] > s[1] ||
		s[1] > s[2] ||
		s[2] > s[3] ||
		s[3] > s[4] ||
		s[4] > s[5] {
		return false
	}

	return true
}

func matchesPasswordCritiera2(val int) bool {
	s := strconv.Itoa(val)

	var found bool
	for i := 0; i < 5; i++ {

		expandsLeft := i > 0 && s[i] == s[i-1]
		expandsRight := i < 4 && s[i+1] == s[i+2]
		if s[i] == s[i+1] && !expandsLeft && !expandsRight {
			found = true
		}
	}

	if !found {
		return false
	}

	if s[0] > s[1] ||
		s[1] > s[2] ||
		s[2] > s[3] ||
		s[3] > s[4] ||
		s[4] > s[5] {
		return false
	}

	return true
}

func findInRange(start, end int, critiera func(int) bool) int {
	var count int
	for i := start; i <= end; i++ {
		if critiera(i) {
			count++
		}
	}
	return count
}

func main() {
	if matchesPasswordCritiera(111111) != true {
		log.Fatal("111111 != true")
	}

	if matchesPasswordCritiera(223450) != false {
		log.Fatal("223450 != false")
	}

	if matchesPasswordCritiera(123789) != false {
		log.Fatal("123789 != false")
	}

	tcs := []struct {
		in  int
		out bool
	}{
		{112233, true},
		{123444, false},
		{111122, true},
		{111123, false},
	}

	for _, tc := range tcs {
		if matchesPasswordCritiera2(tc.in) != tc.out {
			log.Fatalf("%d != %t", tc.in, tc.out)
		}
	}

	pass := findInRange(246540, 787419, matchesPasswordCritiera2)
	log.Println("Answer part 2:", pass)

	log.Println("success!")
}
