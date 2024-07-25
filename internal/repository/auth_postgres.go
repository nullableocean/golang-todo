package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/nullableocean/golang-todo/internal/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password) VALUES ($1, $2, $3) RETURNING ID", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	fmt.Println("CREATED_PASS_HASH: ", user.Password)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, passwordHash string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", usersTable)
	fmt.Println("SEARCH_PASS_HASH: ", passwordHash, username)
	err := r.db.Get(&user, query, username, passwordHash)

	return user, err
}
