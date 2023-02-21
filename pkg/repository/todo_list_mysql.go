package repository

import (
	"fmt"

	"github.com/fm2901/go-todo"
	"github.com/jmoiron/sqlx"
)

type TodoListMysql struct {
	db *sqlx.DB
}

func NewTodoListMysql(db *sqlx.DB) *TodoListMysql {
	return &TodoListMysql{db: db}
}

func (r *TodoListMysql) Create(userId int, list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	createListQuery := fmt.Sprintf("INSERT INTO %s(title, description) VALUES('%s', '%s')", todoListsTable, list.Title, list.Description)
	result, err := r.db.Exec(createListQuery)
	if err != nil {
		return 0, err
	}

	var list_id int64
	if list_id, err = result.LastInsertId(); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s(user_id, list_id) VALUES(%d, %d)", usersListsTable, userId, list_id)
	result, err = r.db.Exec(createUsersListQuery)
	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = result.LastInsertId(); err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(id), tx.Commit()
}

func (r *TodoListMysql) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl LEFT JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id=%d", todoListsTable, usersListsTable, userId)
	err := r.db.Select(&lists, query)

	return lists, err
}

func (r *TodoListMysql) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl LEFT JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id=%d and tl.id=%d", todoListsTable, usersListsTable, userId, listId)
	err := r.db.Get(&list, query)

	return list, err
}

func (r *TodoListMysql) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id=ul.list_id and tl.id=%d and ul.user_id=%d", todoListsTable, usersListsTable, listId, userId)
	_, err := r.db.Exec(query)

	return err
}
