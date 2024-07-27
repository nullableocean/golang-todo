package services

import (
	"errors"
	"github.com/nullableocean/golang-todo/internal/models"
	"github.com/nullableocean/golang-todo/internal/repository"
)

type TasksService struct {
	r     repository.TodoTask
	rList repository.TodoList
}

func NewTasksService(r repository.TodoTask, rList repository.TodoList) *TasksService {
	return &TasksService{r: r, rList: rList}
}

func (s *TasksService) GetAll(userId, listId int) ([]models.Task, error) {
	if _, err := s.rList.GetListById(userId, listId); err != nil {
		return nil, errors.New("error with todo list: " + err.Error())
	}

	return s.r.GetAll(userId, listId)
}

func (s *TasksService) GetTaskById(userId, taskId int) (models.Task, error) {
	return s.r.GetTaskById(userId, taskId)
}

func (s *TasksService) Create(userId, listId int, task models.Task) (int, error) {
	if _, err := s.rList.GetListById(userId, listId); err != nil {
		return 0, errors.New("error with todo list: " + err.Error())
	}

	return s.r.Create(listId, task)
}

func (s *TasksService) Update(userId, taskId int, input models.TaskUpdateInput) error {
	err := input.Validate()
	if err != nil {
		return err
	}

	return s.r.Update(userId, taskId, input)
}

func (s *TasksService) Delete(userId, taskId int) error {
	return s.r.Delete(userId, taskId)
}
