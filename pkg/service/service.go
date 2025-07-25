package service

import (
	todo "github.com/jpegShawty/go_todo_app/pkg"
	"github.com/jpegShawty/go_todo_app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service{ 
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}