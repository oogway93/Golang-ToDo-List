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

func NewToDoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) GetAll(userId int) ([]structs.ToDo, error) {
	var items []structs.ToDo

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1;",
		listTableName, usersListsTable)
	if err := r.db.Select(&items, query, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoListPostgres) GetById(userId, listId int) (structs.ToDo, error) {
	var item structs.ToDo

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = $1 WHERE ul.user_id = $2;", listTableName, usersListsTable)
	if err := r.db.Get(&item, query, listId, userId); err != nil {
		return item, err
	}
	return item, nil
}

func (r *TodoListPostgres) Create(userId int, itemList structs.ToDo) error {
	bg, err := r.db.Begin()
	if err != nil {
		return err
	}
	var listId int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id;", listTableName)
	row := bg.QueryRow(createListQuery, itemList.Title, itemList.Description)
	if err := row.Scan(&listId); err != nil {
		bg.Rollback()
		return err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2);", usersListsTable)
	_, err = bg.Exec(createUsersListQuery, userId, listId)
	if err != nil {
		bg.Rollback()
		return err
	}

	return bg.Commit()
}

func (r *TodoListPostgres) Update(userId, listId int, itemList structs.UpdateToDo) error {
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
	query := fmt.Sprintf("UPDATE %s ti SET %s FROM %s ul WHERE ti.id=ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d;",
		listTableName, queryValues, usersListsTable, order, order+1)
	args = append(args, listId, userId)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TodoListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id=ul.list_id AND ul.user_id=$%d AND ul.list_id=$%d;", listTableName, usersListsTable, userId, listId)
	_, err := r.db.Exec(query, userId, listId)
	return err
}
