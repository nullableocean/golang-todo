package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/nullableocean/golang-todo/internal/models"
	"strings"
)

type TasksPostgres struct {
	db *sqlx.DB
}

func NewTasksPostgres(db *sqlx.DB) *TasksPostgres {
	return &TasksPostgres{db: db}
}

func (r *TasksPostgres) GetAll(userId, listId int) ([]models.Task, error) {
	var tasks []models.Task

	query := fmt.Sprintf(
		`
		SELECT t.id, t.title, t.description, t.is_done FROM %s as t 
		INNER JOIN %s as tl ON t.id = tl.task_id
		INNER JOIN %s as ul ON tl.list_id = ul.list_id
		WHERE ul.user_id = $1 AND tl.list_id = $2
		`,
		tasksTable, tasksListsTable, usersListsTable)

	err := r.db.Select(&tasks, query, userId, listId)
	return tasks, err
}

func (r *TasksPostgres) GetTaskById(userId, taskId int) (models.Task, error) {
	var task models.Task

	query := fmt.Sprintf(
		`
		SELECT t.id, t.title, t.description, t.is_done FROM %s as t 
		INNER JOIN %s as tl ON t.id = tl.task_id
		INNER JOIN %s as ul ON tl.list_id = ul.list_id
		WHERE ul.user_id = $1 AND t.id = $2
		`,
		tasksTable, tasksListsTable, usersListsTable)

	err := r.db.Get(&task, query, userId, taskId)
	return task, err
}

func (r *TasksPostgres) Create(listId int, task models.Task) (int, error) {
	tran, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	queryCreateTask := fmt.Sprintf("INSERT INTO %s (title, description, is_done) VALUES ($1, $2, $3) RETURNING id", tasksTable)
	row := tran.QueryRow(queryCreateTask, task.Title, task.Description, task.Done)
	err = row.Scan(&id)
	if err != nil {
		_ = tran.Rollback()
		return 0, err
	}

	queryCreateTaskListRelation := fmt.Sprintf("INSERT INTO %s (task_id, list_id) VALUES ($1, $2)", tasksListsTable)
	_, err = tran.Exec(queryCreateTaskListRelation, id, listId)
	if err != nil {
		_ = tran.Rollback()
		return 0, err
	}

	return id, tran.Commit()
}

func (r *TasksPostgres) Update(userId, taskId int, input models.TaskUpdateInput) error {
	updatingColumns := make([]string, 0)
	updateArgs := make([]interface{}, 0)
	argsInd := 1

	if input.Title != nil {
		updatingColumns = append(updatingColumns, fmt.Sprintf("title=$%d", argsInd))
		updateArgs = append(updateArgs, *input.Title)
		argsInd++
	}
	if input.Description != nil {
		updatingColumns = append(updatingColumns, fmt.Sprintf("description=$%d", argsInd))
		updateArgs = append(updateArgs, *input.Description)
		argsInd++
	}
	if input.Done != nil {
		updatingColumns = append(updatingColumns, fmt.Sprintf("is_done=$%d", argsInd))
		updateArgs = append(updateArgs, *input.Done)
		argsInd++
	}

	updateArgs = append(updateArgs, userId, taskId)
	setQuerySubstr := strings.Join(updatingColumns, ", ")
	query := fmt.Sprintf(
		`
		UPDATE %s as t SET %s
		FROM %s as tl, %s as ul
		WHERE t.id = tl.task_id
		AND tl.list_id = ul.list_id AND ul.user_id = $%d
		AND t.id = $%d
		`,
		tasksTable,
		setQuerySubstr,
		tasksListsTable,
		usersListsTable,
		argsInd,
		argsInd+1)

	_, err := r.db.Exec(query, updateArgs...)
	return err
}

func (r *TasksPostgres) Delete(userId, taskId int) error {
	query := fmt.Sprintf(
		`
		DELETE FROM %s as t USING %s as tl, %s as ul
		WHERE t.id = tl.task_id AND tl.list_id = ul.list_id
		AND ul.user_id = $1 
		AND t.id = $2
		`,
		tasksTable, tasksListsTable, usersListsTable)

	_, err := r.db.Exec(query, userId, taskId)
	return err
}
