package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	
	db, err := sql.Open("mysql", "root:divyangk1998@tcp(127.0.0.1:3306)/kloudonedata")
	fmt.Println(db)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	//insert, err := db.Query("INSERT INTO users VALUES ( 'Divyang' )")
	insert, err := db.Query("INSERT INTO employee VALUES ( 3,'Dhiraj',10000,'surat','2020-09-28',1,'Golang' ),(4,'Shivang',10000,'bihar','2020-09-28',1,'golang')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

}
