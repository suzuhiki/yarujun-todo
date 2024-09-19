package model

import (
	"fmt"
	"yarujun/app/database"
)

type TaskEntity struct {
	Title        string
	Memo         string
	Deadline     string
	waitlist_num string
	work_time    string
}

func GetAll() (datas []TaskEntity) {
	db := database.SetupDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT title, memo, deadline, waitlist_num, work_time FROM tasks")
	if err != nil {
		fmt.Println(err)
	}

	var tasks []TaskEntity
	for rows.Next() {
		var title string
		var memo string
		var deadline string
		var waitlist_num string
		var work_time string

		rows.Scan(&title, &memo, &deadline, &waitlist_num, &work_time)
		task := TaskEntity{Title: title, Memo: memo, Deadline: deadline, waitlist_num: waitlist_num, work_time: work_time}
		tasks = append(tasks, task)
	}
	fmt.Printf("%v", tasks)
	return tasks
}
