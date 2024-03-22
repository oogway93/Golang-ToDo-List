package service

import (
	"todo_list/repository"
	"todo_list/structs"
)

type TodoList interface {
	GetAll() ([]structs.ToDo, error)
	GetById(itemId int) (structs.ToDo, error)
	Create(itemList structs.ToDo) error
}

type Service struct {
	TodoList
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		TodoList: NewTodoListService(repository.TodoList),
	}
}
