package model

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"yarujun/app/database"
	"yarujun/app/types"
)

func GetAllTask(user_id string, sort string, filter string) (datas []types.ShowTaskResponse, return_err error) {
	db := database.SetupDatabase()
	defer db.Close()

	var rows *sql.Rows
	var err error
	if filter == "waitlist" {
		rows, err = db.Query("SELECT id, title, deadline, done, waitlist_num FROM tasks WHERE user_id = $1 AND waitlist_num IS NOT NULL ORDER BY waitlist_num", user_id)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

	} else if sort == "deadline" {
		rows, err = db.Query("SELECT id, title, deadline, done, waitlist_num FROM tasks WHERE user_id = $1 ORDER BY done, deadline, waitlist_num", user_id)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else if sort == "waitlist_num" {
		rows, err = db.Query("SELECT id, title, deadline, done, waitlist_num FROM tasks WHERE user_id = $1 ORDER BY done, waitlist_num, deadline", user_id)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
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
	return tasks, nil
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

func UpdateDoneTask(user_id string, task_id string, value bool) error {
	db := database.SetupDatabase()
	defer db.Close()

	ins, err := db.Prepare("UPDATE tasks SET done = $1 WHERE user_id = $2 AND id = $3")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = ins.Exec(value, user_id, task_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// タスクの削除
func DeleteTask(user_id string, task_id string) error {
	db := database.SetupDatabase()
	defer db.Close()

	ins, err := db.Prepare("SELECT waitlist_num FROM tasks WHERE user_id = $1 AND id = $2")
	if err != nil {
		fmt.Println(err)
		return err
	}
	var waitlist_num sql.NullInt64
	err = ins.QueryRow(user_id, task_id).Scan(&waitlist_num)
	if err != nil {
		fmt.Println(err)
		return err
	}

	ins, err = db.Prepare("DELETE FROM tasks WHERE user_id = $1 AND id = $2")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = ins.Exec(user_id, task_id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if waitlist_num.Valid {
		// waitlistに含まれていた場合は、その番号より後ろのwaitlist_numを-1する
		ins, err = db.Prepare("UPDATE tasks SET waitlist_num = waitlist_num - 1 WHERE user_id = $1 AND waitlist_num > $2")
		if err != nil {
			fmt.Println(err)
			return err
		}
		_, err = ins.Exec(user_id, waitlist_num.Int64)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

// タスクをやる順リストの末尾に追加する
func AddWaitlist(user_id string, task_id string) error {
	db := database.SetupDatabase()
	defer db.Close()

	// 既にwaitlist_numが設定されているか確認
	ins, err := db.Prepare("SELECT waitlist_num FROM tasks WHERE user_id = $1 AND id = $2")
	if err != nil {
		fmt.Println(err)
		return err
	}
	var waitlist_num sql.NullInt64
	err = ins.QueryRow(user_id, task_id).Scan(&waitlist_num)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if waitlist_num.Valid {
		return errors.New("waitlist_num is already set")
	}

	// waitlist_numの最大値を取得
	ins, err = db.Prepare("SELECT MAX(waitlist_num) FROM tasks WHERE user_id = $1")
	if err != nil {
		fmt.Println(err)
		return err
	}
	var max_waitlist_num sql.NullInt64
	err = ins.QueryRow(user_id).Scan(&max_waitlist_num)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if max_waitlist_num.Valid {
		if max_waitlist_num.Int64 == 9 {
			// waitlist_numが9の場合は、既存のwaitlist_num==9のタスクをnullにセット
			ins, err := db.Prepare("UPDATE tasks SET waitlist_num = $1 WHERE waitlist_num = 9 AND user_id = $2")
			if err != nil {
				fmt.Println(err)
				return err
			}
			_, err = ins.Exec(sql.NullInt32{}, user_id)
			if err != nil {
				fmt.Println(err)
				return err
			}

			ins, err = db.Prepare("UPDATE tasks SET waitlist_num = $1 WHERE user_id = $2 AND id = $3")
			if err != nil {
				fmt.Println(err)
				return err
			}
			_, err = ins.Exec(9, user_id, task_id)
			if err != nil {
				fmt.Println(err)
				return err
			}
		} else {
			// waitlist_numが9未満の場合は、waitlist_numを+1してセット
			ins, err := db.Prepare("UPDATE tasks SET waitlist_num = $1 WHERE user_id = $2 AND id = $3")
			if err != nil {
				fmt.Println(err)
				return err
			}
			_, err = ins.Exec(max_waitlist_num.Int64+1, user_id, task_id)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	} else {
		// waitlist_numが設定されていない場合は、waitlist_num=0(リストの先頭)をセット
		ins, err := db.Prepare("UPDATE tasks SET waitlist_num = $1 WHERE user_id = $2 AND id = $3")
		if err != nil {
			fmt.Println(err)
			return err
		}
		_, err = ins.Exec(0, user_id, task_id)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

// waitlist_numのすべてを更新する
func ReorderWaitlist(user_id string, task_ids []int) error {
	db := database.SetupDatabase()
	defer db.Close()

	ins, err := db.Prepare("UPDATE tasks SET waitlist_num = $1 WHERE user_id = $2 AND id = $3")
	if err != nil {
		fmt.Println(err)
		return err
	}
	// waitlist_numをnullにセット
	for _, task_id := range task_ids {
		_, err = ins.Exec(sql.NullInt32{}, user_id, task_id)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	// waitlist_numがすべてnullになっているか確認
	ins2, err2 := db.Prepare("SELECT COUNT(*) FROM tasks WHERE user_id = $1 AND waitlist_num IS NOT NULL")
	if err2 != nil {
		fmt.Println(err2)
		return err2
	}
	var count int
	err2 = ins2.QueryRow(user_id).Scan(&count)
	if err2 != nil {
		fmt.Println(err2)
		return err2
	}
	if count != 0 {
		return errors.New("provided task_ids not include all tasks in waitlist")
	}

	// waitlist_numを更新
	for i, task_id := range task_ids {
		_, err = ins.Exec(i, user_id, task_id)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}
