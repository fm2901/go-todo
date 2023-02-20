package repository

import (
	"github.com/fm2901/go-todo"
	"github.com/jmoiron/sqlx"
)

type AuthMysql struct {
	db *sqlx.DB
}

func NewAuthMysql(db *sqlx.DB) *AuthMysql {
	return &AuthMysql{db: db}
}

func (c *AuthMysql) CreateUser(user todo.User) (int, error) {
	return 0, nil
}
