package model

import (
	"fmt"
	"yarujun/app/database"
)

type UserEntity struct {
	Name     string
	Password string
}

func CreateAccount(name string, password string) error {
	db := database.SetupDatabase()
	defer db.Close()

	ins, err := db.Prepare("INSERT INTO users (name, password) VALUES ($1, $2)")
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = ins.Exec(name, password)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetPassword(name string) (password string) {
	db := database.SetupDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT password FROM users WHERE name = $1", name)
	if err != nil {
		fmt.Println(err)
	}

	var pass string
	for rows.Next() {
		rows.Scan(&pass)
	}
	return pass
}
