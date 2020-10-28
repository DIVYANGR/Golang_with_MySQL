package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main1() {
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
	insert, err := db.Query("INSERT INTO employee VALUES (1,'Divyang',10000,'76-Nilkanth','2020-09-28',1,'Golang' ),(2,'Aarsh',10000,'Junagadh','2020-09-28',1,'Golang' ),(3,'Dhiraj',10000,'surat','2020-09-28',1,'Golang' ),(4,'Shivang',10000,'bihar','2020-09-28',1,'Golang'),(5,'Yash',10000,'Gujarat','2020-09-28',1,'Golang'),(6,'Manish',12000,'Punjab','2020-09-30',2,'front-end'),(7,'Kapil',12000,'Tamilnadu','2020-10-14',2,'front-end'),(8,'Shreya',15000,'Punjab','2020-10-14',3,'MySQL'),(9,'Ahad',12000,'Tamilnadu','2020-09-14',3,'MySQL'),(10,'Ashish',15000,'Gujarat','2020-08-01',4,'Python')")

	if err != nil {
		panic(err.Error())
	}

	insert1, err := db.Query("INSERT INTO department VALUES (1,'Golang','Chennai'),(2,'fron-end','Chennai'),(3,'MySQL','Chennai'),(4,'Python','Chennai')")
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
	defer insert1.Close()

}
