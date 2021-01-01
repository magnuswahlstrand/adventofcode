package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"regexp"
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
	evaluatesTo := map[string]string{}
	re := evalAndRegex(rules, evaluatesTo)
	count, _ := validMessages(input, re)
	fmt.Println("answer part 1:", count)
}

func part2(input []byte) {
	rules := parseRules(input)
	regexs := calculateRegexPart2(rules)
	count, _ := validMessagesPart2(input, regexs)
	fmt.Println("answer part 1:", count)
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

func evalAndRegex(rules map[string]rule, evaluatesTo map[string]string) *regexp.Regexp {
	val := evaluatePart1(rules, evaluatesTo, "0")
	re, err := regexp.Compile("^" + val + "$")
	if err != nil {
		log.Fatal("failed compiling regex", err)
	}
	return re
}

func validMessages(input []byte, re *regexp.Regexp) (int, []string) {
	receievedMessages := strings.Split(strings.Split(string(input), "\n\n")[1], "\n")
	var count int
	var validMessages []string
	for _, message := range receievedMessages {
		if re.MatchString(message) {
			count++
			validMessages = append(validMessages, message)
		}
	}
	return count, validMessages
}

func validMessagesPart2(input []byte, regexs []*regexp.Regexp) (int, []string) {
	receievedMessages := strings.Split(strings.Split(string(input), "\n\n")[1], "\n")
	var count int
	var validMessages []string
	for _, message := range receievedMessages {
		for _, re := range regexs {
			if re.MatchString(message) {
				count++
				validMessages = append(validMessages, message)
				break
			}
		}
	}
	return count, validMessages
}

func evaluatePart1(rules map[string]rule, evaluatesTo map[string]string, name string) string {
	if val, ok := evaluatesTo[name]; ok {
		// Return from cache
		return val
	}

	r := rules[name]
	var result string
	switch {
	case strings.Contains(r.expression, "\""):
		// Is end rule "a" --> a
		result = r.expression[1:2]
	default:
		// Split into sub rules 1 2 | 3 4 --> 1 2, 3 4
		//fmt.Println("expr", r.expression)
		var parts []string
		for _, subSection := range strings.Split(r.expression, " | ") {

			//fmt.Println("sub sect", subSection)
			//// Evaluate each expression, e.g. 3 and 4
			var res2 string
			subRules := strings.Split(subSection, " ")
			for _, name := range subRules {
				res2 += evaluatePart1(rules, evaluatesTo, name)
				//fmt.Println(name)
			}
			//	//if i == 0 {
			//	//	res2 = append(res2, evaluatePart1(rules, evaluatesTo, name)...)
			//	//	continue
			//	//}
			//
			//	eval := evaluatePart1(rules, evaluatesTo, name)
			//	var tmpResult []string
			//	for _, s1 := range res2 {
			//		for _, s2 := range eval {
			//			tmpResult = append(tmpResult, s1+s2)
			//		}
			//	}
			//	res2 = tmpResult
			//}
			//result = append(result, res2...)
			parts = append(parts, res2)
		}
		result = "(" + strings.Join(parts, "|") + ")"
	}

	evaluatesTo[name] = result
	return result
}

func calculateRegexPart2(rules map[string]rule) []*regexp.Regexp {
	evaluatesTo := map[string]string{}
	rules42 := evaluatePart1(rules, evaluatesTo, "42")
	rules31 := evaluatePart1(rules, evaluatesTo, "31")
	var regexs []*regexp.Regexp
	for num8 := 1; num8 <= 6; num8++ {
		for num11 := 1; num11 <= 6; num11++ {
			// Prepend rule 42s and add suffix of rule31
			regString := fmt.Sprintf("^%s{%d}%s{%d}%s{%d}$", rules42, num8, rules42, num11, rules31, num11)
			re, err := regexp.Compile(regString)
			if err != nil {
				log.Fatal(err)
			}
			regexs = append(regexs, re)
		}
	}
	return regexs
}

// Guessed 348 --> too low
