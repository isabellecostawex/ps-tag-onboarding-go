package postgresql

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:123456@postgres:5432/users_database?sslmode=disable")
	if err != nil {
		return err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        first_name VARCHAR(100) NOT NULL,
        last_name VARCHAR(100) NOT NULL,
        email VARCHAR(250) NOT NULL,
        age INT NOT NULL
    )`)

	if err != nil {
		return err
	}
	return nil
}
