package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	todo "github.com/jpegShawty/go_todo_app/pkg"
	"github.com/jpegShawty/go_todo_app/pkg/handler"
	"github.com/jpegShawty/go_todo_app/pkg/repository"
	"github.com/jpegShawty/go_todo_app/pkg/service"
	"github.com/spf13/viper"
)

func main(){
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err !=nil{
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: "localhost",
		Port: "5436",
		Username: "postgres",
		Password: "qwerty",
		DBName: "postgres",
		SSLMode: "disable",
	})
	if err != nil{
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(todo.Server) // Инициализируем экземпляр сервера
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil{
		logrus.Fatalf("error while running http server: %s", err.Error())

	}

}

// Инициализация конф. файлов
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}