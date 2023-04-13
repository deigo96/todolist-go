package domain

import (
	"context"
	"time"
)

type Activities struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TodolistRepository interface {
	// activity
	CreateActivity(context.Context, Activities) (*Activities, error)
	GetAllActivity(context.Context) ([]Activities, error)
	GetActivityById(context.Context, int) (*Activities, error)
	UpdateActivity(context.Context, int, string) error
	DeleteActivity(context.Context, int) error

	// todo
	CreateTodo(context.Context, Todos) (*Todos, error)
	GetAllTodo(context.Context, string) ([]Todos, error)
	GetTodoById(context.Context, int) (*Todos, error)
	UpdateTodo(context.Context, int, TodoParams) error
	DeleteTodo(context.Context, int) error
}
