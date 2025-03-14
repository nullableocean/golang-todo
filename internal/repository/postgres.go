package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	tasksTable      = "tasks"
	usersListsTable = "users_lists"
	tasksListsTable = "tasks_lists"
)

func NewPostgresDB(cnf Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cnf.Host, cnf.Port, cnf.Username, cnf.DBName, cnf.Password, cnf.SSLMode))

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
