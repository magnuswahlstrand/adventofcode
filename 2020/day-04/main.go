package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func must(t bool, txt string) {
	if !t {
		log.Fatal("condition failed: ", txt)
	}
}

func isValidBirthYear(byr string) bool {
	return byr >= "1920" && byr <= "2002"
}

func isValidIssueYear(iyr string) bool {
	return iyr >= "2010" && iyr <= "2020"
}

func isValidExpirationYear(eyr string) bool {
	return eyr >= "2020" && eyr <= "2030"
}
func isValidEyeColor(ecl string) bool {
	switch ecl {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}
}
func isValidPassportID(pid string) bool {
	if len(pid) != 9 {
		return false
	}

	if _, err := strconv.Atoi(pid); err != nil {
		return false
	}
	return true
}

func isValidHeight(hgt string) bool {
	i := len(hgt) - 2
	if i < 0 {
		return false
	}

	switch v := hgt[i:]; v {
	case "in":
		return hgt[:i] >= "59" && hgt[:i] <= "76"
	case "cm":
		return hgt[:i] >= "150" && hgt[:i] <= "193"
	default:
		return false
	}
}

func isValidHairColor(hcl string) bool {
	if len(hcl) != 7 {
		return false
	}
	if hcl[0] != '#' {
		return false
	}
	_, err := hex.DecodeString(hcl[1:])
	if err != nil {
		return false
	}

	return true
}

func part2Tests() {
	must(isValidBirthYear("2002"), "")
	must(!isValidBirthYear("2003"), "")

	must(isValidIssueYear("2010"), "")
	must(!isValidIssueYear("2009"), "")

	must(isValidExpirationYear("2020"), "")
	must(!isValidExpirationYear("2031"), "")

	must(isValidHairColor("#123abc") == true, "")
	must(isValidHairColor("#123abz") == false, "")
	must(isValidHairColor("123abc") == false, "")

	must(isValidEyeColor("brn") == true, "")
	must(isValidEyeColor("wat") == false, "")

	must(isValidPassportID("000000001") == true, "")
	must(isValidPassportID("0123456789") == false, "")

	invalidInput := "eyr:1972 cid:100\nhcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926\n\niyr:2019\nhcl:#602927 eyr:1967 hgt:170cm\necl:grn pid:012533040 byr:1946\n\nhcl:dab227 iyr:2012\necl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277\n\nhgt:59cm ecl:zzz\neyr:2038 hcl:74454a iyr:2023\npid:3556412378 byr:2007"
	valid, total := countValid2(invalidInput)
	must(valid == 0, "valid == 0")
	must(total == 4, "total == 4")

	validInput := "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980\nhcl:#623a2f\n\neyr:2029 ecl:blu cid:129 byr:1989\niyr:2014 pid:896056539 hcl:#a97842 hgt:165cm\n\nhcl:#888785\nhgt:164cm byr:2001 iyr:2015 cid:88\npid:545766238 ecl:hzl\neyr:2022\n\niyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"
	valid, total = countValid2(validInput)
	must(valid == 4, "valid == 4")
	must(total == 4, "total == 4")
}

func part1Tests() {
	input := "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm\n\niyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\nhcl:#cfa07d byr:1929\n\nhcl:#ae17e1 iyr:2013\neyr:2024\necl:brn pid:760753108 byr:1931\nhgt:179cm\n\nhcl:#cfa07d eyr:2025 pid:166559648\niyr:2011 ecl:brn hgt:59in"
	valid, total := countValid(input)
	must(valid == 2, "valid == 2")
	must(total == 4, "valid == 4")
}

func part1(input string) {
	valid, _ := countValid(input)
	fmt.Println("valid passports for part 1:", valid)
}

func part2(input string) {
	valid, _ := countValid2(input)
	fmt.Println("valid passports for part 2:", valid)
}

func countValid(input string) (int, int) {
	var nValid, nTotal int

	lines := strings.Split(input, "\n")
	current := map[string]string{}
	for i, line := range lines {
		passportComplete := line == ""
		parts := strings.Split(line, " ")

		if !passportComplete {
			for _, part := range parts {
				current[part[:3]] = part[4:]
			}
		}

		if passportComplete || i == len(lines)-1 {
			nTotal++

			if !(current["byr"] == "" ||
				current["iyr"] == "" ||
				current["eyr"] == "" ||
				current["hgt"] == "" ||
				current["hcl"] == "" ||
				current["ecl"] == "" ||
				current["pid"] == "") {
				nValid++
			}

			current = map[string]string{}
		}
	}
	return nValid, nTotal
}

func countValid2(input string) (int, int) {
	var nValid, nTotal int

	lines := strings.Split(input, "\n")
	current := map[string]string{}
	for i, line := range lines {
		passportComplete := line == ""
		parts := strings.Split(line, " ")

		if !passportComplete {
			for _, part := range parts {
				current[part[:3]] = part[4:]
			}
		}

		if passportComplete || i == len(lines)-1 {
			nTotal++

			if isValidBirthYear(current["byr"]) &&
				isValidIssueYear(current["iyr"]) &&
				isValidExpirationYear(current["eyr"]) &&
				isValidHeight(current["hgt"]) &&
				isValidHairColor(current["hcl"]) &&
				isValidEyeColor(current["ecl"]) &&
				isValidPassportID(current["pid"]) {
				nValid++
			}

			current = map[string]string{}
		}
	}
	fmt.Println(nValid, nTotal)
	return nValid, nTotal
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1Tests()
	part1(string(input))

	part2Tests()
	part2(string(input))

	log.Println("success")
}
