package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("templates/login.tmpl")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		name := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"name":     name,
			"password": password,
		})
	})

	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	r.POST("/json", func(c *gin.Context) {
		name := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"name":     name,
			"password": password,
		})
	})

	type msg struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Message string `json:"message"`
	}

	r.GET("/", func(c *gin.Context) {
		m := msg{
			Name:    "hello",
			Age:     18,
			Message: "go",
		}
		c.JSON(http.StatusOK, m)
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}

}
