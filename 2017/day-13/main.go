package main

import (
	"fmt"
	"strings"

	"github.com/pkg/profile"
)

// ScannerBoard holds the state of all scanners
type ScannerBoard struct {
	pos      int
	scanners []*scanner
}

func (sb *ScannerBoard) move() {
	sb.pos++
}

func (sb *ScannerBoard) updateInitial(steps int) {

	// Optimization
	modMap := make(map[int]int)

	var modSteps int
	for i, sc := range sb.scanners {

		// Optimization
		if val, ok := modMap[sc.length]; ok {
			modSteps = val
		} else {
			cycleTime := (2 * (sc.length - 1))
			modSteps = (steps % cycleTime)
			modMap[sc.length] = modSteps
		}

		offset := sc.id - i

		for d := 0; d < modSteps+offset; d++ {
			sc.pos += sc.v
			if sc.pos >= sc.length+1 || sc.pos == 0 {
				sc.v *= -1
				sc.pos += 2 * sc.v
			}
		}
	}

}

func (sb *ScannerBoard) setInitial() {

	for i, sc := range sb.scanners {

		offset := sc.id - i

		for d := 0; d < offset; d++ {

			sc.pos += sc.v
			if sc.pos >= sc.length+1 || sc.pos == 0 {
				sc.v *= -1
				sc.pos += 2 * sc.v
			}
		}

		sc.initV = sc.v
		sc.initPos = sc.pos
	}
}

func (sb *ScannerBoard) update() {

	for _, sc := range sb.scanners {

		sc.pos += sc.v
		if sc.pos >= sc.length+1 || sc.pos == 0 {
			sc.v *= -1
			sc.pos += 2 * sc.v
		}
	}
}

func (sb ScannerBoard) String() string {
	s := ""
	var max int
	for i, sc := range sb.scanners {
		s += fmt.Sprintf(" %d ", i)
		if sc.length > max {
			max = sc.length
		}
	}

	s += "\n"
	for r := 1; r <= max; r++ {

		for i, sc := range sb.scanners {

			var marker, format string
			marker = " "
			if sc.length >= r {
				format = "[%s]"
			} else {
				format = " %s "
				if r == 1 {
					format = " %s "
					marker = "."
				}
			}

			if r == 1 && sb.pos == i {
				format = "(%s)"
			}

			if sc.pos == r {
				marker = "S"
			}

			s += fmt.Sprintf(format, marker)
		}
		s += "\n"
	}

	return s
}

func (sb *ScannerBoard) reset() {
	for _, sc := range sb.scanners {
		sc.pos = sc.initPos
		sc.v = sc.initV
	}

	// sb.pos = -1
	// for _, sc := range sb.scanners {
	// 	if sc.length > 0 {
	// 		sc.v = 1
	// 		sc.pos = 1
	// 	}
	// }
}

func (sb *ScannerBoard) setNext() {
	for _, sc := range sb.scanners {
		var pos, v int
		pos = sc.initPos
		v = sc.initV

		pos += v
		if pos >= sc.length+1 || pos == 0 {
			v *= -1
			pos += 2 * v
		}

		sc.initPos = pos
		sc.initV = v
	}
}

func (sb *ScannerBoard) updateInitialV2() {

	for _, sc := range sb.scanners {

		var pos, v int
		pos = sc.initPos
		v = sc.initV

		pos += sc.v
		if pos >= sc.length+1 || pos == 0 {
			v *= -1
			pos += 2 * sc.v
		}

		sc.initPos = pos
		sc.initV = v
	}
}

type scanner struct {
	id      int
	length  int
	v       int
	pos     int
	initV   int
	initPos int
}

func (sb ScannerBoard) check(strict bool) int {

	if sb.pos < 0 {
		return 0
	}

	currentScanner := *sb.scanners[sb.pos]
	if currentScanner.pos != 1 {
		return 0
	}

	dmg := sb.pos * currentScanner.length
	// fmt.Printf("Was hit for %d damage\n", sb.pos*currentScanner.length)
	if strict && dmg == 0 {
		return 10000000
	}
	return sb.pos * currentScanner.length

}

func main() {

	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	var id, length int
	// Get highest ID
	rows := strings.Split(input, "\n")
	fmt.Sscanf(rows[len(rows)-1], "%d: %d", &id, &length)

	// ss := make([]*scanner, id+1)
	// for i := 0; i < id; i++ {
	// 	ss[i] = &scanner{
	// 		length: 0,
	// 		v:      0,
	// 		pos:    0,
	// 	}
	// }

	ss := []*scanner{}
	for _, row := range rows {
		fmt.Sscanf(row, "%d: %d", &id, &length)
		ss = append(ss, &scanner{
			id:      id,
			length:  length,
			pos:     1,
			v:       1,
			initV:   1,
			initPos: 1,
		})
	}

	sb := ScannerBoard{
		pos:      0,
		scanners: ss,
	}
	sb.setInitial()

	strict := true
	for delay := 0; delay < 10000000; delay++ {
		// fmt.Println("GOGO", delay)
		sb.reset()
		sb.setNext()
		sb.pos = -1

		// fmt.Printf("Initial\n")

		// fmt.Println(sb)
		// sb.updateInitialV2()
		// sb.updateInitial()
		// sb.reset()
		// sb.setNext()
		// sb.updateInitial(delay)

		var dmg int

	Inner:
		for round := 0; sb.pos < len(ss)-1; round++ {
			// fmt.Printf("Pico sec %d BEFORE (%d)\n", round, delay)
			// fmt.Println(sb)
			sb.move()
			dmg += sb.check(strict)

			if dmg > 0 {
				break Inner
			}
			sb.update()
			// fmt.Printf("Pico sec %d AFTER (%d)\n", round, delay)
			// fmt.Println(sb)
		}

		if dmg == 0 {
			fmt.Printf("Total damage: %d with a delay of %d\n", dmg, delay)
			break
		}
	}

	// Initial solution took 11s for 10k
	// Moving inital delay outside of loop took 1s for 10k
	// Using modulus to remove steps --> 2.7 for 100k
	// Skip modulus, and just step 1 ahead of initial previous step --> 3.9 for 10M steps

	// 309984 too low
}

var example = `0: 3
1: 2
4: 4
6: 4`

var input = `0: 4
1: 2
2: 3
4: 5
6: 8
8: 4
10: 6
12: 6
14: 6
16: 10
18: 6
20: 12
22: 8
24: 9
26: 8
28: 8
30: 8
32: 12
34: 12
36: 12
38: 8
40: 10
42: 14
44: 12
46: 14
48: 12
50: 12
52: 12
54: 14
56: 14
58: 14
60: 12
62: 14
64: 14
68: 12
70: 14
74: 14
76: 14
78: 14
80: 17
82: 28
84: 18
86: 14`
