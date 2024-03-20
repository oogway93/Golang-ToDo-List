package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"todo_list"
	handler_todo "todo_list/handler"
	"todo_list/repository"
	service2 "todo_list/service"
)

func main() {
	repos := repository.NewRepository()
	services := service2.NewService(repos)
	handlers := handler_todo.NewHandler(services)

	server := new(todo_list.Server)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	if err := server.Run(PORT, handlers.InitRoutes()); err != nil {
		log.Fatalf("Some errors %s", err.Error())
	}

}
