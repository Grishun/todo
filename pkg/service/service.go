package service

import (
	"github.com/Grishun/todo"
	"github.com/Grishun/todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	CreateList(userId int, list todo.TodoList) (int, error)
	GetAllLists(userId int) ([]todo.TodoList, error)
	GetListById(userId, listId int) (todo.TodoList, error)
	DeleteList(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	NewItem(userId, listId int, input todo.TodoItem) (itemId int, err error)
	GetAllItems(userId, listId int) ([]todo.TodoItem, error)
	GetItemById(userId, itemId int) (todo.TodoItem, error)
	Update(userId, itemId int, input todo.UpdateItemInput) error
	Delete(userId, itemId int) error
}

type Service struct {
	Auth     Authorization
	Todolist TodoList
	TodoItem TodoItem
}

func NewService(rep *repository.Rep) *Service {
	return &Service{
		Auth:     NewAuthService(*rep),
		Todolist: NewListService(*rep),
		TodoItem: NewItemService(*rep),
	}
}
