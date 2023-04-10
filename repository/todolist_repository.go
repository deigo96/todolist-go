package repository

import (
	"github.com/deigo96/todolist-go.git/domain"
	"gorm.io/gorm"
)

type todoRepo struct {
	db *gorm.DB
}

func NewTodolistRepo(db *gorm.DB) domain.TodolistRepository{
	return &todoRepo{
		db: db,
	}
}