package repository

import (
	"github.com/jmoiron/sqlx"
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
	Create(listId int, item structs.ToDoItem) error
}

type User interface {
	SignUp(user structs.UserAdd) error
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
