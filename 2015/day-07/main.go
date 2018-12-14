package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func dependenciesOk(registry map[string]uint16, rs []string) bool {

	for _, r := range rs {
		if _, ok := registry[r]; !ok {
			if r == "1" {
				log.Fatal("WTF")
			}
			return false
		}
	}

	return true
}

func run(operations []string, override bool) map[string]uint16 {
	registry := make(map[string]uint16)
	registry["0"] = 0
	registry["1"] = 1
	if override {
		registry["b"] = 3176
	}

	for i := 0; i < 1000; i++ {

		for _, op := range operations {
			// fmt.Println(op)
			var val uint16
			var r1, r2, rec string

			switch {
			case strings.Contains(op, "AND"):
				fmt.Sscanf(op, "%s AND %s -> %s", &r1, &r2, &rec)
				if dependenciesOk(registry, []string{r1, r2}) {
					registry[rec] = registry[r1] & registry[r2]
				}

			case strings.Contains(op, "OR"):
				fmt.Sscanf(op, "%s OR %s -> %s", &r1, &r2, &rec)
				if dependenciesOk(registry, []string{r1, r2}) {
					registry[rec] = registry[r1] | registry[r2]
				}

			case strings.Contains(op, "LSHIFT"):
				fmt.Sscanf(op, "%s LSHIFT %d -> %s", &r1, &val, &rec)
				if dependenciesOk(registry, []string{r1}) {
					registry[rec] = registry[r1] << val
				}

			case strings.Contains(op, "RSHIFT"):
				fmt.Sscanf(op, "%s RSHIFT %d -> %s", &r1, &val, &rec)
				if dependenciesOk(registry, []string{r1}) {
					registry[rec] = registry[r1] >> val
				}

			case strings.Contains(op, "NOT"):
				fmt.Sscanf(op, "NOT %s -> %s", &r1, &rec)
				if dependenciesOk(registry, []string{r1}) {
					registry[rec] = ^registry[r1]
				}

			default:
				fmt.Sscanf(op, "%s -> %s", &r1, &rec)

				if v, err := strconv.ParseUint(r1, 10, 16); err == nil {

					if rec == "b" && override {
						continue
					}
					registry[rec] = uint16(v)
					continue
				}

				if dependenciesOk(registry, []string{r1}) {
					registry[rec] = registry[r1]
				}
			}

		}
		if _, ok := registry["a"]; ok {
			break

		}
	}

	return registry
}
func runOperations(filename string, override bool) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	operations := strings.Split(string(content), "\n")

	results := run(operations, override)

	return fmt.Sprint("Result for a=", results["a"])
}

func main() {
	fmt.Println("Day 7 - 2015")
	fmt.Println("Test 1:", runOperations("example.txt", false))
	fmt.Println("Part 1:", runOperations("input.txt", false))
	fmt.Println("Part 1:", runOperations("input.txt", true))
}
