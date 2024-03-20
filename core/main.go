package main

import (
	"log"
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
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Some errors %s", err.Error())
	}

}
