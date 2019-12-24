package infrastructure

import (
	"database/sql"

	"github.com/ksmt88/taskManager-api/internal/task/interfaces/database"
	_ "github.com/mattn/go-sqlite3"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewHandler() *SqlHandler {
	connection, err := sql.Open("sqlite3", "task.sqlite")
	if err != nil {
		return nil
	}

	handler := new(SqlHandler)
	handler.Conn = connection

	return handler
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	var res SqlResult
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}
