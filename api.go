package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Start(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "Server Started"})
}

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
		err := rows.Scan(&temp_emp.Id, &temp_emp.Name, &temp_emp.Gender, &temp_emp.FromDate,
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

// func POSTin(c *gin.Context) {
// 	// Container for Storing all rows in Array
// 	var Employee_details = []Employee{}

// 	var temp_emp Employee

// 	if err := c.BindJSON(&temp_emp); err != nil {
// 		return
// 	}

// 	Employee_details = append(Employee_details, temp_emp)

// 	//Database Connection
// 	DB := db_connection()

// 	/***********************************************/
// 	/***********************************************/
// 	/***********************************************/
// 	// close database
// 	defer DB.Close()

// 	//delete  this
// 	c.IndentedJSON(http.StatusCreated, Employee_details)

// 	// insert
// 	insertQry := `insert into "employee_details"("name", "gender") values($1,$2)`
// 	_, e := DB.Exec(insertQry, temp_emp.FullName, temp_emp.Gender)
// 	fmt.Printf(temp_emp.Gender)
// 	CheckError(e)

// }

func POST(c *gin.Context) {
	// Fname := c.PostForm("fname")

	// Extracting Values from submitted form
	name := c.Request.FormValue("name")
	gender := c.Request.FormValue("gender")
	from := c.Request.FormValue("startDate")
	to := c.Request.FormValue("tillDate")
	phone := c.Request.FormValue("phone")
	email := c.Request.FormValue("email")

	finalName := strings.TrimSpace(name)
	finalGender := gender

	// number validation
	resPhn, _ := regexp.MatchString(`((\+|\(|0)?\d{1,3})?((\s|\)|\-))?(\d{10})$`, phone)
	if !resPhn {
		c.JSON(400, gin.H{"error": "number not valid"})
		return
	}

	finalPhone, _ := strconv.ParseInt(phone, 10, 0)

	// temp storing file
	file, err := c.FormFile("resume")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//file validating
	if file.Size > 5000000 || file.Size == 0 {
		c.JSON(400, gin.H{"error": "file size can't exceed 5 MB"})
		return
	}

	fname := file.Filename

	resPdf, _ := regexp.MatchString(`\.pdf$`, fname)
	resPng, _ := regexp.MatchString(`\.png$`, fname)
	var fileName string
	if resPdf {
		fileName = (name + "_" + phone + ".pdf")
	} else if resPng {
		fileName = (name + "_" + phone + ".png")
	} else {
		c.JSON(400, gin.H{"error": "file type not acceptable"})
		return
	}

	// Define the path where the file will be saved

	filePath := filepath.Join("uploads", fileName)
	fmt.Println(filePath)
	// Save the file to the defined path
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(500, gin.H{"error": "Failed to save file"})
		return
	}

	//email validation
	resMail, _ := regexp.MatchString(`^[a-zA-Z0-9.!#$%&'*+=?^_{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`, email)
	if !resMail {
		c.JSON(400, gin.H{"error": "email not valid"})
		return
	}

	finalMail := email

	//date validation

	finalFrom, _ := time.Parse("2006-01-02", from)
	finalTo, _ := time.Parse("2006-01-02", to)

	// Save file metadata to database
	finalData := Employee{
		Name:     finalName,
		Gender:   finalGender,
		FromDate: finalFrom,
		ToDate:   finalTo,
		Phone:    finalPhone,
		Resume:   fileName,
		Email:    finalMail,
	}

	//DATABASE OPERATIONS
	/************************************/
	/************************************/
	/************************************/
	/************************************/
	// DB.Create(&finalData)
	c.JSON(201, gin.H{"message": "Details uploaded successfully", "Details": finalData})

	// c.JSON(200, gin.H{
	// 	"status":  "Submitted the form",
	// 	"message": "whoo",
	// 	"fname":   name,
	// 	"gender":  gender,
	// 	"email":   email,
	// 	"Dt1":     from,
	// 	"Dt":      to,
	// 	"phone":   phone,
	// })
}
