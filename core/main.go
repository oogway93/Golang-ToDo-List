package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"todo_list"
	handler_todo "todo_list/handler"
	"todo_list/repository"
	service2 "todo_list/service"
)

// @title ToDo List
// @version 1.0
// @description The simple ToDo List App

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file",
			err)
	}
	PORT := os.Getenv("PORT")
	db, err := repository.NewPostgresDB(repository.Config{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMode"),
	})

	if err != nil {
		logrus.Fatal("Failed to initialized db",
			err)
	}

	repos := repository.NewRepository(db)         // db layer
	services := service2.NewService(repos)        // business layer
	handlers := handler_todo.NewHandler(services) // http/handlers layer

	server := new(todo_list.Server)

	if err := server.Run(PORT, handlers.InitRoutes()); err != nil {
		logrus.Fatal("Some errors in initialization routes",
			err)
	}
}
