package main

import (
	_ "embed"
	"github.com/gin-gonic/gin"
)
import "net/http"

//go:embed version.txt
var version string

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "world",
		})
	})

	r.Run()
}
