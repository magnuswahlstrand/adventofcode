package main

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput2 = []byte(`0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`)

var testExpected2 = []string{"aaaabb", "aaabab", "abbabb", "abbbab", "aabaab", "aabbbb", "abaaab", "ababbb"}

var testExpected = []string{"aab", "aba"}

func TestPart1(t *testing.T) {
	// Setup
	rules := parseRules(testInput)

	// Iterate
	evaluatesTo := map[string][]string{}
	val := evaluate(rules, evaluatesTo, "0")
	require.EqualValues(t, testExpected, val)
}

func TestPart1Attempt2(t *testing.T) {
	// Setup
	input := testInput2
	rules := parseRules(input)

	// Iterate
	evaluatesTo := map[string]string{}
	val := evaluatePart1(rules, evaluatesTo, "0")

	re, err := regexp.Compile(val)
	require.NoError(t, err)

	receievedMessages := strings.Split(strings.Split(string(input), "\n\n")[1], "\n")

	var count int
	var validMessages []string
	for _, message := range receievedMessages {
		if re.MatchString(message) {
			count++
			validMessages = append(validMessages, message)
		}
	}

	require.Contains(t, validMessages, "ababbb")
	require.Contains(t, validMessages, "abbbab")
}

var testInput = []byte(`0: 1 2
1: "a"
2: 1 3 | 3 1
3: "b"`)

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
		fmt.Println("expr", r.expression)
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

func TestPart1Full(t *testing.T) {
	// Setup
	input := testInput2
	rules := parseRules(input)

	// Iterate
	evaluatesTo := map[string][]string{}
	allowedMessages := evaluate(rules, evaluatesTo, "0")
	require.EqualValues(t, testExpected2, allowedMessages)

	count := countAllowed(allowedMessages, input)
	require.EqualValues(t, 2, count)
}

var part2TestInput = []byte(`42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1

abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba`)

func TestPart2Full(t *testing.T) {
	_ = []string{"bbabbbbaabaabba", "ababaaaaaabaaab", "ababaaaaabbbaba"}
	_ = []string{
		"bbabbbbaabaabba",
		"babbbbaabbbbbabbbbbbaabaaabaaa",
		"aaabbbbbbaaaabaababaabababbabaaabbababababaaa",
		"bbbbbbbaaaabbbbaaabbabaaa",
		"bbbababbbbaaaaaaaabbababaaababaabab",
		"ababaaaaaabaaab",
		"ababaaaaabbbaba",
		"baabbaaaabbaaaababbaababb",
		"abbbbabbbbaaaababbbbbbaaaababb",
		"aaaaabbaabaaaaababaa",
		"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa",
		"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba",
	}

	// Setup
	input := part2TestInput
	rules := parseRules(input)

	// Manually expand these rules 4 steps
	rules["8"] = rule{expression: "42 | 42 42 | 42 42 42"}
	rules["11"] = rule{expression: "42 31 | 42 Y 31"}
	//rules["8"] = rule{expression: "42 | 42 42 | 42 42 42 | 42 42 42 42 | 42 42 42 42"}
	//rules["11"] = rule{expression: "42 31 | 42 42 31 31 | 42 42 42 31 31 31 | 42 42 42 42 31 31 31 31"}

	evaluatesTo := map[string][]string{}

	//allowedMessages := evaluate2(rules, evaluatesTo, "11")
	//naiveFilterAndCount := countAllowed(allowedMessages, input)

	//for _, message := range receievedMessages {
	//	evaluate2(rules, evaluatesTo, "0", message)
	//}
	//require.Equal(t, 12, naiveFilterAndCount)
	//
	//evaluatesTo["8"] = []string{""}
	//evaluatesTo["11"] = []string{"Y"}
	rules42 := evaluate(rules, evaluatesTo, "42")
	rules31 := evaluate(rules, evaluatesTo, "31")

	fmt.Println(strings.Join(rules42, "|"))
	fmt.Println(strings.Join(rules31, "|"))
	//rules0 := evaluate(rules, evaluatesTo, "0")
	//

	countAll, countOK := naiveFilterAndCount(input, rules42, rules31)
	fmt.Printf("%d/%d\n", countOK, countAll)

	//for key, val := range evaluatesTo {
	//	fmt.Println(key, val)
	//}
	//var naiveFilterAndCount int
	//for _, message := range receievedMessages {
	//	// Check against possible combinations
	//	fmt.Println(rules0[0])
	//	for _, p := range rules0[0] {
	//		switch p {
	//		case 'X':
	//
	//		case 'Y':
	//
	//		default:
	//
	//		}
	//	}
	//
	//	//(		if isAllowed[message] {
	//	//			naiveFilterAndCount++
	//	//		}
	//	isMatch := true
	//	if isMatch {
	//		fmt.Println(message)
	//		naiveFilterAndCount++
	//	}
	//}

	//
	//evaluatesTo["8"] = []string{"8"}
	//evaluatesTo["11"] = []string{"11"}
	//allowedMessages := evaluate(rules, evaluatesTo, "0")
	//require.Equal(t, preExpected, allowedMessages)
	//
	//naiveFilterAndCount := countAllowed(allowedMessages, input)
	//require.EqualValues(t, 3, naiveFilterAndCount)
}
