package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pyprism/Hiren-Finder/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/create", controllers.AddSingleIndex)
	r.POST("/search", controllers.Search)
	r.POST("/update", controllers.UpdateIndex)
	r.Run("127.0.0.1:8000") // listen and serve on 0.0.0.0:8080
}
