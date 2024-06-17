package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	err := r.Run()
	if err != nil {
		fmt.Printf("r run error: %v", err)
		return
	}
}
