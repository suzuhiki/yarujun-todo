package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ENV struct {
	DB_USER     string
	DB_PASSWORD string
	DB_PORT     string
}

// .envを読み込む
func LoadEnv() ENV {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASSWORD")
	db_port := os.Getenv("DB_PORT")

	if db_user == "" || db_pass == "" || db_port == "" {
		fmt.Println("環境変数が設定されていません")
	}

	env := ENV{DB_USER: db_user, DB_PASSWORD: db_pass, DB_PORT: db_port}

	return env
}
