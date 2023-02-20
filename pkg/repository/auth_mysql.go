package repository

import (
	"fmt"

	"github.com/fm2901/go-todo"
	"github.com/jmoiron/sqlx"
)

type AuthMysql struct {
	db *sqlx.DB
}

func NewAuthMysql(db *sqlx.DB) *AuthMysql {
	return &AuthMysql{db: db}
}

func (r *AuthMysql) CreateUser(user todo.User) (int, error) {
	var id int64
	sql := fmt.Sprintf("INSERT INTO %s(name, username, password_hash) values('%s', '%s', '%s')", usersTable, user.Name, user.Userame, user.Password)
	result, err := r.db.Exec(sql)

	if err != nil {
		return 0, err
	}

	if id, err = result.LastInsertId(); err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *AuthMysql) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	sql := fmt.Sprintf("SELECT id FROM %s WHERE username='%s' and password_hash='%s'", usersTable, username, password)
	err := r.db.Get(&user, sql)

	if err != nil {
		return user, err
	}

	return user, err
}
