package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
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
	rules := parseRules(input)

	// Iterate
	evaluatesTo := map[string][]string{}
	allowedMessages := evaluate(rules, evaluatesTo, "0")
	count := countAllowed(allowedMessages, input)
	fmt.Println("answer part 1:", count)
}

func part2(input []byte) {
	//fmt.Println("answer part 2:", sum)
	rules := parseRules(input)
	evaluatesTo := map[string][]string{}
	rules42 := evaluate(rules, evaluatesTo, "42")
	rules31 := evaluate(rules, evaluatesTo, "31")

	fmt.Println(rules42)
	fmt.Println(rules31)
	//rules0 := evaluate(rules, evaluatesTo, "0")
	//

	countAll, countOK := naiveFilterAndCount(input, rules42, rules31)
	fmt.Printf("%d/%d\n", countOK, countAll)
	//
	//for key, val := range evaluatesTo {
	//	fmt.Println(key, val)
	//}
}

type rule struct {
	expression string
}

func countAllowed(allowedMessages []string, input []byte) int {
	isAllowed := allowedMap(allowedMessages)
	receievedMessages := strings.Split(strings.Split(string(input), "\n\n")[1], "\n")
	var count int
	max := math.MinInt32
	for _, line := range receievedMessages {
		if len(line) > max {
			max = len(line)
		}
		if isAllowed[line] {
			count++
		}
	}
	fmt.Println(max)
	return count
}

func allowedMap(allowedMessages []string) map[string]bool {
	isAllowed := map[string]bool{}
	for _, v := range allowedMessages {
		isAllowed[v] = true
	}
	return isAllowed
}

func parseRules(input []byte) map[string]rule {
	rules := map[string]rule{}
	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}

		split := strings.SplitN(line, ": ", 2)
		ruleName := split[0]
		ruleExpression := split[1]
		rules[ruleName] = rule{
			expression: ruleExpression,
		}
	}
	return rules
}

func evaluate(rules map[string]rule, evaluatesTo map[string][]string, name string) []string {
	if val, ok := evaluatesTo[name]; ok {
		// Return from cache
		return val
	}

	r := rules[name]
	var result []string
	switch {
	case strings.Contains(r.expression, "\""):
		// Is end rule "a" --> a
		result = []string{r.expression[1:2]}
	default:
		// Split into sub rules 1 2 | 3 4 --> 1 2, 3 4
		for _, subSection := range strings.Split(r.expression, " | ") {

			// Evaluate each expression, e.g. 3 and 4
			var res2 []string
			subRules := strings.Split(subSection, " ")
			for i, name := range subRules {
				if i == 0 {
					res2 = append(res2, evaluate(rules, evaluatesTo, name)...)
					continue
				}

				eval := evaluate(rules, evaluatesTo, name)
				var tmpResult []string
				for _, s1 := range res2 {
					for _, s2 := range eval {
						tmpResult = append(tmpResult, s1+s2)
					}
				}
				res2 = tmpResult
			}
			result = append(result, res2...)
		}
	}

	evaluatesTo[name] = result
	return result
}

//func evaluate2(rules map[string]rule, evaluatesTo map[string][]string, currentRule string, s string) bool {
//	fmt.Println(s)
//	if val, ok := evaluatesTo[currentRule]; ok {
//		// Return from cache
//		for _, v := range val {
//			if s == v {
//				return true
//			}
//		}
//		return false
//	}
//
//	r := rules[currentRule]
//	switch {
//	case strings.Contains(r.expression, "\""):
//		// Is end rule "a" --> a
//		result = []string{r.expression[1:2]}
//	default:
//		// Split into sub rules 1 2 | 3 4 --> 1 2, 3 4
//		for _, subSection := range strings.Split(r.expression, " | ") {
//
//			// Evaluate each expression, e.g. 3 and 4
//			var res2 []string
//			subRules := strings.Split(subSection, " ")
//			for i, name := range subRules {
//				if i == 0 {
//					res2 = append(res2, evaluate(rules, evaluatesTo, name)...)
//					continue
//				}
//
//				eval := evaluate(rules, evaluatesTo, name)
//				var tmpResult []string
//				for _, s1 := range res2 {
//					for _, s2 := range eval {
//						tmpResult = append(tmpResult, s1+s2)
//					}
//				}
//				res2 = tmpResult
//			}
//			result = append(result, res2...)
//		}
//	}
//	return false
//	//
//	//
//	//var result []string
//	//switch {
//	//case strings.Contains(r.expression, "\""):
//	//	// Is end rule "a" --> a
//	//	result = []string{r.expression[1:2]}
//	//default:
//	//	// Split into sub rules 1 2 | 3 4 --> 1 2, 3 4
//	//	for _, subSection := range strings.Split(r.expression, " | ") {
//	//
//	//		// Evaluate each expression, e.g. 3 and 4
//	//		var res2 []string
//	//		subRules := strings.Split(subSection, " ")
//	//		for i, name := range subRules {
//	//			if i == 0 {
//	//				res2 = append(res2, evaluate(rules, evaluatesTo, name)...)
//	//				continue
//	//			}
//	//
//	//			eval := evaluate(rules, evaluatesTo, name)
//	//			var tmpResult []string
//	//			for _, s1 := range res2 {
//	//				for _, s2 := range eval {
//	//					tmpResult = append(tmpResult, s1+s2)
//	//				}
//	//			}
//	//			res2 = tmpResult
//	//		}
//	//		result = append(result, res2...)
//	//	}
//	//}
//	//
//	//for _, val := range result {
//	//	fmt.Println(len(val))
//	//}
//	//
//	//evaluatesTo[currentRule] = result
//	//return result
//}

func naiveFilterAndCount(input []byte, rules42 []string, rules31 []string) (int, int) {
	receievedMessages := strings.Split(strings.Split(string(input), "\n\n")[1], "\n")
	countAll := len(receievedMessages)
	countOK := 0

	for _, message := range receievedMessages {
		// Count number of suffixes from 31 occur

		// Try trimming any of the 8 prefixes, and search from there
		var ok, suffixOK bool
		for _, prefix := range rules42 {
			for _, prefix2 := range rules42 {
				if strings.HasPrefix(message, prefix+prefix2) {
					ok = true
					break
				}
			}
		}

		for _, suffix := range rules31 {
			if strings.HasSuffix(message, suffix) {
				suffixOK = true
				break
			}
		}

		if ok && suffixOK {
			countOK++
		}
	}
	return countAll, countOK
}
