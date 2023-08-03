package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GET(c *gin.Context) {

	// Container for Storing all rows in Array
	var Employee_details = []Employee{}

	var temp_emp Employee

	//Database Connection
	DB := db_connection()
	rows, err := DB.Query("select * from employee_details")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//Extracting Rows from DB
	for rows.Next() {
		err := rows.Scan(&temp_emp.Id, &temp_emp.FullName, &temp_emp.Gender, &temp_emp.StartDate,
			&temp_emp.ToDate, &temp_emp.Phone, &temp_emp.Resume, &temp_emp.Email)

		if err != nil {
			log.Fatal(err)
		}
		//Adding data to Employee_details array
		Employee_details = append(Employee_details, temp_emp)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// Prints the Json on page
	c.IndentedJSON(http.StatusOK, Employee_details)
}

func POST(c *gin.Context) {
	// Container for Storing all rows in Array
	// var Employee_details = []Employee{}

	var temp_emp Employee

	if err := c.BindJSON(&temp_emp); err != nil {
		return
	}

	//Database Connection
	DB := db_connection()

	/***********************************************/
	/***********************************************/
	/***********************************************/
	// close database
	// defer db.Close()

	//delete  this
	// Employee_details = append(Employee_details, temp_emp)
	c.IndentedJSON(http.StatusCreated, temp_emp)

	// insert
	insertQry := `insert into "employee_details"("fname", "lname") values($1,$2)`
	_, e := DB.Exec(insertQry, temp_emp.FullName, temp_emp.Gender)
	CheckError(e)

}
