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

func (r *ToDoItemPostgres) Create(listId int, item structs.ToDoItem) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", itemTableName)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", listItemsTableName)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
