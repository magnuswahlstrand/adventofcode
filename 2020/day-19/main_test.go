package main

import (
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
	re := evalAndRegex(rules, evaluatesTo)

	count, validMessages := validMessages(input, re)

	require.Contains(t, validMessages, "ababbb")
	require.Contains(t, validMessages, "abbbab")
	require.EqualValues(t, 2, count)
}

var testInput = []byte(`0: 1 2
1: "a"
2: 1 3 | 3 1
3: "b"`)

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
	// Setup
	input := part2TestInput
	rules := parseRules(input)
	regexs := calculateRegexPart2(rules)
	count, _ := validMessagesPart2(input, regexs)
	require.Equal(t, 12, count)
}
