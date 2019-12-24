package domain

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Categories []Category
