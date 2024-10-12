package main

import (
	// "github.com/labstack/echo/v4"
	// echoSwagger "github.com/swaggo/echo-swagger"
	"database/sql"

	_ "github.com/lib/pq"

	"test/pkg/database"
)

func main() {

	db, err := sql.Open("postgres", "user=postgres password=calcal321 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	database.CreateUserTable(db)
	database.CreateReadHoursTable(db)
	// database.InsertUser(db, "test", "test", "test")

	database.GetAllUsers(db)
	database.DeleteUserById(db, 2)

	database.InsertReadHours(db, 1, 10)
	database.InsertReadHours(db, 1, 10)
	database.GetUserReadHours(db, 1)
}
