package model

import (
	"fmt"
	"yarujun/app/database"
)

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

func GetLoginInfo(name string) (password string, user_id string) {
	db := database.SetupDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT password, id FROM users WHERE name = $1", name)
	if err != nil {
		fmt.Println(err)
	}

	var pass string
	var id string
	for rows.Next() {
		rows.Scan(&pass, &id)
	}
	if pass == "" {
		fmt.Println("not found")
	}
	return pass, id
}
