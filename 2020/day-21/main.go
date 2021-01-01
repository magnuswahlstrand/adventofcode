package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
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
	sum, _ := calcPart1And2(string(input))
	fmt.Println("answer part 1:", sum)
}

func part2(input []byte) {
	_, dangerous := calcPart1And2(string(input))
	fmt.Println("answer part 2:", dangerous)
}

func calcPart1And2(input string) (int, string) {
	possibleIngredients := map[string]map[string]bool{}
	countIngredients := map[string]int{}
	for _, line := range strings.Split(input, "\n") {
		// mxmxvkd kfcds sqjhc nhms (contains dairy, fish) -->
		// 1) mxmxvkd kfcds sqjhc nhms
		// 2) dairy, fish
		split := strings.Split(line[:len(line)-1], " (contains ")

		// Count ingredients
		newIngredients := strings.Split(split[0], " ")
		for _, ing := range newIngredients {
			countIngredients[ing]++
		}

		newAllergens := strings.Split(split[1], ", ")
		for _, allergen := range newAllergens {
			ingredientSet, exists := possibleIngredients[allergen]

			// Create new map if doesn't exist
			if !exists {
				ingredientsMap := map[string]bool{}
				for _, ingredient := range newIngredients {
					ingredientsMap[ingredient] = true
				}
				possibleIngredients[allergen] = ingredientsMap
				continue
			}

			// Perform intersection of two sets (map and slice)
			for ingredient := range ingredientSet {
				var found bool
				for _, newIngredients := range newIngredients {
					if ingredient == newIngredients {
						found = true
						break
					}
				}

				if !found {
					delete(possibleIngredients[allergen], ingredient)
				}
			}
		}
	}

	// Iterate to find all ingredients
	allergenToIngredient := map[string]string{}
	isMappedIngredient := map[string]bool{}
	for len(allergenToIngredient) != len(possibleIngredients) {
		for allergen, ingredients := range possibleIngredients {
			if allergenToIngredient[allergen] != "" {
				continue
			}

			if len(ingredients) != 1 {
				continue
			}

			// 1. Get single ingredient
			var ingredient string
			for ingr := range ingredients {
				ingredient = ingr
				break
			}

			// 2. Mark as solved
			allergenToIngredient[allergen] = ingredient
			isMappedIngredient[ingredient] = true // Quick fix

			// 3. Remove from possible ingredients
			for a, ingredients := range possibleIngredients {
				if allergen == a {
					continue
				}

				delete(ingredients, ingredient)
			}
		}
	}

	//for allergen, ingredients := range possibleIngredients {
	//	fmt.Println(allergen, ingredients)
	//}

	var sum int
	for ing, count := range countIngredients {
		if isMappedIngredient[ing] {
			continue
		}

		sum += count
	}

	// Calculate dangerous list
	var sortableAllergens []string
	for allergen := range allergenToIngredient {
		sortableAllergens = append(sortableAllergens, allergen)
	}
	sort.Slice(sortableAllergens, func(i, j int) bool {
		return sortableAllergens[i] < sortableAllergens[j]
	})
	var dangerousIngredients []string
	for _, allergen := range sortableAllergens {
		dangerousIngredients = append(dangerousIngredients, allergenToIngredient[allergen])
	}

	return sum, strings.Join(dangerousIngredients, ",")
}
