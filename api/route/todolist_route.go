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

	// activity
	route.GET("/activity-groups", tc.GetAllActivity)
	route.POST("/activity-groups", tc.CreateActivity)
	route.GET("/activity-groups/:id", tc.GetActivityById)
	route.PATCH("/activity-groups/:id", tc.UpdateActivity)
	route.DELETE("/activity-groups/:id", tc.DeleteActivity)
}
