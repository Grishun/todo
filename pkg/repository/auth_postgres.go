package repository

import (
	"fmt"
	"github.com/Grishun/todo"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", userTab)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (user todo.User, err error) {
	// Строим запрос к базе данных
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 AND password_hash=$2", userTab)

	row := r.db.QueryRow(query, username, password)

	// Сканируем результат в структуру пользователя
	err = row.Scan(&user.Id, &user.Name, &user.Username, &user.Password)

	return user, err
}
