package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	userAppService "github.com/tonouchi510/golang-ddd-layout/internal/application/users"
	"github.com/tonouchi510/golang-ddd-layout/internal/domain/models/users"
	userInfra "github.com/tonouchi510/golang-ddd-layout/internal/infrastructure/sqlboiler/users"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	ctx := context.Background()

	dsn := "test_user:password@tcp(localhost:3306)/golang_ddd_layout?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	boil.SetDB(db)

	ur, err := userInfra.NewUserRepository(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	us := users.NewUserService(ur)
	uf := userInfra.NewUserFactory(ctx)

	uas, err := userAppService.NewUserApplicationService(ur, us, uf)
	if err != nil {
		log.Fatalln(err.Error())
	}

	c1, err := userAppService.NewRegistorUserCommand("fuga")
	if err != nil {
		log.Fatalln(err.Error())
	}

	userData, err := uas.Registor(c1)
	if err != nil {
		log.Fatalln(err.Error())
	}
	c2, err := userAppService.NewGetUserCommand(userData.Id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	u, err := uas.Get(c2)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(u)
}
