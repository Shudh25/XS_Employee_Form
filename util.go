package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Custom DATA TYPE STRUCTURE
type Employee struct {
	Id        string `json:"id"`
	FullName  string `json:"FullName"`
	Gender    string `json:"Gender"`
	StartDate string `json:"StartDate"`
	ToDate    string `json:"ToDate"`
	Phone     string `json:"Phone"`
	Resume    string `json:"Resume"`
	Email     string `json:"Email"`
}

// For Routing the routes of SERVER
func routing() {
	router := gin.Default()
	router.Static("/assets", "./assets")

	// router.POST("/sendData", POST)
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
