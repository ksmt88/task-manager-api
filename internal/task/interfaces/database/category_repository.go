package database

import (
	"github.com/ksmt88/taskManager-api/internal/task/domain"
)

type CategoryRepository struct {
	SqlHandler
}

func (repo *CategoryRepository) Store(c domain.Category) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO categories (name) VALUES (?)", c.Name,
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

func (repo *CategoryRepository) Update(id int, p domain.Category) (err error) {
	_, err = repo.Execute(
		"UPDATE categories SET name = ? WHERE id = ?", p.Name, id,
	)
	if err != nil {
		return
	}

	return
}

func (repo *CategoryRepository) Delete(id int) (err error) {
	_, err = repo.Execute(
		"DELETE FROM categories WHERE id = ?", id,
	)
	if err != nil {
		return
	}

	return
}

func (repo *CategoryRepository) FindById(id int) (category domain.Category, err error) {
	row, err := repo.Query(
		"SELECT id, name FROM categories WHERE id = ?", id,
	)
	defer row.Close()
	if err != nil {
		return
	}

	row.Next()
	if err = row.Scan(&category.Id, &category.Name); err != nil {
		return
	}

	return
}

func (repo *CategoryRepository) FindAll() (categories domain.Categories, err error) {
	rows, err := repo.Query("SELECT id, name FROM categories")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var category domain.Category
		if err = rows.Scan(&category.Id, &category.Name); err != nil {
			return
		}
		categories = append(categories, category)
	}

	return
}
