package repository

import (
	"github.com/Grishun/todo"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	CreateList(userid int, list todo.TodoList) (int, error)
	GetAllLists(userId int) ([]todo.TodoList, error)
	GetListById(userId, listId int) (todo.TodoList, error)
	Update(userId, listId int, input todo.UpdateListInput) error
	DeleteList(userId, listId int) error
}

type TodoItem interface {
	NewItem(listId int, input todo.TodoItem) (int, error)
	GetAllItems(userId, listId int) ([]todo.TodoItem, error)
	GetItemById(userId, itemId int) (todo.TodoItem, error)
	Update(userId, itemId int, input todo.UpdateItemInput) error
	Delete(userId, itemId int) error
}

type Rep struct {
	Authorization
	TodoList
	TodoItem
}

func NewRep(db *sqlx.DB) *Rep {
	return &Rep{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewListPostgres(db),
		TodoItem:      NewItemPostgres(db),
	}
}
