package repository

import (
	"fmt"
	"github.com/Grishun/todo"
	"github.com/jmoiron/sqlx"
	"strings"
)

type ListPostgres struct {
	db *sqlx.DB
}

func NewListPostgres(db *sqlx.DB) *ListPostgres {
	return &ListPostgres{db: db}
}

func (r *ListPostgres) CreateList(userId int, list todo.TodoList) (listId int, err error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTab)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&listId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTab)

	if _, err = tx.Exec(createUsersListQuery, userId, listId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return listId, tx.Commit()
}

func (r *ListPostgres) GetAllLists(userId int) ([]todo.TodoList, error) {

	lists := make([]todo.TodoList, 0)

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id=ul.list_id WHERE ul.user_id=$1", todoListsTab, usersListsTab)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *ListPostgres) GetListById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id=$1 WHERE ul.user_id=$2", todoListsTab, usersListsTab)
	row := r.db.QueryRow(query, listId, userId)

	err := row.Scan(&list.Id, &list.Title, &list.Description)

	return list, err
}

func (r *ListPostgres) Update(userId, listId int, input todo.UpdateListInput) error {
	setVals := make([]string, 0)
	args := make([]any, 0)
	argId := 1

	if input.Title != nil {
		setVals = append(setVals, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setVals = append(setVals, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setVals, ", ")
	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id=ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListsTab, setQuery, usersListsTab, argId, argId+1)
	args = append(args, listId, userId)

	_, err := r.db.Exec(query, args...)

	return err
}

func (r *ListPostgres) DeleteList(userId, listId int) error {

	delFromListsQuery := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE ul.user_id=$1 AND tl.id=$2", todoListsTab, usersListsTab)
	_, err := r.db.Exec(delFromListsQuery, userId, listId)

	return err
}
