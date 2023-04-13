package route

import (
	"time"

	"todolist/api/controller"
	"todolist/bootstrap"
	"todolist/repository"
	"todolist/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewTodoListRoute(env *bootstrap.Env, timeout time.Duration, db gorm.DB, gin *gin.RouterGroup) {
	tr := repository.NewTodolistRepo(&db)
	tc := &controller.TodolistController{
		Todolist: usecase.NewTodolistUsecase(tr, timeout),
	}

	route := gin.Group("")

	// activity
	route.GET("/activity-groups", tc.GetAllActivity)
	route.POST("/activity-groups", tc.CreateActivity)
	route.GET("/activity-groups/:id", tc.GetActivityById)
	route.PATCH("/activity-groups/:id", tc.UpdateActivity)
	route.DELETE("/activity-groups/:id", tc.DeleteActivity)

	// todo
	route.POST("/todo-items", tc.CreatedTodo)
	route.GET("/todo-items", tc.GetAllTodo)
	route.GET("/todo-items/:id", tc.GetTodoById)
	route.PATCH("/todo-items/:id", tc.UpdateTodo)
	route.DELETE("/todo-items/:id", tc.DeleteTodo)
}
