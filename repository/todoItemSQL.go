package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todo_list/structs"
)

type ToDoItemPostgres struct {
	db *sqlx.DB
}

func NewToDoItemPostgres(db *sqlx.DB) *ToDoItemPostgres {
	return &ToDoItemPostgres{db}
}

func (r *ToDoItemPostgres) GetAll(listId int) ([]structs.ToDoItem, error) {
	var items []structs.ToDoItem
	query := fmt.Sprintf("SELECT i.id, i.title, i.description, i.done FROM %s i INNER JOIN %s li ON li.item_id=i.id WHERE li.list_id=$1",
		itemTableName, listItemsTableName)
	if err := r.db.Select(&items, query, listId); err != nil {
		return nil, err
	}
	return items, nil
}
