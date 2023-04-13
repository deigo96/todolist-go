package controller

import (
	"net/http"

	"todolist/domain"

	"github.com/gin-gonic/gin"
)

func (tr *TodolistController) CreateActivity(c *gin.Context) {
	var req domain.ActivityParams
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.BuildError("Bad request", "invalid request body"))
		return
	}

	activity, err := tr.Todolist.CreateActivity(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.BuildError("Internal server error", "Something went wrong"))
		return
	}

	c.JSON(http.StatusCreated, domain.BuildResponse("Success", "Success", activity))

}

func (tr *TodolistController) GetAllActivity(c *gin.Context) {
	all, err := tr.Todolist.GetAllActivity(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.BuildError("Something went wrong", "Internal server error"))
		return
	}

	c.JSON(http.StatusOK, domain.BuildResponse("Success", "Success", all))
}

func (tr *TodolistController) GetActivityById(c *gin.Context) {
	id := c.Param("id")

	activity, err := tr.Todolist.GetActivityById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.BuildError("Something went wrong", "Internal server error"))
		return
	}

	if activity == nil {
		message := "Activity with ID " + id + " Not Found"
		c.JSON(http.StatusNotFound, domain.BuildError("Not Found", message))
		return
	}

	c.JSON(http.StatusOK, domain.BuildResponse("Success", "Success", activity))
}

func (tr *TodolistController) UpdateActivity(c *gin.Context) {
	id := c.Param("id")
	var req domain.ActivityParams
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.BuildError("Bad request", "invalid request body"))
		return
	}

	if req.Title == "" {
		c.JSON(http.StatusBadRequest, domain.BuildError("Bad request", "title cannot be null"))
		return
	}

	activity, err := tr.Todolist.UpdateActivity(c, id, req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.BuildError("Something went wrong", "Internal server error"))
		return
	}

	if activity == nil {
		message := "Activity with ID " + id + " Not Found"
		c.JSON(http.StatusNotFound, domain.BuildError("Not Found", message))
		return
	}

	c.JSON(http.StatusOK, domain.BuildResponse("Success", "Success", activity))
}

func (tr *TodolistController) DeleteActivity(c *gin.Context) {
	id := c.Param("id")
	if err := tr.Todolist.DeleteActivity(c, id); err != nil {
		message := "Activity with ID " + id + " Not Found"
		c.JSON(http.StatusNotFound, domain.BuildError("Not Found", message))
		return
	}

	c.JSON(http.StatusOK, domain.BuildResponse("Success", "Success", domain.EmptyObj{}))
}
