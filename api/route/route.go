package route

import (
	"time"

	"todolist/bootstrap"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db gorm.DB, router *gin.RouterGroup) {
	protectedRoute := router.Group("")

	NewTodoListRoute(env, timeout, db, protectedRoute)
}
