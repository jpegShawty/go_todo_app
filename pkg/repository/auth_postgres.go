package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	todo "github.com/jpegShawty/go_todo_app/pkg"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres{
	return &AuthPostgres{db: db}
}

// NewAuthPostgres.CreateUser()  - интерфейс Authorization структуры Repository
func (r *AuthPostgres) CreateUser(user todo.User) (int, error){
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil{
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (todo.User, error){
	var user todo.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	// записываем в структуру user результат выборки
	err := r.db.Get(&user, query, username, password)

	return user, err
}