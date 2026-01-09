package main

import (
	"fmt"
	"sort"
)

type BrainrotMeme struct {
	Name       string
	TrendLevel int
	Category   string
	Views      float64
}

func FindTopTrending(memes []BrainrotMeme, minViews float64) []BrainrotMeme {
	var res []BrainrotMeme
	for _, meme := range memes {
		if meme.Views > minViews {
			res = append(res, meme)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].TrendLevel > res[j].TrendLevel
	})
	return res
}

func CalculateCategoryImpact(memes []BrainrotMeme) map[string]float64 {
	cat_views := make(map[string]float64)
	for _, meme := range memes {
		cat_views[meme.Category] += meme.Views
	}
	return cat_views
}

func FilterByComplexCondition(memes []BrainrotMeme) []string {
	var res []string
	for _, meme := range memes {
		if meme.TrendLevel >= 7 || (meme.Views > 50 && meme.Category == "Sigma") {
			res = append(res, meme.Name)
		}
	}
	return res
}

func main() {
	memes := []BrainrotMeme{
		{"Skibidi Toilet", 9, "Skibidi", 120.5},
		{"Sigma Male Grindset", 10, "Sigma", 85.2},
		{"Mewing Tutorial", 8, "Mewing", 45.1},
		{"Subo Bratik Dance", 6, "Subo Bratik", 30.7},
		{"TUNTUNTUNSAHUR", 7, "TUNTUNTUNSAHUR", 15.3},
		{"Ohio Final Boss", 5, "Other", 200.0},
		{"Gyatt Rizzler", 9, "Sigma", 67.8},
	}

	fmt.Println("--- Тестовые данные ---")
	for _, m := range memes {
		fmt.Printf("%s (Trend: %d, Cat: %s, Views: %.1fM)\n", m.Name, m.TrendLevel, m.Category, m.Views)
	}

	top := FindTopTrending(memes, 40)
	fmt.Printf("\n--- Топ трендовые (>40M просмотров) ---\n")
	for _, m := range top {
		fmt.Printf("%s (Trend: %d, Views: %.1fM)\n", m.Name, m.TrendLevel, m.Views)
	}

	cat_views := CalculateCategoryImpact(memes)
	fmt.Printf("\n--- Влияние категорий (просмотры, M) ---\n")
	for cat, views := range cat_views {
		fmt.Printf("%s: %.1f\n", cat, views)
	}

	filtered := FilterByComplexCondition(memes)
	fmt.Printf("\n--- По сложному условию (названия) ---\n")
	for _, name := range filtered {
		fmt.Println("- " + name)
	}
}
