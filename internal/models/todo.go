package models

import "errors"

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type Task struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title"  db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"is_done" db:"is_done"`
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

type TodoListUpdateInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type TaskUpdateInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"is_done"`
}

func (i *TodoListUpdateInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("empty input")
	}

	return nil
}

func (i *TaskUpdateInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("empty input")
	}

	return nil
}
