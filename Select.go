package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//Tag struct
type Tag struct {
	Employeeid   int    `json:"id"`
	Employeename string `json:"name"`
	//Employeesalary  int    `json:"salary"`
	//Employeeaddress string `json:"address"`
	//joiningdate     string `json:"date"`
	Departmentname string `json:"department"`
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:divyangk1998@tcp(127.0.0.1:3306)/kloudonedata")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	// perform a db.Query insert
	//insert, err := db.Query("INSERT INTO users VALUES ( 'Divyang' )")
	select1, err := db.Query("SELECT Employeeid,Employeename FROM employee")

	if err != nil {
		panic(err.Error())
	}

	for select1.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = select1.Scan(&tag.Employeeid, &tag.Employeename)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		fmt.Println("Employee Name:", tag.Employeename)
	}

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	select2, err := db.Query("SELECT Employeeid,Employeename,Departmentname FROM employee NATURAL JOIN department")

	if err != nil {
		panic(err.Error())
	}

	for select2.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = select2.Scan(&tag.Employeeid, &tag.Employeename, &tag.Departmentname)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		fmt.Println(tag.Employeename, tag.Employeeid, tag.Departmentname)
	}
	select3, err := db.Query("SELECT Employeeid,Employeename FROM employee Where Departmentid IN (SELECT departmentid FROM department WHERE departmentname='Golang')")

	if err != nil {
		panic(err.Error())
	}

	for select3.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = select3.Scan(&tag.Employeeid, &tag.Employeename)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		fmt.Println(tag.Employeename, tag.Employeeid)
	}

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer select1.Close()
	defer select2.Close()
	defer select3.Close()

}
