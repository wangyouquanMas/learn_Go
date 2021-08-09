package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/something", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": "message",
			"nick":    "nick",
		})
		c.XML(200, gin.H{
			"status": "posted",
			"nick":   "nick",
		})
		c.String(http.StatusOK, "hello world")
	})

	if err := router.Run(); err != nil {
		log.Println("something error")
	}
}
