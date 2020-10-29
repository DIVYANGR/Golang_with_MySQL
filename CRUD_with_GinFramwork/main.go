package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id         int    `json:id`
	Name       string `json:name`
	City       string `json:city`
	Department string `json:department`
	Salary     int    `json:salary`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "divyangk1998"
	dbName := "goblog"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
func main() {
	router := gin.Default()

	router.POST("/add", func(c *gin.Context) {

		name := c.Query("name")
		city := c.Query("city")
		department := c.Query("department")
		salary := c.Query("salary")

		c.JSON(200, gin.H{
			"name":       name,
			"city":       city,
			"department": department,
			"salary":     salary,
		})
		db := dbConn()
		insForm, err := db.Prepare("INSERT INTO Employee(name, city, department, salary) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city, department, salary)
		fmt.Printf("name: %s; city: %s; department: %s; salary: %d", name, city, department, salary)
	})

	router.GET("/GET", func(c *gin.Context) {
		id := c.Query("id")
		db := dbConn()
		selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", id)
		if err != nil {
			panic(err.Error())
		}
		var salary int
		var name, city, department string
		for selDB.Next() {

			err = selDB.Scan(&id, &name, &city, &department, &salary)
			if err != nil {
				panic(err.Error())
			}
		}
		fmt.Printf("name: %s; city: %s; department: %s; salary: %d", name, city, department, salary)

		c.JSON(200, gin.H{
			"id":         id,
			"name":       name,
			"city":       city,
			"department": department,
			"salary":     salary,
		})

	})
	router.GET("/delete", func(c *gin.Context) {
		name := c.Query("name")
		db := dbConn()
		delForm, err := db.Prepare("DELETE FROM Employee WHERE name=?")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(name)
		log.Println("DELETE")
		defer db.Close()

	})

	router.Run(":8080")
}
