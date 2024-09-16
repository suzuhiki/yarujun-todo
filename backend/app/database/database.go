package database

import (
	"database/sql"
	"fmt"

	"yarujun/app/env"
)

func SetupDatabase() *sql.DB {
	env := env.LoadEnv()
	db, err := sql.Open("postgres", "host=127.0.0.1 port="+env.DB_PORT+" user="+env.DB_USER+" password="+env.DB_PASSWORD+" dbname=yarujun sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}

	return db
}

func CloseDatabase(db *sql.DB) {
	db.Close()
}
