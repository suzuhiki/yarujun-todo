package model

import (
	"fmt"
	"yarujun/app/database"
)

func CreateAccount(name string, password string) error {
	db := database.SetupDatabase()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	ins, err := tx.Prepare("INSERT INTO users (name, password) VALUES ($1, $2)")
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = ins.Exec(name, password)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func GetLoginInfo(name string) (password string, user_id string, err error) {
	db := database.SetupDatabase()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return "", "", err
	}

	rows, err := tx.Query("SELECT password, id FROM users WHERE name = $1", name)
	if err != nil {
		tx.Rollback()
		return "", "", err
	}

	var pass string
	var id string
	for rows.Next() {
		rows.Scan(&pass, &id)
	}
	if pass == "" {
		return "", "", fmt.Errorf("user not found")
	}

	if err := tx.Commit(); err != nil {
		return "", "", err
	}

	return pass, id, nil
}
