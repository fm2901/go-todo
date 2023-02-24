package repository

import (
	"fmt"
	"strings"

	"github.com/fm2901/go-todo"
	"github.com/jmoiron/sqlx"
)

type TodoItemMysql struct {
	db *sqlx.DB
}

func NewTodoItemMysql(db *sqlx.DB) *TodoItemMysql {
	return &TodoItemMysql{db: db}
}

func (r *TodoItemMysql) Create(listId int, item todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	createItemQuery := fmt.Sprintf("INSERT INTO %s(title, description) VALUES('%s', '%s')", todoItemsTable, item.Title, item.Description)
	result, err := r.db.Exec(createItemQuery)
	if err != nil {
		return 0, err
	}

	var itemId int64
	if itemId, err = result.LastInsertId(); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s(list_id, item_id) VALUES(%d, %d)", listsItemsTable, listId, itemId)
	result, err = tx.Exec(createUsersListQuery)
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

func (r *TodoItemMysql) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem

	query := fmt.Sprintf("SELECT ti.* FROM %s ti INNER JOIN %s li on li.item_id = ti.id_item INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = %d and ul.user_id = %d", todoItemsTable, listsItemsTable, usersListsTable, listId, userId)

	if err := r.db.Select(&items, query); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoItemMysql) GetById(userId, itemId int) (todo.TodoItem, error) {
	var item todo.TodoItem

	query := fmt.Sprintf("SELECT ti.* FROM %s ti INNER JOIN %s li on li.item_id = ti.id_item INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = %d and ul.user_id = %d", todoItemsTable, listsItemsTable, usersListsTable, itemId, userId)

	if err := r.db.Get(&item, query); err != nil {
		return item, err
	}

	return item, nil
}

func (r *TodoItemMysql) Delete(userId, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=%d", todoItemsTable, itemId)
	_, err := r.db.Exec(query)

	query = fmt.Sprintf("DELETE FROM %s WHERE user_id=%d and item_id=%d", listsItemsTable, userId, itemId)
	_, err = r.db.Exec(query)

	return err
}

func (r *TodoItemMysql) Update(userId, itemId int, input todo.UpdateItemInput) error {
	setValues := make([]string, 0)

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title='%s'", *input.Title))
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description='%s'", *input.Description))
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done='%s'", *input.Done))
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=%d)",
		todoItemsTable, setQuery, itemId)

	_, err := r.db.Exec(query)
	return err
}
