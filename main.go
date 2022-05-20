package main

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)
import "net/http"

//go:embed version.txt
var version string

func main() {
	r := gin.Default()

	prefix := ""
	if len(os.Args) > 1 {
		prefix = os.Args[1]
	}

	message := "Hello!"
	if len(os.Args) > 2 {
		message = strings.Join(os.Args[2:], " ")
	}

	api := r.Group(prefix)

	api.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})

	r.Run()
}
