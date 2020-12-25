package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const testInput = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

const testInput2 = `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`

func TestParseColors(t *testing.T) {
	inner, outers := parseColors("light red bags contain 1 bright white bag, 2 muted yellow bags.")
	require.Equal(t, "light red", inner)
	require.Len(t, outers, 2)
	require.Equal(t, "bright white", outers[0])
	require.Equal(t, "muted yellow", outers[1])

	inner, outers = parseColors("dotted black bags contain no other bags.")
	require.Equal(t, "dotted black", inner)
	require.Len(t, outers, 0)

	inner, outers = parseColors("bright white bags contain 1 shiny gold bag.")
	require.Equal(t, "bright white", inner)
	require.Len(t, outers, 1)
	require.Equal(t, "shiny gold", outers[0])
}

func TestParseRules(t *testing.T) {
	rules := parseRules(testInput)
	require.Contains(t, rules["shiny gold"], "bright white")
	require.Contains(t, rules["shiny gold"], "muted yellow")
	require.Contains(t, rules["muted yellow"], "light red")
}

func TestParseRules2(t *testing.T) {
	rules := parseRules2(testInput2)
	shinyGold := rules["shiny gold"]
	require.Contains(t, shinyGold, "dark red")
	require.Equal(t, shinyGold["dark red"], 2)
}

func TestFindBags(t *testing.T) {
	rules := parseRules(testInput)
	colors := map[string]bool{}
	findColorsThatLeadToColor("shiny gold", colors, rules)
	require.Len(t, colors, 4)
}
func TestFindBags2(t *testing.T) {
	rules := parseRules2(testInput2)
	n := countBags("shiny gold", rules) - 1
	require.Equal(t, 126, n)
}
