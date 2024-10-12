package database

const userTable = (`
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(50) UNIQUE NOT NULL,
        password VARCHAR(50) NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL 
		)`)

const readHoursTable = (`
		CREATE TABLE IF NOT EXISTS read_hours (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL,
			hours INTEGER NOT NULL,
			timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`)

const allUsers = (`
		SELECT * FROM users
	`)

const deleteUserById = (`
		DELETE FROM users WHERE id = $1
	`)

const insertReadHours = (`
	INSERT INTO read_hours (user_id, hours) VALUES ($1, $2)
	`)
