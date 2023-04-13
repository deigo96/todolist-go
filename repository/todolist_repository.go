package repository

import (
	"context"
	"errors"

	"todolist/domain"

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

func (t *todoRepo) CreateTodo(c context.Context, req domain.Todos) (*domain.Todos, error) {
	if err := t.db.Create(&req).Error; err != nil {
		return nil, err
	}

	return &req, nil
}

func (t *todoRepo) GetAllTodo(c context.Context, aId string) (record []domain.Todos, err error) {
	query := "SELECT * FROM todos"
	if aId != "" {
		query += " where activity_group_id = " + aId
	}

	if err := t.db.Raw(query).Find(&record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (t *todoRepo) GetTodoById(c context.Context, id int) (record *domain.Todos, err error) {
	if err := t.db.Table("todos").Where("id = ?", id).Scan(&record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (t *todoRepo) UpdateTodo(c context.Context, id int, req domain.TodoParams) error {
	if err := t.db.Table("todos").Where("id = ?", id).Updates(&req).Error; err != nil {
		return err
	}

	return nil

}

func (t *todoRepo) DeleteTodo(c context.Context, id int) error {
	if err := t.db.Delete(&domain.Todos{}, id); err.Error != nil || err.RowsAffected == 0 {
		return errors.New("error")
	}

	return nil
}
