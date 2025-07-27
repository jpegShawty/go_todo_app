package service

import (
	todo "github.com/jpegShawty/go_todo_app/pkg"
	"github.com/jpegShawty/go_todo_app/pkg/repository"
)

type TodoItemService struct{
	repo repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService{
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error){
	_, err := s.listRepo.GetById(userId,listId)
	if err != nil{
		// list doesnt exist or doesnt belong to user
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error){
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId int, listId int) (todo.TodoItem, error){
	return s.repo.GetById(userId, listId)
}

func (s *TodoItemService) Delete(userId, itemId int) error{
	return s.repo.Delete(userId,itemId)
}

func (s *TodoItemService) Update(userId int, listId int, input todo.UpdateItemInput) error{
	return s.repo.Update(userId, listId, input)
}