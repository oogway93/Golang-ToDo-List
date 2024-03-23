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
}

type Repository struct {
	ToDoList
	ToDoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		ToDoList: NewToDoListPostgres(db),
		ToDoItem: NewToDoItemPostgres(db),
	}
}
