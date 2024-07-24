package models

type TodoList struct {
	Id          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Task struct {
	Id          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"is_done"`
}

type UserList struct {
	Id     int `json:"-"`
	UserId int `json:"user_id"`
	ListId int `json:"list_id"`
}

type TaskList struct {
	Id     int `json:"-"`
	TaskId int `json:"task_id"`
	ListId int `json:"list_id"`
}
