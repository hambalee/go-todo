package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	db.Create(&User{Name: "Hello"})

	r := gin.Default()

	userHandler := UserHandler{db: db}
	r.GET("/users", userHandler.User)

	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.GET("/ping", pingpongHandler)
	r.Run(":8080")
}

type UserHandler struct {
	db *gorm.DB
}

func (h *UserHandler) User(c *gin.Context) {
	var u User
	h.db.First(&u)
	c.JSON(200, u)
}

func pingpongHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}