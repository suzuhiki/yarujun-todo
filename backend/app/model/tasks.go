package model

import (
	"database/sql"
	"errors"
	"strconv"
	"yarujun/app/database"
	"yarujun/app/types"
)

func GetAllTask(user_id string, sort string, filter string) (datas []types.ShowTaskResponse, return_err error) {
	db := database.SetupDatabase()
	defer db.Close()

	var rows *sql.Rows
	var err error

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	if filter == "waitlist" {
		rows, err = tx.Query("SELECT id, title, deadline, done, waitlist_num FROM tasks WHERE user_id = $1 AND waitlist_num IS NOT NULL ORDER BY waitlist_num", user_id)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

	} else if sort == "deadline" {
		rows, err = tx.Query("SELECT id, title, deadline, done, waitlist_num FROM tasks WHERE user_id = $1 ORDER BY done, deadline, waitlist_num", user_id)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	} else if sort == "waitlist_num" {
		rows, err = tx.Query("SELECT id, title, deadline, done, waitlist_num FROM tasks WHERE user_id = $1 ORDER BY done, waitlist_num, deadline", user_id)
		if err != nil {
			tx.Rollback()
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

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func CreateTask(user_id string, data types.CreateTaskRequest) error {
	db := database.SetupDatabase()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if data.Waitlist_num <= 9 && data.Waitlist_num >= 0 {
		ins, err := tx.Prepare("INSERT INTO tasks (user_id, title, deadline, waitlist_num) VALUES ($1, $2, $3, $4)")
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = ins.Exec(user_id, data.Title, data.Deadline, data.Waitlist_num)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else if data.Waitlist_num == -1 {
		ins, err := tx.Prepare("INSERT INTO tasks (user_id, title, deadline) VALUES ($1, $2, $3)")
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = ins.Exec(user_id, data.Title, data.Deadline)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		tx.Rollback()
		return errors.New("waitlist_num is invalid")
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func UpdateDoneTask(user_id string, task_id string, value bool) error {
	db := database.SetupDatabase()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	ins, err := tx.Prepare("UPDATE tasks SET done = $1 WHERE user_id = $2 AND id = $3")
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = ins.Exec(value, user_id, task_id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// タスクの削除
func DeleteTask(user_id string, task_id string) error {
	db := database.SetupDatabase()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// 削除対象のwaitlist_numを取得
	ins, err := tx.Prepare("SELECT waitlist_num FROM tasks WHERE user_id = $1 AND id = $2")
	if err != nil {
		tx.Rollback()
		return err
	}
	var waitlist_num sql.NullInt64
	err = ins.QueryRow(user_id, task_id).Scan(&waitlist_num)
	if err != nil {
		tx.Rollback()
		return err
	}

	// タスクを削除
	ins, err = tx.Prepare("DELETE FROM tasks WHERE user_id = $1 AND id = $2")
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = ins.Exec(user_id, task_id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if waitlist_num.Valid {
		// waitlistに含まれていた場合は、その番号より後ろのwaitlist_numを-1する
		ins, err = tx.Prepare("UPDATE tasks SET waitlist_num = waitlist_num - 1 WHERE user_id = $1 AND waitlist_num > $2")
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = ins.Exec(user_id, waitlist_num.Int64)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// タスクをやる順リストの末尾に追加する
func AddWaitlist(user_id string, task_id string) error {
	db := database.SetupDatabase()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// 既にwaitlist_numが設定されているか確認
	ins, err := tx.Prepare("SELECT waitlist_num FROM tasks WHERE user_id = $1 AND id = $2")
	if err != nil {
		tx.Rollback()
		return err
	}
	var waitlist_num sql.NullInt64
	err = ins.QueryRow(user_id, task_id).Scan(&waitlist_num)
	if err != nil {
		tx.Rollback()
		return err
	}
	if waitlist_num.Valid {
		return errors.New("waitlist_num is already set")
	}

	// waitlist_numの最大値を取得
	ins, err = tx.Prepare("SELECT MAX(waitlist_num) FROM tasks WHERE user_id = $1")
	if err != nil {
		tx.Rollback()
		return err
	}
	var max_waitlist_num sql.NullInt64
	err = ins.QueryRow(user_id).Scan(&max_waitlist_num)
	if err != nil {
		tx.Rollback()
		return err
	}

	if max_waitlist_num.Valid {
		if max_waitlist_num.Int64 == 9 {
			// waitlist_numが9の場合は、既存のwaitlist_num==9のタスクをnullにセット
			ins, err := tx.Prepare("UPDATE tasks SET waitlist_num = $1 WHERE waitlist_num = 9 AND user_id = $2")
			if err != nil {
				tx.Rollback()
				return err
			}
			_, err = ins.Exec(sql.NullInt32{}, user_id)
			if err != nil {
				tx.Rollback()
				return err
			}

			ins, err = tx.Prepare("UPDATE tasks SET waitlist_num = $1 WHERE user_id = $2 AND id = $3")
			if err != nil {
				tx.Rollback()
				return err
			}
			_, err = ins.Exec(9, user_id, task_id)
			if err != nil {
				tx.Rollback()
				return err
			}
		} else {
			// waitlist_numが9未満の場合は、waitlist_numを+1してセット
			ins, err := tx.Prepare("UPDATE tasks SET waitlist_num = $1 WHERE user_id = $2 AND id = $3")
			if err != nil {
				tx.Rollback()
				return err
			}
			_, err = ins.Exec(max_waitlist_num.Int64+1, user_id, task_id)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	} else {
		// waitlist_numが設定されていない場合は、waitlist_num=0(リストの先頭)をセット
		ins, err := tx.Prepare("UPDATE tasks SET waitlist_num = $1 WHERE user_id = $2 AND id = $3")
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = ins.Exec(0, user_id, task_id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// waitlist_numのすべてを更新する
func ReorderWaitlist(user_id string, task_ids []int) error {
	db := database.SetupDatabase()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	ins, err := tx.Prepare("UPDATE tasks SET waitlist_num = $1 WHERE user_id = $2 AND id = $3")
	if err != nil {
		tx.Rollback()
		return err
	}
	// waitlist_numをnullにセット
	for _, task_id := range task_ids {
		_, err = ins.Exec(sql.NullInt32{}, user_id, task_id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// waitlist_numがすべてnullになっているか確認
	ins2, err2 := tx.Prepare("SELECT COUNT(*) FROM tasks WHERE user_id = $1 AND waitlist_num IS NOT NULL")
	if err2 != nil {
		tx.Rollback()
		return err2
	}
	var count int
	err2 = ins2.QueryRow(user_id).Scan(&count)
	if err2 != nil {
		tx.Rollback()
		return err2
	}
	if count != 0 {
		tx.Rollback()
		return errors.New("provided task_ids not include all tasks in waitlist")
	}

	// waitlist_numを更新
	for i, task_id := range task_ids {
		_, err = ins.Exec(i, user_id, task_id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
