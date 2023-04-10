package route

import (
	"time"

	"github.com/deigo96/todolist-go.git/api/controller"
	"github.com/deigo96/todolist-go.git/bootstrap"
	"github.com/deigo96/todolist-go.git/repository"
	"github.com/deigo96/todolist-go.git/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewTodoListRoute(env *bootstrap.Env, timeout time.Duration, db gorm.DB, gin *gin.RouterGroup) {
	tr := repository.NewTodolistRepo(&db)
	tc := &controller.TodolistController{
		Todolist: usecase.NewTodolistUsecase(tr, timeout),
	}

	route := gin.Group("")

	route.GET("/todolist", tc.GetTodolist)
}