package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"todo_list/structs"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) GetAll() ([]structs.ToDo, error) {
	var items []structs.ToDo

	query := fmt.Sprintf("SELECT * FROM %s;", listTableName)
	if err := r.db.Select(&items, query); err != nil {
		return items, err
	}

	return items, nil
}

func (r *TodoListPostgres) GetById(itemId int) (structs.ToDo, error) {
	var item structs.ToDo

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1;", listTableName)
	if err := r.db.Get(&item, query, itemId); err != nil {
		return item, err
	}
	return item, nil
}

func (r *TodoListPostgres) Create(itemList structs.ToDo) error {
	bg, err := r.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2);", listTableName)
	_, err = bg.Exec(query, itemList.Title, itemList.Description)
	if err != nil {
		return err
	}
	return bg.Commit()
}

func (r *TodoListPostgres) Update(itemId int, itemList structs.UpdateToDo) error {
	values := make([]string, 0)
	args := make([]interface{}, 0)
	order := 1
	if itemList.Title != nil {
		values = append(values, fmt.Sprintf("title=$%d", order))
		args = append(args, *itemList.Title)
		order++
	}
	if itemList.Description != nil {
		values = append(values, fmt.Sprintf("description=$%d", order))
		args = append(args, *itemList.Description)
		order++
	}
	queryValues := strings.Join(values, ", ")
	query := fmt.Sprintf("UPDATE %s ti SET %s WHERE ti.id=$%d;",
		listTableName, queryValues, itemId)
	args = append(args, itemId)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TodoListPostgres) Delete(itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1;", listTableName)
	_, err := r.db.Exec(query, itemId)
	return err
}
