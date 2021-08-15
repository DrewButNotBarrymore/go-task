package models

type Task struct {
	ID          uint   `json:"id" gorm:"primary key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
	StatusID    uint   `json:"status_id"`
	Priv1       string `json:"priv1"`
	Priv2       string `json:"priv2"`
}

type User struct {
	ID    uint   `json:"id" gorm:"primary key"`
	Name  string `json:"name"`
	Tasks []Task
}

type History struct {
	ID     uint `json:"id" gorm:"primary key"`
	TaskID uint `json:"task_id"`
	Status uint `json:"status_id"`
	UserID uint `json:"user_id"`
}

type Status struct {
	ID    uint   `json:"id" gorm:"primary key"`
	Name  string `json:"name"`
	Tasks []Task
}
