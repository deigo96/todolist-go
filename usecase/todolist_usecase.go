package usecase

import (
	"time"

	"github.com/deigo96/todolist-go.git/domain"
)

type todolist struct {
	todoList domain.TodolistRepository
	timeout time.Duration
}

func NewTodolistUsecase (todoList domain.TodolistRepository, timeout time.Duration) domain.TodolistUsecase{
	return &todolist{
		todoList: todoList,
		timeout: timeout,
	}
}