package service

import (
	"github.com/Grishun/todo"
	"github.com/Grishun/todo/pkg/repository"
)

type ItemService struct {
	rep repository.Rep
}

func NewItemService(rep repository.Rep) *ItemService {
	return &ItemService{rep: rep}
}

func (s *ItemService) NewItem(userId, listId int, input todo.TodoItem) (itemId int, err error) {

	if _, err := s.rep.GetListById(userId, listId); err != nil {
		return 0, err
	}
	return s.rep.TodoItem.NewItem(listId, input)
}

func (s *ItemService) GetAllItems(userId, listId int) ([]todo.TodoItem, error) {

	return s.rep.TodoItem.GetAllItems(userId, listId)

}

func (s *ItemService) GetItemById(userId, itemId int) (todo.TodoItem, error) {
	return s.rep.GetItemById(userId, itemId)
}

func (s *ItemService) Update(userId, itemId int, input todo.UpdateItemInput) error {
	return s.rep.TodoItem.Update(userId, itemId, input)
}

func (s *ItemService) Delete(userId, itemId int) error {
	return s.rep.TodoItem.Delete(userId, itemId)
}
