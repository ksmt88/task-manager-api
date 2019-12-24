package database

import (
	"fmt"
	"time"

	"../../domain"
)

type TaskRepository struct {
	SqlHandler
}

func (repo *TaskRepository) Store(t domain.Task) (task domain.Task, err error) {
	result, err := repo.Execute(
		"INSERT INTO tasks (title, project_id, category_id, deadline, man_hour) VALUES (?, ?, ?, ?, ?)", t.Title, t.ProjectId, t.CategoryId, t.Deadline, t.ManHour,
	)
	if err != nil {
		return
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return
	}

	task, err = repo.FindById(int(id64))
	fmt.Printf("%v", err)
	if err != nil {
		return
	}
	return
}

func (repo *TaskRepository) Update(id int, t domain.Task) (err error) {
	dt := time.Now()
	t.UpdatedAt = fmt.Sprintf(dt.Format("2006-01-02 15:04:05"))
	_, err = repo.Execute(
		"UPDATE tasks SET title = ?, project_id = ?, category_id = ?, deadline = ?, complete_date = ?, man_hour = ?, delete_flg = ?, updated_at = ? WHERE id = ?",
		t.Title, t.ProjectId, t.CategoryId, t.Deadline, t.CompleteDate, t.ManHour, t.DeleteFlg, t.UpdatedAt, id,
	)
	if err != nil {
		return
	}

	return
}

func (repo *TaskRepository) Delete(id int) (err error) {
	_, err = repo.Execute(
		"DELETE FROM tasks WHERE id = ?", id,
	)
	if err != nil {
		return
	}

	return
}

func (repo *TaskRepository) FindById(id int) (t domain.Task, err error) {
	row, err := repo.Query(`
SELECT id,
       title,
       project_id,
       category_id,
       deadline,
       complete_date,
       man_hour,
       created_at,
       updated_at
FROM tasks
WHERE id = ?
		`, id)
	defer row.Close()
	if err != nil {
		return
	}

	row.Next()
	if err = row.Scan(&t.Id, &t.Title, &t.ProjectId,
		&t.CategoryId, &t.Deadline, &t.CompleteDate,
		&t.ManHour, &t.CreatedAt, &t.UpdatedAt); err != nil {
		return
	}

	return
}

func (repo *TaskRepository) FindAll() (tasks domain.Tasks, err error) {
	rows, err := repo.Query(`
SELECT id,
       title,
       project_id,
       category_id,
       deadline,
       complete_date,
       man_hour,
       created_at,
       updated_at
FROM tasks
		`)
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var t domain.Task
		if err = rows.Scan(&t.Id, &t.Title, &t.ProjectId,
			&t.CategoryId, &t.Deadline, &t.CompleteDate,
			&t.ManHour, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return
		}
		tasks = append(tasks, t)
	}

	return
}
