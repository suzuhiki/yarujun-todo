package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type EMPLOYEE struct {
	ID     string
	NUMBER string
}

func main() {
	db_user, db_pass, db_port := loadEnv()

	db, err := sql.Open("postgres", "host=127.0.0.1 port="+db_port+" user="+db_user+" password="+db_pass+" dbname=yarujun sslmode=disable")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}

	// INSERT
	var task_id string
	title := "test"
	memo := "this is memo desu"
	deadline := time.Now()
	err = db.QueryRow("INSERT INTO tasks(title, memo, deadline) VALUES($1, $2, $3) RETURNING id", title, memo, deadline).Scan(&task_id)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(task_id)

	// SELECT
	rows, err := db.Query("SELECT title FROM tasks")

	if err != nil {
		fmt.Println(err)
	}

	var titles []string
	for rows.Next() {
		var title string
		rows.Scan(&title)
		titles = append(titles, title)
	}
	fmt.Printf("%v", titles)

	// start api server
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!!!!!!!!")
	})

	r.Run()
}

// .envを読み込む
func loadEnv() (string, string, string) {

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

	return db_user, db_pass, db_port
}
