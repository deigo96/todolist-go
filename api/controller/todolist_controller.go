package controller

import (
	"net/http"

	"github.com/deigo96/todolist-go.git/domain"
	"github.com/gin-gonic/gin"
)

type TodolistController struct {
	Todolist domain.TodolistUsecase
}

func (tc *TodolistController) GetTodolist (c *gin.Context) {


	c.JSON(http.StatusOK, domain.BuildResponse(http.StatusOK, "Success", domain.EmptyObj{}))
}