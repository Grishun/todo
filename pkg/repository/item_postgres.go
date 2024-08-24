package repository

import (
	"fmt"
	"github.com/Grishun/todo"
	"github.com/jmoiron/sqlx"
)

type ItemPostgres struct {
	db *sqlx.DB
}

func NewItemPostgres(db *sqlx.DB) *ItemPostgres {
	return &ItemPostgres{db: db}
}

func (r *ItemPostgres) NewItem(listId int, input todo.TodoItem) (itemId int, err error) {
	tx, err := r.db.Begin()

	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemsTab)
	row := tx.QueryRow(createItemQuery, input.Title, input.Description)
	if err = row.Scan(&itemId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemQuery := fmt.Sprintf("INSERT INTO %s (item_id, list_id) VALUES ($1, $2)", listsItemsTab)
	if _, err := tx.Exec(createListItemQuery, itemId, listId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *ItemPostgres) GetAllItems(userId, listId int) (items []todo.TodoItem, err error) {
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON li.id = $1 INNER JOIN %s ul ON ul.user_id=$2", todoItemsTab, todoListsTab)

	err = r.db.Select(&items, query, listId, userId)

	return
}

func (r *ItemPostgres) GetItemById(userId, itemId int) (item todo.TodoItem, err error) {
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON ti.id=li.item_id INNER JOIN %s ul ON ul.user_id=$1 WHERE ti.id=$2", todoItemsTab, listsItemsTab, usersListsTab)

	row := r.db.QueryRow(query, userId, itemId)

	err = row.Scan(&item.Id, &item.Title, &item.Description, &item.Done)

	return
}

func (r *ItemPostgres) Update(userId, itemId int, input todo.UpdateItemInput) (err error) {

	updItemQuery := fmt.Sprintf("UPDATE %s SET title=$1, description=$2, done=$3", todoItemsTab)
	_, err = r.db.Exec(updItemQuery, input.Title, input.Description, input.Done)

	return
}

func (r *ItemPostgres) Delete(userId, itemId int) error {
	tx, err := r.db.Begin()

	delFromItemsQuery := fmt.Sprintf("DELETE FROM %s WHERE id=$1", todoItemsTab)
	if _, err = tx.Exec(delFromItemsQuery, itemId); err != nil {
		tx.Rollback()
		return err
	}
	delFromListsItemsQuery := fmt.Sprintf("DELETE FROM %s li USING %s ul WHERE ul.user_id=$1 AND li.id=$2", listsItemsTab, usersListsTab)
	if _, err = tx.Exec(delFromListsItemsQuery, userId, itemId); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
