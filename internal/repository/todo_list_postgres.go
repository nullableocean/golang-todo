package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/nullableocean/golang-todo/internal/models"
	"strings"
)

type TodoListsPostgres struct {
	db *sqlx.DB
}

func NewTodoListsPostgres(db *sqlx.DB) *TodoListsPostgres {
	return &TodoListsPostgres{db: db}
}

func (r *TodoListsPostgres) GetAll(userId int) ([]models.TodoList, error) {
	var lists []models.TodoList

	query := fmt.Sprintf(
		"SELECT tl.id, tl.title, tl.description FROM %s AS tl INNER JOIN %s as ul ON tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable,
		usersListsTable)

	err := r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *TodoListsPostgres) GetListById(userId, listId int) (models.TodoList, error) {
	var list models.TodoList

	query := fmt.Sprintf(
		"SELECT tl.id, tl.title, tl.description FROM %s AS tl INNER JOIN %s as ul ON tl.id = ul.list_id WHERE ul.user_id = $1 and ul.list_id=$2",
		todoListsTable,
		usersListsTable)

	err := r.db.Get(&list, query, userId, listId)
	return list, err
}

func (r *TodoListsPostgres) Create(userId int, list models.TodoList) (int, error) {
	tran, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	queryCreateList := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tran.QueryRow(queryCreateList, list.Title, list.Description)
	err = row.Scan(&id)
	if err != nil {
		_ = tran.Rollback()
		return 0, err
	}

	queryUsersListsRelation := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tran.Exec(queryUsersListsRelation, userId, id)
	if err != nil {
		_ = tran.Rollback()
		return 0, err
	}

	return id, tran.Commit()
}

func (r *TodoListsPostgres) Update(userId, listId int, input models.TodoListUpdateInput) error {
	updatingColumns := make([]string, 0)
	updateArgs := make([]interface{}, 0)
	argsIndex := 1

	if input.Title != nil {
		updatingColumns = append(updatingColumns, fmt.Sprintf("title=$%d", argsIndex))
		updateArgs = append(updateArgs, *input.Title)
		argsIndex++
	}
	if input.Description != nil {
		updatingColumns = append(updatingColumns, fmt.Sprintf("description=$%d", argsIndex))
		updateArgs = append(updateArgs, *input.Description)
		argsIndex++
	}

	updateArgs = append(updateArgs, userId, listId)
	setQuerySubstr := strings.Join(updatingColumns, ", ")
	query := fmt.Sprintf(
		"UPDATE %s as tl SET %s FROM %s as ul WHERE tl.id = ul.list_id AND ul.user_id = $%d AND ul.list_id = $%d",
		todoListsTable,
		setQuerySubstr,
		usersListsTable,
		argsIndex,
		argsIndex+1)

	_, err := r.db.Exec(query, updateArgs...)
	return err
}

func (r *TodoListsPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf(
		"DELETE FROM %s as tl USING %s as ul WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2",
		todoListsTable,
		usersListsTable)

	_, err := r.db.Exec(query, userId, listId)
	return err
}
