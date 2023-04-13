package controller

import (
	"net/http"

	"todolist/domain"

	"github.com/gin-gonic/gin"
)

type TodolistController struct {
	Todolist domain.TodolistUsecase
}

func (tc *TodolistController) GetAllTodo(c *gin.Context) {
	qry := c.Query("activity_group_id")

	all, err := tc.Todolist.GetAllTodo(c, qry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.BuildError("Something went wrong", "Internal server error"))
		return
	}

	if len(all) == 0 {
		message := "Todo with activity group id " + qry + " Not Found"
		c.JSON(http.StatusNotFound, domain.BuildError("Not Found", message))
		return
	}

	c.JSON(http.StatusOK, domain.BuildResponse("Success", "Success", all))
}

func (tc *TodolistController) GetTodoById(c *gin.Context) {
	id := c.Param("id")

	todo, err := tc.Todolist.GetTodoById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.BuildError("Something went wrong", "Internal server error"))
		return
	}

	if todo == nil {
		message := "Todo with ID " + id + " Not Found"
		c.JSON(http.StatusNotFound, domain.BuildError("Not Found", message))
		return
	}

	c.JSON(http.StatusOK, domain.BuildResponse("Success", "Success", todo))
}

func (tc *TodolistController) CreatedTodo(c *gin.Context) {
	var req domain.TodoParams
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.BuildError("Bad request", "invalid request body"))
		return
	}

	if req.Title == nil {
		c.JSON(http.StatusBadRequest, domain.BuildError("Bad request", "title cannot be null"))
		return
	}

	if req.ActivityGroupId == nil {
		c.JSON(http.StatusBadRequest, domain.BuildError("Bad request", "activity group id cannot be null"))
		return
	}

	if req.IsActive == nil {
		c.JSON(http.StatusBadRequest, domain.BuildError("Bad request", "is active cannot be null"))
		return
	}

	activity, err := tc.Todolist.CreateTodo(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.BuildError("Internal server error", "Something went wrong"))
		return
	}

	c.JSON(http.StatusCreated, domain.BuildResponse("Success", "Success", activity))
}

func (tr *TodolistController) UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var req domain.TodoParams
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.BuildError("Bad request", "invalid request body"))
		return
	}

	todo, err := tr.Todolist.UpdateTodo(c, id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.BuildError("Something went wrong", "Internal server error"))
		return
	}

	if todo == nil {
		message := "Todo with ID " + id + " Not Found"
		c.JSON(http.StatusNotFound, domain.BuildError("Not Found", message))
		return
	}

	c.JSON(http.StatusOK, domain.BuildResponse("Success", "Success", todo))
}

func (tr *TodolistController) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if err := tr.Todolist.DeleteTodo(c, id); err != nil {
		message := "Todo with ID " + id + " Not Found"
		c.JSON(http.StatusNotFound, domain.BuildError("Not Found", message))
		return
	}

	c.JSON(http.StatusOK, domain.BuildResponse("Success", "Success", domain.EmptyObj{}))
}
