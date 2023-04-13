package usecase

import (
	"context"
	"time"

	"todolist/domain"
)

type todolist struct {
	todoList domain.TodolistRepository
	timeout  time.Duration
}

func NewTodolistUsecase(todoList domain.TodolistRepository, timeout time.Duration) domain.TodolistUsecase {
	return &todolist{
		todoList: todoList,
		timeout:  timeout,
	}
}

func (t *todolist) CreateActivity(c context.Context, req domain.ActivityParams) (*domain.Activity, error) {
	ctx, cancel := context.WithTimeout(c, t.timeout)
	defer cancel()

	request := domain.Activities{
		Title:     req.Title,
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	activity, err := t.todoList.CreateActivity(ctx, request)
	if err != nil {
		return nil, err
	}

	res := domain.Activity{
		Id:        int64(activity.ID),
		Title:     activity.Title,
		Email:     activity.Email,
		CreatedAt: domain.TimeToString(activity.CreatedAt),
		UpdatedAt: domain.TimeToString(activity.UpdatedAt),
	}

	return &res, nil
}

func (t *todolist) GetAllActivity(c context.Context) ([]domain.Activity, error) {
	ctx, cancel := context.WithTimeout(c, t.timeout)
	defer cancel()

	var allActivities []domain.Activity
	var activity domain.Activity

	result, err := t.todoList.GetAllActivity(ctx)
	if err != nil {
		return nil, err
	}

	for _, val := range result {
		activity.Id = int64(val.ID)
		activity.Title = val.Title
		activity.Email = val.Email
		activity.CreatedAt = domain.TimeToString(val.CreatedAt)
		activity.UpdatedAt = domain.TimeToString(val.UpdatedAt)

		allActivities = append(allActivities, activity)
	}

	return allActivities, nil
}

func (t *todolist) GetActivityById(c context.Context, str string) (*domain.Activities, error) {
	ctx, cancel := context.WithTimeout(c, t.timeout)
	defer cancel()

	id := domain.StringToInt(str)

	activity, err := t.todoList.GetActivityById(ctx, id)
	if err != nil {
		return nil, err
	}

	return activity, nil
}

func (t *todolist) UpdateActivity(c context.Context, str, title string) (*domain.Activities, error) {
	ctx, cancel := context.WithTimeout(c, t.timeout)
	defer cancel()

	id := domain.StringToInt(str)

	err := t.todoList.UpdateActivity(ctx, id, title)
	if err != nil {
		return nil, err
	}

	activity, err := t.todoList.GetActivityById(ctx, id)
	if err != nil {
		return nil, err
	}

	return activity, nil

}

func (t *todolist) DeleteActivity(c context.Context, str string) error {
	ctx, cancel := context.WithTimeout(c, t.timeout)
	defer cancel()

	id := domain.StringToInt(str)

	err := t.todoList.DeleteActivity(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (t *todolist) CreateTodo(c context.Context, req domain.TodoParams) (*domain.Todos, error) {
	ctx, cancel := context.WithTimeout(c, t.timeout)
	defer cancel()

	request := domain.Todos{
		Title:           *req.Title,
		ActivityGroupId: *req.ActivityGroupId,
		IsActive:        *req.IsActive,
		Priority:        req.Priority,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	todos, err := t.todoList.CreateTodo(ctx, request)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (t *todolist) GetAllTodo(c context.Context, aId string) ([]domain.Todos, error) {
	ctx, cancel := context.WithTimeout(c, t.timeout)
	defer cancel()

	allTodos, err := t.todoList.GetAllTodo(ctx, aId)
	if err != nil {
		return nil, err
	}

	return allTodos, nil
}

func (t *todolist) GetTodoById(c context.Context, str string) (*domain.Todos, error) {
	ctx, cancel := context.WithTimeout(c, t.timeout)
	defer cancel()

	id := domain.StringToInt(str)

	todo, err := t.todoList.GetTodoById(ctx, id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *todolist) UpdateTodo(c context.Context, str string, req domain.TodoParams) (*domain.Todos, error) {
	ctx, cancel := context.WithTimeout(c, t.timeout)
	defer cancel()

	id := domain.StringToInt(str)

	err := t.todoList.UpdateTodo(ctx, id, req)
	if err != nil {
		return nil, err
	}

	todo, err := t.todoList.GetTodoById(ctx, id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *todolist) DeleteTodo(c context.Context, str string) error {
	ctx, cancel := context.WithTimeout(c, t.timeout)
	defer cancel()

	id := domain.StringToInt(str)

	err := t.todoList.DeleteTodo(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
