package main

import (
	_ "embed"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strings"
)
import "net/http"

//go:embed version.txt
var version string

type Ping_Response struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}

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

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"name":    prefix,
		})
	})

	api.GET("/headers", func(c *gin.Context) {
		c.JSON(http.StatusOK, c.Request.Header)
	})

	api.GET("/hello", func(c *gin.Context) {
		resp, err := http.Get("http://phero-hidden:8080/ping")

		var ping_response = Ping_Response{"Unknown", "Unknown"}

		if err == nil {
			if err = json.NewDecoder(resp.Body).Decode(&ping_response); err == nil {
				log.Println("Unable to decode data form hidden phero server.", err)
			}
		} else {
			log.Println("Unable to make request to hidden phero server. ", err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  message,
			"upstream": ping_response,
		})
	})

	r.Run()
}
