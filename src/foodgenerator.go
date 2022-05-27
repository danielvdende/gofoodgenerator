package webserver

import (
	"encoding/csv"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
)

type Rule struct {
	PreparationTime []string
}

var daysOfWeek = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
var rules = map[string][]Rule{
	"Monday" : {{PreparationTime: []string{"Snel", "Medium Snel"}}},
	"Tuesday" : {{PreparationTime: []string{"Snel", "Medium Snel"}}},
	"Wednesday" : {{PreparationTime: []string{"Snel", "Medium Snel"}}},
	"Thursday" : {{PreparationTime: []string{"Snel", "Medium Snel"}}},
	"Friday" : {{PreparationTime: []string{"Snel", "Medium Snel"}}},
	"Saturday" : {{PreparationTime: []string{"Snel", "Medium Snel", "Lang"}}},
	"Sunday" : {{PreparationTime: []string{"Snel", "Medium Snel", "Lang"}}},
}

func contains(haystack []string, needle string) bool {
	for _, v := range haystack {
		if strings.EqualFold(needle, v) {
			return true
		}
	}
	return false
}

func checkRule(dish Dish, rule Rule) bool {
	// For now, only support preparation time filtering
	if contains(rule.PreparationTime, dish.PreparationTime) {
		return true
	}
	return false
}

func generateFoodForDay(day string, ruleSet []Rule,  allDishes []Dish) Dish {
	// For now, use dumb algorithm:
	// 1. grab random dish
	// 2. check if rules are satisfied
	// 3. if yes, done. If not, go back to 1.
	// This will not guarantee uniqueness in the list yet. That check would have to be done on the whole menu
	OUTER:
	for {
		randomIndex := rand.Intn(len(allDishes))
		dish := allDishes[randomIndex]
		for _, rule := range ruleSet {
			if !checkRule(dish, rule)	{
				continue OUTER
			}
		}
		return dish
	}
}

func CreateMenu(allDishes []Dish) map[string]Dish {
	menu := make(map[string]Dish)
	for day, ruleSet := range rules {
		menu[day] = generateFoodForDay(day, ruleSet, allDishes)
	}
	return menu
}

func GenerateFood() map[string]Dish {
	allDishes := loadDishes()
	menu := CreateMenu(allDishes)
	//db := InitDb()
	//defer db.Close()
	//db.Create(allDishes)

	//db.Take(&user)
	//println(user.Name, user.Category, user.PreparationTime)

	return menu
}

func loadDishes() []Dish {
	pathToDishes := "food.csv"
	var dishes []Dish
	f, err := os.Open(pathToDishes)
	if err != nil {
		log.Fatal("Unable to read input file " + pathToDishes, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	// Read the first line to skip the header
	csvReader.Read()
	for  {
		record, error := csvReader.Read()
		if error == io.EOF {
			break
		}
		if error != nil {
			log.Fatal(error)
		}
		dishes = append(dishes, Dish{
			Name: record[0],
			Category: record[1],
			PreparationTime: record[2],
		})
	}
	return dishes
}
