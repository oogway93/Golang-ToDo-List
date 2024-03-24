package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todo_list/structs"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) SignUp(user structs.UserAdd) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (username, password) VALUES ($1, $2);", userTableName)
	_, err = r.db.Exec(query, user.Username, user.Password)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
