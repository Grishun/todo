package service

import (
	"github.com/Grishun/todo"
	"github.com/Grishun/todo/pkg/repository"
)

type ListService struct {
	rep repository.Rep
}

func NewListService(rep repository.Rep) *ListService {
	return &ListService{rep: rep}
}

func (s *ListService) CreateList(userId int, list todo.TodoList) (int, error) {
	return s.rep.CreateList(userId, list)
}

func (s *ListService) GetAllLists(userId int) ([]todo.TodoList, error) {
	return s.rep.GetAllLists(userId)
}

func (s *ListService) GetListById(userId, listId int) (todo.TodoList, error) {
	return s.rep.GetListById(userId, listId)
}

func (s *ListService) Update(userId, listId int, input todo.UpdateListInput) error {
	err := todo.ValidateUpdate(input)
	if err != nil {
		return err
	}
	return s.rep.TodoList.Update(userId, listId, input)

}

func (s *ListService) DeleteList(userId, listId int) error {
	return s.rep.DeleteList(userId, listId)
}
