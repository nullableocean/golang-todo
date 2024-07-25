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
	}
}
