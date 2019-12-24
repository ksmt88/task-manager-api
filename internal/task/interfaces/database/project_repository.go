package database

import (
	"../../domain"
)

type ProjectRepository struct {
	SqlHandler
}

func (repo *ProjectRepository) Store(p domain.Project) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO projects (name) VALUES (?)", p.Name,
	)
	if err != nil {
		return
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return
	}

	id = int(id64)
	return
}

func (repo *ProjectRepository) Update(id int, p domain.Project) (err error) {
	_, err = repo.Execute(
		"UPDATE projects SET name = ? WHERE id = ?", p.Name, id,
	)
	if err != nil {
		return
	}

	return
}

func (repo *ProjectRepository) Delete(id int) (err error) {
	_, err = repo.Execute(
		"DELETE FROM projects WHERE id = ?", id,
	)
	if err != nil {
		return
	}

	return
}

func (repo *ProjectRepository) FindById(id int) (project domain.Project, err error) {
	row, err := repo.Query(
		"SELECT id, name FROM projects WHERE id = ?", id,
	)
	defer row.Close()
	if err != nil {
		return
	}

	row.Next()
	if err = row.Scan(&project.Id, &project.Name); err != nil {
		return
	}

	return
}

func (repo *ProjectRepository) FindAll() (projects domain.Projects, err error) {
	rows, err := repo.Query("SELECT id, name FROM projects")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var project domain.Project
		if err = rows.Scan(&project.Id, &project.Name); err != nil {
			return
		}
		projects = append(projects, project)
	}

	return
}
