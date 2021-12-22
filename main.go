package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Print("a")
	r := gin.Default()
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// r.LoadHTMLGlob("public/*.html")
	// r.StaticFile("/favicon.ico", "./resources/favicon.ico")
	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", nil)
	// })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
