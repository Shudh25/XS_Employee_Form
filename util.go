package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
	router.Run("localhost:8080")
}

// DATABASE CONNECTION
const (
	host     = "localhost"
	user     = "postgres"
	password = "lusifer25"
	dbname   = "xenonstack_db"
)

func db_connection() (db *sql.DB) {
	//Connection String
	psqlconn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)

	// Open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	// defer db.Close()
	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
