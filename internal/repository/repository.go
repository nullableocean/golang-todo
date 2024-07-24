package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type TodoList interface {
}

type TodoTask interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoTask
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
