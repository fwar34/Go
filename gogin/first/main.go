package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello go")
	})
	if err := r.Run(":9000"); err != nil {
		log.Fatal(err.Error())
	}
}
