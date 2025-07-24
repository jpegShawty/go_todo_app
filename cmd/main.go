package main

import (
	"log"

	"github.com/jpegShawty/go_todo_app"
	todo "github.com/jpegShawty/go_todo_app/pkg"
	"github.com/jpegShawty/go_todo_app/pkg/handler"
)

func main(){
	handler := new(handler.Handler)

	srv := new(todo.Server) // Инициализируем экземпляр сервера
	if err := srv.Run("8080", handler.InitRoutes()); err != nil{
		log.Fatal("error while running http server: %s", err.Error())

	}



}