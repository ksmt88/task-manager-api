package domain

import (
	_ "github.com/mattn/go-sqlite3"
)

type Project struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Projects []Project
