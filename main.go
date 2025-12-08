package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var db *gorm.DB

type Todo struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	Text   string `json:"title"`
	Status string `json:"status"`
}

func initMysql() (err error) {
	dsn := "root:password@(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.DB().Ping()
	return err
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	err := initMysql()
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	db.AutoMigrate(&Todo{})

	v1Group := r.Group("v1")
	{
		// Scan all todo list
		v1Group.GET("/todo", func(c *gin.Context) {
			var todos []Todo
			if err = db.Find(&todos).Error; err != nil {
				panic(err)
			} else {
				c.JSON(http.StatusOK, todos)
			}
		})
		// Create todo
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo Todo
			err := c.BindJSON(&todo) // 接收从前端发来的json
			if err != nil {
				panic(err)
			}

			if err = db.Create(&todo).Error; err != nil { // 返回响应
				panic(err)
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		// Update todo by id
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			var todo Todo
			id, ok := c.Params.Get("id")
			if !ok {
				return
			}
			if err = db.Where("id = ?", id).First(&todo).Error; err != nil {
				panic(err)
			}
			c.BindJSON(&todo)
			if err = db.Save(&todo).Error; err != nil {
				panic(err)
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		// Delete todo by id
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				return
			}
			if err = db.Where("id = ?", id).Delete(&Todo{}).Error; err != nil {
				panic(err)
			} else {
				c.JSON(http.StatusOK, gin.H{"id": id})
			}
		})
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	err = r.Run()
	if err != nil {
		return
	}
}
