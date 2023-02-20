package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db, err :=sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/golang2023")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Connected to the DataBases:")
}