package repository

import (
	"context"
	"errors"

	"github.com/deigo96/todolist-go.git/domain"
	"gorm.io/gorm"
)

type todoRepo struct {
	db *gorm.DB
}

func NewTodolistRepo(db *gorm.DB) domain.TodolistRepository {
	return &todoRepo{
		db: db,
	}
}

func (t *todoRepo) CreateActivity(c context.Context, req domain.Activities) (*domain.Activities, error) {
	if err := t.db.Create(&req).Error; err != nil {
		return nil, err
	}

	return &req, nil

}

func (t *todoRepo) GetAllActivity(context.Context) (record []domain.Activities, err error) {
	if err := t.db.Find(&record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (t *todoRepo) GetActivityById(c context.Context, id int) (record *domain.Activities, err error) {
	if err := t.db.Table("activities").Where("id = ?", id).Scan(&record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (t *todoRepo) UpdateActivity(c context.Context, id int, title string) error {
	if err := t.db.Table("activities").Where("id = ?", id).Update("title", title).Error; err != nil {
		return nil
	}

	return nil
}

func (t *todoRepo) DeleteActivity(c context.Context, id int) error {
	if err := t.db.Delete(&domain.Activities{}, id); err.Error != nil || err.RowsAffected == 0 {
		return errors.New("error")
	}

	return nil
}
