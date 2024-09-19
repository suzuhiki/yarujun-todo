package model

import (
	"fmt"
	"yarujun/app/database"
	"yarujun/app/types"
)

func GetAllTask() (datas []types.TaskEntity) {
	db := database.SetupDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT title, memo, deadline, waitlist_num, work_time FROM tasks")
	if err != nil {
		fmt.Println(err)
	}

	var tasks []types.TaskEntity
	for rows.Next() {
		var title string
		var memo string
		var deadline string
		var waitlist_num string

		rows.Scan(&title, &memo, &deadline, &waitlist_num)
		task := types.TaskEntity{Title: title, Memo: memo, Deadline: deadline, Waitlist_num: waitlist_num}
		tasks = append(tasks, task)
	}
	fmt.Printf("%v", tasks)
	return tasks
}

func CreateTask(user_id string, data types.CreateTaskRequest) error {
	db := database.SetupDatabase()
	defer db.Close()

	ins, err := db.Prepare("INSERT INTO tasks (user_id, title, memo, deadline, waitlist_num) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = ins.Exec(user_id, data.Title, data.Memo, data.Deadline, data.Waitlist_num)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
