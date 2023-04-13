package domain

import "context"

type TodolistUsecase interface {
	// activity usecase
	CreateActivity(context.Context, ActivityParams) (*Activity, error)
	GetAllActivity(context.Context) ([]Activity, error)
	GetActivityById(context.Context, string) (*Activities, error)
	UpdateActivity(context.Context, string, string) (*Activities, error)
	DeleteActivity(context.Context, string) error

	// todo usecase
	CreateTodo(context.Context, TodoParams) (*Todos, error)
	GetAllTodo(context.Context, string) ([]Todos, error)
	GetTodoById(context.Context, string) (*Todos, error)
	UpdateTodo(context.Context, string, TodoParams) (*Todos, error)
	DeleteTodo(context.Context, string) error
}
