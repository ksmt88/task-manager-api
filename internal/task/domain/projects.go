package domain

type Project struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Projects []Project
