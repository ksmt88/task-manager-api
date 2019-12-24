package domain

type Task struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	ProjectId    int    `json:"project_id"`
	CategoryId   int    `json:"category_id"`
	Deadline     string `json:"deadline"`
	CompleteDate string `json:"complete_date"`
	ManHour      int    `json:"man_hour"`
	DeleteFlg    int    `json:"delete_flg"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type Tasks []Task
