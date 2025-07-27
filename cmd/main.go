package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	todo "github.com/jpegShawty/go_todo_app/pkg"
	"github.com/jpegShawty/go_todo_app/pkg/handler"
	"github.com/jpegShawty/go_todo_app/pkg/repository"
	"github.com/jpegShawty/go_todo_app/pkg/service"
	"github.com/spf13/viper"

	_ "github.com/jpegShawty/go_todo_app/docs"
)

// @title Todo App API
// @version 1.0
// @description API Server for TodoList Application

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

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

	go func(){
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil{
			logrus.Fatalf("error while running http server: %s", err.Error())

		}
	}()

	logrus.Print("Todoapp started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Todoapp shutting down")

	if err := srv.Shutdown(context.Background()); err != nil{
		logrus.Errorf("error occured on server while shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil{
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

// Инициализация конф. файлов
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}