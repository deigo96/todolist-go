package domain

import "context"

type TodolistUsecase interface {
	CreateActivity(context.Context, ActivityParams) (*Activity, error)
	GetAllActivity(context.Context) ([]Activity, error)
	GetActivityById(context.Context, string) (*Activities, error)
	UpdateActivity(context.Context, string, string) (*Activities, error)
	DeleteActivity(context.Context, string) error
}
