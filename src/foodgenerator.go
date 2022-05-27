package webserver

import (
	"strings"
)

type Rule struct {
	PreparationTime []string
}

var rules = map[string][]Rule{
	"Monday":    {{PreparationTime: []string{"Snel", "Medium Snel"}}},
	"Tuesday":   {{PreparationTime: []string{"Snel", "Medium Snel"}}},
	"Wednesday": {{PreparationTime: []string{"Snel", "Medium Snel"}}},
	"Thursday":  {{PreparationTime: []string{"Snel", "Medium Snel"}}},
	"Friday":    {{PreparationTime: []string{"Snel", "Medium Snel"}}},
	"Saturday":  {{PreparationTime: []string{"Snel", "Medium Snel", "Lang"}}},
	"Sunday":    {{PreparationTime: []string{"Snel", "Medium Snel", "Lang"}}},
}

func contains(haystack []string, needle string) bool {
	// TODO: this can/should be removed, leaving it in for now as it's the only tested function, so it's a nice example
	for _, v := range haystack {
		if strings.EqualFold(needle, v) {
			return true
		}
	}
	return false
}

func generateFoodForDay(ruleSet []Rule) Dish {
	// For now, use dumb algorithm:
	// 1. grab random dish
	// 2. check if rules are satisfied
	// 3. if yes, done. If not, go back to 1.
	// TODO: This will not guarantee uniqueness in the list yet. That check would have to be done on the whole menu
	// TODO: this is now hardcoded for a single rule. This should change.
	return GetDishWithConditions(ruleSet[0])
}

func CreateMenu() map[string]Dish {
	menu := make(map[string]Dish)
	for day, ruleSet := range rules {
		menu[day] = generateFoodForDay(ruleSet)
	}
	return menu
}

func GenerateFood() map[string]Dish {
	menu := CreateMenu()
	return menu
}
