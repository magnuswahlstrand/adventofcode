package main

import (
	"fmt"
	"strings"
)

// ScannerBoard holds the state of all scanners
type ScannerBoard struct {
	pos      int
	scanners []*scanner
}

func (sb *ScannerBoard) reset() {
	sb.pos = -1
	for _, sc := range sb.scanners {
		if sc.length > 0 {
			sc.v = 1
			sc.pos = 1
		}
	}
}

func (sb *ScannerBoard) move() {
	if sb.pos >= 0 {
		sb.pos++
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

type scanner struct {
	length int
	v      int
	pos    int
}

func (sb ScannerBoard) check() int {

	if sb.pos < 0 {
		return 0
	}

	currentScanner := *sb.scanners[sb.pos]
	if currentScanner.pos != 1 {
		return 0
	}
	fmt.Printf("Was hit for %d damage\n", sb.pos*currentScanner.length)
	return sb.pos * currentScanner.length
}

func main() {

	var id, length int
	// Get highest ID
	rows := strings.Split(example, "\n")
	fmt.Sscanf(rows[len(rows)-1], "%d: %d", &id, &length)

	ss := make([]*scanner, id+1)
	for i := 0; i < id; i++ {
		ss[i] = &scanner{
			length: 0,
			v:      0,
			pos:    0,
		}
	}

	for _, row := range rows {
		fmt.Sscanf(row, "%d: %d", &id, &length)
		ss[id] = &scanner{
			length: length,
			v:      1,
			pos:    1,
		}
	}

	sb := ScannerBoard{
		pos:      0,
		scanners: ss,
	}

	var dmg int
	for round := 0; round < len(ss); round++ {
		fmt.Println(sb)
		dmg += sb.check()
		sb.move()
		sb.update()
	}
	fmt.Printf("Total damage: %d", dmg)
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
