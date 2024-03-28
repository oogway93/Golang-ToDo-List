package service

import (
	"todo_list/repository"
	"todo_list/structs"
)

type ToDoList interface {
	GetAll(userId int) ([]structs.ToDo, error)
	GetById(userId int, listId int) (structs.ToDo, error)
	Create(userId int, itemList structs.ToDo) error
	Delete(userId, listId int) error
	Update(userId, listId int, itemList structs.UpdateToDo) error
}

type ToDoItem interface {
	GetAll(userId, listId int) ([]structs.ToDoItem, error)
	Create(listId int, item structs.ToDoItem) error
}

type User interface {
	CreateUser(user structs.User) error
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Service struct {
	ToDoList
	ToDoItem
	User
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		ToDoList: NewTodoListService(repository.ToDoList),
		ToDoItem: NewToDoItemService(repository.ToDoItem, repository.ToDoList),
		User:     NewUserRepository(repository.User),
	}
}
