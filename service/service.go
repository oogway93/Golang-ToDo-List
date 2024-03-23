package service

import (
	"todo_list/repository"
	"todo_list/structs"
)

type ToDoList interface {
	GetAll() ([]structs.ToDo, error)
	GetById(itemId int) (structs.ToDo, error)
	Create(itemList structs.ToDo) error
	Delete(itemId int) error
	Update(itemId int, itemList structs.UpdateToDo) error
}

type ToDoItem interface {
	GetAll(listId int) ([]structs.ToDoItem, error)
}

type Service struct {
	ToDoList
	ToDoItem
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		ToDoList: NewTodoListService(repository.ToDoList),
		ToDoItem: NewToDoItemService(repository.ToDoItem, repository.ToDoList),
	}
}
