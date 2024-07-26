package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nullableocean/golang-todo/internal/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type TodoList interface {
	GetAll(userId int) ([]models.TodoList, error)
	GetListById(userId, listId int) (models.TodoList, error)
	Create(userId int, list models.TodoList) (int, error)
	Update(userId, listId int, input models.TodoListUpdateInput) error
	Delete(userId, listId int) error
}

type TodoTask interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoTask
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListsPostgres(db),
	}
}
