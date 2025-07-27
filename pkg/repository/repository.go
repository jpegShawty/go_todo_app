package repository

import (
	"github.com/jmoiron/sqlx"
	todo "github.com/jpegShawty/go_todo_app/pkg"
)
// Для чего интерфейс? Чтобы можно было создать NewAuthMongo итд.
// То есть поменять реализацию, вот и все

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId int, listId int) (todo.TodoList, error)
	Delete(serId int, listId int) error
	Update(userId int, listId int, input todo.UpdateListInput) error
 }

type TodoItem interface {
	Create(listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId int, listId int) (todo.TodoItem, error)
	Delete(userId int, itemId int) error
	Update(userId int, listId int, input todo.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList: NewTodoListPostgres(db),
		TodoItem: NewTodoItemPostgres(db),
	}
}