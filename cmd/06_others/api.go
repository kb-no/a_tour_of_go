package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/sample", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "サンプルAPI",
			"status":  200,
			"int":     1,
			"map": map[string]int{
				"key1": 100,
				"key2": 200,
			},
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}