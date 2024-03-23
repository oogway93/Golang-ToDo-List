package repository

import (
	"github.com/jmoiron/sqlx"
	"todo_list/structs"
)

type TodoList interface {
	GetAll() ([]structs.ToDo, error)
	GetById(itemId int) (structs.ToDo, error)
	Create(itemList structs.ToDo) error
	Delete(itemId int) error
	Update(itemId int, itemList structs.UpdateToDo) error
}

type Repository struct {
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoList: NewTodoListPostgres(db),
	}
}
