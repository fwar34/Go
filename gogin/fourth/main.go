package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type student struct {
	Name string
	Age  uint32
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	stu1 := &student{"Li", 10}
	stu2 := &student{"Wang", 11}
	r.GET("/arr", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gin.H{
			"title":  "Gin",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	r.Run(":9000")
}
