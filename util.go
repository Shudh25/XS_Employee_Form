package main

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Custom DATA TYPE STRUCTURE
type Employee struct {
	gorm.Model
	Id       int       `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name"`
	Gender   string    `json:"gender"`
	FromDate time.Time `json:"from"`
	ToDate   time.Time `json:"to"`
	Phone    int64     `json:"phone"`
	Resume   string    `json:"resume"`
	Email    string    `json:"email"`
}

// For Routing the routes of SERVER
func routing() {
	router := gin.Default()
	router.Static("/assets", "./assets")

	router.GET("/", Start)
	router.POST("/sendData", POST)
	router.GET("/getData", GET)
	router.GET("/:file", GetFile)
	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(404, gin.H{}) })
	router.Run("localhost:8080")
}

// DATABASE CONNECTION
func db_connection() (db *gorm.DB) {
	//Connection String
	dsn := "host=localhost user=postgres password=lusifer25 dbname=api sslmode=disable TimeZone=Asia/Shanghai"

	// Open database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	CheckError(err)

	db.AutoMigrate(&Employee{})

	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
