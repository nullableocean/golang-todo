package services

import (
	"github.com/nullableocean/golang-todo/internal/models"
	"github.com/nullableocean/golang-todo/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	FindUser(username, password string) (models.User, error)
	GenerateJwtToken(user models.User) (string, error)
	ParseToken(tokenString string) (int, error)
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

type Services struct {
	Authorization
	TodoList
	TodoTask
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoService(repo.TodoList),
	}
}
