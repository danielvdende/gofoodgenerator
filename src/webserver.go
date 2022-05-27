package webserver

import (
	"github.com/gin-gonic/gin"
)

type Dish struct {
	Name            string
	Category        string
	PreparationTime string
}

type Menu struct {
	Monday    Dish
	Tuesday   Dish
	Wednesday Dish
	Thursday  Dish
	Friday    Dish
	Saturday  Dish
	Sunday    Dish
}

func Run() {
	router := gin.Default()
	router.StaticFile("/", "./frontend/index.html")

	router.GET("/food", func(c *gin.Context) {
		dishes := GenerateFood()
		c.JSON(200, gin.H{
			"status": "OK",
			"code":   200,
			"data":   dishes,
		})
	})

	router.POST("/food/:day", func(c *gin.Context) {
		var json map[string]Dish
		// TODO: this needs more params: needs full week menu to avoid duplicates. Assumes ruleset is static
		day := c.Param("day")
		c.BindJSON(&json)

		// TODO: no guarantee you won't get the same dish again
		newDish := generateFoodForDay(day, rules[day], allDishes)

		c.JSON(200, gin.H{
			"status": "OK",
			"code":   200,
			"data":   newDish,
		})
	})

	router.Run("0.0.0.0:8080")
}
