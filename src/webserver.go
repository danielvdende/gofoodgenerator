package webserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Dish struct {
	Name            string
	Category        string
	PreparationTime string
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

	router.GET("/food/:day", func(c *gin.Context) {
		// TODO: this needs more params: needs full week menu to avoid duplicates. Assumes ruleset is static
		day := c.Param("day")
		c.String(http.StatusOK, "Hello %s", day)
	})

	router.Run("0.0.0.0:8080")
}
