package services

import (
	"github.com/nullableocean/golang-todo/internal/models"
	"github.com/nullableocean/golang-todo/internal/repository"
)

type TodoService struct {
	repo repository.TodoList
}

func NewTodoService(repo repository.TodoList) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) GetAll(userId int) ([]models.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoService) GetListById(userId, listId int) (models.TodoList, error) {
	return s.repo.GetListById(userId, listId)
}

func (s *TodoService) Create(userId int, list models.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoService) Update(userId, listId int, input models.TodoListUpdateInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, listId, input)
}

func (s *TodoService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}
