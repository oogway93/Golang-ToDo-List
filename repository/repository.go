package repository

import (
	"github.com/jmoiron/sqlx"
	"todo_list/structs"
)

type ToDoList interface {
	GetAll(userId int) ([]structs.ToDo, error)
	GetById(userId, listId int) (structs.ToDo, error)
	Create(userId int, itemList structs.ToDo) error
	Update(userId, listId int, itemList structs.UpdateToDo) error
	Delete(userId, listId int) error
}

type ToDoItem interface {
	GetAll(userId, listId int) ([]structs.ToDoItem, error)
	Create(listId int, item structs.ToDoItem) error
}

type User interface {
	CreateUser(user structs.User) error
	GetUser(username, password string) (structs.User, error)
}

type Repository struct {
	ToDoList
	ToDoItem
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		ToDoList: NewToDoListPostgres(db),
		ToDoItem: NewToDoItemPostgres(db),
		User:     NewAuthPostgres(db),
	}
}
