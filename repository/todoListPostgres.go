package repository

import "todo_list/structs"

type ToDoList interface {
	GetAll() ([]structs.ToDo, error)
	GetById(itemId int) (structs.ToDo, error)
	Create(itemList structs.ToDo) error
	Delete(itemId int) error
	Update(itemId int, itemList structs.UpdateToDo) error
}
