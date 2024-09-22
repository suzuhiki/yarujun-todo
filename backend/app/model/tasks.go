package model

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"yarujun/app/database"
	"yarujun/app/types"
)

func GetAllTask(user_id string) (datas []types.ShowTaskResponse) {
	db := database.SetupDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT id, title, deadline, done, waitlist_num FROM tasks WHERE user_id = $1", user_id)
	if err != nil {
		fmt.Println(err)
	}

	var tasks []types.ShowTaskResponse
	for rows.Next() {
		var id string
		var title string
		var deadline string
		var waitlist_num sql.NullInt64
		var done bool

		var task types.ShowTaskResponse

		rows.Scan(&id, &title, &deadline, &done, &waitlist_num)

		deadline = deadline[:10]

		if waitlist_num.Valid {
			task = types.ShowTaskResponse{Id: id, Title: title, Deadline: deadline, Done: done, Waitlist_num: strconv.FormatInt(waitlist_num.Int64, 10)}
		} else {
			task = types.ShowTaskResponse{Id: id, Title: title, Deadline: deadline, Done: done, Waitlist_num: ""}
		}

		tasks = append(tasks, task)
	}
	return tasks
}

func CreateTask(user_id string, data types.CreateTaskRequest) error {
	print("test")
	db := database.SetupDatabase()
	defer db.Close()

	if data.Waitlist_num <= 9 && data.Waitlist_num >= 0 {
		print("valid")
		ins, err := db.Prepare("INSERT INTO tasks (user_id, title, deadline, waitlist_num) VALUES ($1, $2, $3, $4)")
		if err != nil {
			fmt.Println(err)
			return err
		}
		_, err = ins.Exec(user_id, data.Title, data.Deadline, data.Waitlist_num)
		if err != nil {
			fmt.Println(err)
			return err
		}
	} else if data.Waitlist_num == -1 {
		print("-1")
		ins, err := db.Prepare("INSERT INTO tasks (user_id, title, deadline) VALUES ($1, $2, $3)")
		if err != nil {
			fmt.Println(err)
			return err
		}
		_, err = ins.Exec(user_id, data.Title, data.Deadline)
		if err != nil {
			fmt.Println(err)
			return err
		}
	} else {
		fmt.Println("waitlist_num is invalid")
		return errors.New("waitlist_num is invalid")
	}

	return nil
}

func UpdateDoneTask(user_id string, task_id string) error {
	db := database.SetupDatabase()
	defer db.Close()

	ins, err := db.Prepare("UPDATE tasks SET done = true WHERE user_id = $1 AND id = $2")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = ins.Exec(user_id, task_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
