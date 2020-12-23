package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type RuleAndPassword struct {
	val1      int
	val2      int
	character uint8
	password  string
}

func (rap *RuleAndPassword) IsValid() bool {
	n := strings.Count(rap.password, string(rap.character))
	return n >= rap.val1 && n <= rap.val2
}

func (rap *RuleAndPassword) IsValidRule2() bool {
	c1 := rap.password[rap.val1-1]
	c2 := rap.password[rap.val2-1]
	return (c1 == rap.character || c2 == rap.character) && c1 != c2
}

func part1Tests() {
	line := mustParse([]byte("1-3 a: abcde"))
	expected := RuleAndPassword{
		val1:      1,
		val2:      3,
		character: 'a',
		password:  "abcde",
	}
	if line != expected {
		log.Fatal("not the same!")
	}
	test1 := RuleAndPassword{val1: 1, val2: 3, character: 'a', password: "abcde"}
	test2 := RuleAndPassword{val1: 1, val2: 3, character: 'b', password: "cdefg"}
	test3 := RuleAndPassword{val1: 2, val2: 9, character: 'c', password: "ccccccccc"}
	if test1.IsValid() != true {
		log.Fatal("expected", test1, "to be valid")
	}
	if test2.IsValid() != false {
		log.Fatal("expected", test2, "to be invalid")
	}
	if test3.IsValid() != true {
		log.Fatal("expected", test3, "to be valid")
	}
}

func part2Tests() {
	test1 := RuleAndPassword{val1: 1, val2: 3, character: 'a', password: "abcde"}
	test2 := RuleAndPassword{val1: 1, val2: 3, character: 'b', password: "cdefg"}
	test3 := RuleAndPassword{val1: 2, val2: 9, character: 'c', password: "ccccccccc"}
	if test1.IsValidRule2() != true {
		log.Fatal("expected", test1, "to be valid")
	}
	if test2.IsValidRule2() != false {
		log.Fatal("expected", test2, "to be invalid")
	}
	if test3.IsValidRule2() != false {
		log.Fatal("expected", test3, "to be invalid")
	}
}

func parseRulesAndPasswords(input []byte) []RuleAndPassword {
	var lines []RuleAndPassword

	rows := bytes.Split(input, []byte{'\n'})
	for _, r := range rows {
		line := mustParse(r)
		lines = append(lines, line)
	}
	return lines
}

func mustParse(r []byte) RuleAndPassword {
	var line RuleAndPassword
	_, err := fmt.Sscanf(string(r), "%d-%d %c: %s", &line.val1, &line.val2, &line.character, &line.password)
	if err != nil {
		log.Fatal(string(r), err)
	}
	return line
}

func part1(raps []RuleAndPassword) {
	var valid int
	for _, rap := range raps {
		if rap.IsValid() {
			valid++
		}
	}
	fmt.Println("valid passwords for part 1:", valid)
}

func part2(raps []RuleAndPassword) {
	var valid int
	for _, rap := range raps {
		if rap.IsValidRule2() {
			valid++
		}
	}
	fmt.Println("valid passwords for part 2:", valid)
}


func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	entries := parseRulesAndPasswords(input)

	part1Tests()
	part1(entries)

	part2Tests()
	part2(entries)

	log.Println("success")
}
