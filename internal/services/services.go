package services

import "github.com/nullableocean/golang-todo/internal/repository"

type Authorization interface {
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
	return &Services{}
}
