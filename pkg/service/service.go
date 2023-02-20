package service

import (
	"github.com/fm2901/go-todo"
	"github.com/fm2901/go-todo/pkg/repository"
)

type Authorization interface {
	CreateUser(todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(*repos.Authorization),
	}
}
