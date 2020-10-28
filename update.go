package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main2() {
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
	update, err := db.Query("UPDATE employee SET Employeesalary = 20000 WHERE Departmentname = 'Golang' ")
	if err != nil {
		panic(err.Error())
	}

	update1, err := db.Query("UPDATE employee SET Departmentname = 'Python',Departmentid = 4 WHERE Employeename IN ('Shivang','Yash') ")

	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer update.Close()
	defer update1.Close()

}
