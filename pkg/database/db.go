package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func CreateUserTable(db *sql.DB) {
	_, err := db.Exec(userTable)

	if err != nil {
		panic(err)
	}
	fmt.Println("Table created successfully")
}

func CreateReadHoursTable(db *sql.DB) {
	_, err := db.Exec(readHoursTable)

	if err != nil {
		panic(err)
	}
	fmt.Println("Table created successfully")
}

func InsertUser(db *sql.DB, username, password, email string) {
	insertUserQuery := `
	INSERT INTO users (username, password, email)
	VALUES ($1, $2, $3);`

	_, err := db.Exec(insertUserQuery, username, password, email)
	if err != nil {
		panic(err)
	}
	fmt.Println("User inserted successfully")
}

func GetAllUsers(db *sql.DB) {
	rows, err := db.Query(allUsers)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username, password, email string
		err = rows.Scan(&id, &username, &password, &email)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, username, password, email)
	}
}

func DeleteUserById(db *sql.DB, id int) {
	_, err := db.Exec(deleteUserById, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("User deleted successfully")
}

func InsertReadHours(db *sql.DB, userId, hours int) {
	_, err := db.Exec(insertReadHours, userId, hours)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read hours inserted successfully")
}

func GetUserReadHours(db *sql.DB, userId int) {
	rows, err := db.Query(`SELECT * FROM read_hours WHERE user_id = $1`, userId)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var userId int
		var hours int
		var timestamp string
		err = rows.Scan(&id, &userId, &hours, &timestamp)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, userId, hours, timestamp)
	}
}
