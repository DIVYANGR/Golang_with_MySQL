package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:divyangk1998@tcp(127.0.0.1:3306)/kloudonedata")
	fmt.Println(db)
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	// perform a db.Query insert
	//insert, err := db.Query("INSERT INTO users VALUES ( 'Divyang' )")
	update, err := db.Query("INSERT INTO employee VALUES ( 3,'Dhiraj',10000,'surat','2020-09-28',1,'Golang' ),(4,'Shivang',10000,'bihar','2020-09-28',1,'golang')")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer update.Close()

}
