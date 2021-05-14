package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/huroshotoku/golang-ddd-layout/internal/domain/models/users"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	// Open handle to database like normal
	dsn := "test_user:password@tcp(localhost:3306)/golang_ddd_layout"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	boil.SetDB(db)

	userName, err := users.NewUserName("hoge")
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(userName)
	user, err := users.NewUserByName(userName)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(user)
}
