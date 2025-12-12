package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Ingredient struct {
	Name       string
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
}

type MixResult struct {
	BestScore int
	BestMix   []int
}

func Day15() {
	content := Read_file("./files/day15.txt")
	ingredients, err := parseIngredient(content)
	if err != nil {
		panic(err)
	}
	result := MaxScore(ingredients, 100)
	fmt.Printf("Día 15 - Parte 1: La mejor puntuación es %d con la mezcla %v\n", result.BestScore, result.BestMix)

	resultWithCalories := MaxScoreWithCalories(ingredients, 100, 500)
	fmt.Printf("Día 15 - Parte 2: La mejor puntuación con 500 calorías es %d con la mezcla %v\n", resultWithCalories.BestScore, resultWithCalories.BestMix)
}

func parseIngredient(content string) ([]Ingredient, error) {
	var items []Ingredient
	sc := bufio.NewScanner(strings.NewReader(content))

	for sc.Scan() {
		var item Ingredient
		_, err := fmt.Sscanf(sc.Text(), "%s capacity %d, durability %d, flavor %d, texture %d, calories %d",
			&item.Name, &item.Capacity, &item.Durability, &item.Flavor, &item.Texture, &item.Calories)
		if err != nil {
			return nil, fmt.Errorf("error al analizar línea %q: %w", sc.Text(), err)
		}
		item.Name = strings.TrimSuffix(item.Name, ":")
		items = append(items, item)
	}
	return items, nil
}

func MaxScore(ingredients []Ingredient, TotalTeaspoons int) MixResult {
	return maximize(ingredients, TotalTeaspoons, -1)
}

func MaxScoreWithCalories(ingredients []Ingredient, totalTeaspoons int, caloriesTarget int) MixResult {
	return maximize(ingredients, totalTeaspoons, caloriesTarget)
}

func maximize(ingredients []Ingredient, totalTeaspoons int, caloriesTarget int) MixResult {
	n := len(ingredients)
	best := MixResult{BestScore: 0, BestMix: make([]int, n)}
	currentMix := make([]int, n)

	var dfs func(idx, remaining, capSum, furSum, flaSum, texSum, calSum int)
	dfs = func(idx, remaining, capSum, furSum, flaSum, texSum, calSum int) {
		if idx == n-1 {
			// last ingredient takes the remaining teaspoons
			q := remaining
			capTot := capSum + ingredients[idx].Capacity*q
			furTot := furSum + ingredients[idx].Durability*q
			flaTot := flaSum + ingredients[idx].Flavor*q
			texTot := texSum + ingredients[idx].Texture*q
			calTot := calSum + ingredients[idx].Calories*q

			if caloriesTarget != -1 && calTot != caloriesTarget {
				return
			}
			if capTot < 0 || furTot < 0 || flaTot < 0 || texTot < 0 {
				return
			}
			score := capTot * furTot * flaTot * texTot
			if score > best.BestScore {
				best.BestScore = score
				copy(best.BestMix, currentMix)
				best.BestMix[idx] = q
			}
			return
		}
		ing := ingredients[idx]
		for q := 0; q <= remaining; q++ {
			currentMix[idx] = q
			dfs(idx+1, remaining-q,
				capSum+ing.Capacity*q,
				furSum+ing.Durability*q,
				flaSum+ing.Flavor*q,
				texSum+ing.Texture*q,
				calSum+ing.Calories*q)
		}
	}
	dfs(0, totalTeaspoons, 0, 0, 0, 0, 0)
	return best
}
