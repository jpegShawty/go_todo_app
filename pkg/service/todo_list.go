package service

import (
	todo "github.com/jpegShawty/go_todo_app/pkg"
	"github.com/jpegShawty/go_todo_app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService{
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error){
// передаем данные на репозиторий при создании списка
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error){
	return s.repo.GetAll(userId)
}
func (s *TodoListService) GetById(userId int, listId int) (todo.TodoList, error){
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Delete(userId int, listId int) (error){
	return s.repo.Delete(userId, listId)
}

func (s *TodoListService) Update(userId int, listId int, input todo.UpdateListInput) error{
	if err := input.Validate(); err != nil{
		return err
	}
	return s.repo.Update(userId, listId, input)
}